package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/jwt"
	"github.com/muhhylmi/store-api/utils/objects"
	"github.com/muhhylmi/store-api/utils/wrapper"
	"golang.org/x/crypto/bcrypt"
)

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	_, errCheckUserName := service.Repository.FindByUsername(ctx, request.Username)
	if errCheckUserName == nil {
		l.Error("Username Already Exists")
		panic(wrapper.NewStatuConflictError("username already exists"))
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if errHash != nil {
		l.Error(err)
		panic(err)
	}
	user := domain.Users{
		BaseModel: domain.BaseModel{
			ID:        uuid.NewString(),
			IsDeleted: objects.ToPointer(false),
		},
		Username: request.Username,
		Role:     request.Role,
		Password: string(hashedPassword),
	}
	result, err := service.Repository.Save(ctx, user)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToUserRersponse(result)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	l := service.Logger.LogWithContext("product_service", "FindById")

	checkUser, errCheck := service.Repository.FindByUsername(ctx, request.Username)
	if errCheck != nil {
		l.Error("User Is Not Found")
		panic(wrapper.NewNotFoundError("user is not found"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(request.Password))
	if err != nil {
		l.Error(err)
		panic(err)
	}

	token, err := jwt.CreateToken(checkUser, service.Config)
	if err != nil {
		l.Error(err)
		panic(err)
	}

	return web.ToLoginResponse(*checkUser, token)
}
