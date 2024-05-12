package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/controller"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func NewRouter(productController controller.ProductController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	// products
	// router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	// router.PUT("/api/products/:productId", productController.Update)
	// router.DELETE("/api/products/:productId", productController.Delete)

	// users
	router.POST("/api/users", userController.Create)
	router.POST("/api/users/login", userController.Login)

	router.PanicHandler = wrapper.ErrorHandler

	return router
}
