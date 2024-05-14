package middleware

import (
	"net/http"
	"strings"

	"go-store-api/model/web"
	"go-store-api/utils/config"
	"go-store-api/utils/jwt"
	"go-store-api/utils/objects"
	"go-store-api/utils/wrapper"
)

type AuthMiddleware struct {
	Handler http.Handler
	Config  *config.Configurations
}

func NewAuthMiddleware(handler http.Handler, config *config.Configurations) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
		Config:  config,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	bearerToken := request.Header.Get("Authorization")

	if strings.HasPrefix(request.RequestURI, "/api/users") && request.Header.Get("X-API-Key") == middleware.Config.API_KEY {
		middleware.Handler.ServeHTTP(writer, request)
	} else if bearerToken != "" {
		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Invalid Token Format",
			}
			wrapper.WriteToResponseBody(writer, webResponse)
			return
		}

		jwtClaim, err := jwt.ValidateJwt(token[1], middleware.Config)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   err.Error(),
			}
			wrapper.WriteToResponseBody(writer, webResponse)
			return
		}

		if ok := CheckRole(jwtClaim.Role, request.RequestURI, request.Method); !ok {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "role is restricted",
			}
			wrapper.WriteToResponseBody(writer, webResponse)
			return
		}

		request.Header.Set("userId", jwtClaim.UserId)
		request.Header.Set("Role", jwtClaim.Role)

		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		wrapper.WriteToResponseBody(writer, webResponse)
	}

}

func CheckRole(role string, uri string, method string) bool {
	checkUri := objects.AnyInSlice([]string{
		"/api/products",
		"/api/categories",
	}, uri)
	if checkUri && (method == "POST" || method == "PUT") && role != "ADMIN" {
		return false
	}

	return true
}
