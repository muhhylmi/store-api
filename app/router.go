package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/controller"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func NewRouter(productController controller.ProductController, userController controller.UserController,
	categoryController controller.CategoryController, cartController controller.ShoppingCartController) *httprouter.Router {
	router := httprouter.New()

	// products
	router.GET("/api/products", productController.List)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)

	// users
	router.POST("/api/users", userController.Create)
	router.POST("/api/users/login", userController.Login)
	router.PUT("/api/users/top-up/:user_id", userController.TopUp)

	//categories
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories", categoryController.List)

	// shopping cart
	router.GET("/api/shopping-cart", cartController.List)
	router.POST("/api/shopping-cart", cartController.Create)
	router.PATCH("/api/shopping-cart/checkout", cartController.Checkout)
	router.PUT("/api/shopping-cart/:shopping_cart_id", cartController.Update)
	router.DELETE("/api/shopping-cart/:shopping_cart_id", cartController.Delete)

	router.PanicHandler = wrapper.ErrorHandler
	return router
}
