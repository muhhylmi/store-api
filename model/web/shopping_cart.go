package web

import (
	"github.com/google/uuid"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/objects"
)

type ShopingCartCreateRequest struct {
	AuthData

	Items []CartItem `json:"items"`
}

type ShopingCartResponse struct {
	ShoppingCartId string `json:"shoppingCartId"`
	ProductId      string `json:"productId"`
	Qty            int    `json:"qty"`
}

type CartItem struct {
	ProductId string `json:"productId"`
	Qty       int    `json:"qty"`
}

type ListCartRequest struct {
	AuthData
	Status string `query:"status" validate:"oneof=PENDING SUCCESS"`
}

type ListCartResponse struct {
	ShoppingCartId string `json:"shoppingCartId"`
	ProductId      string `json:"productId"`
	ProductPrice   int64  `json:"productPrice"`
	ProductName    string `json:"productName"`
	Qty            int    `json:"qty"`
}

type CartItemResponse struct {
	ProductId    string `json:"productId"`
	ProductPrice int64  `json:"productPrice"`
	ProductName  string `json:"productName"`
	Qty          int    `json:"qty"`
}

type UpdateCartRequest struct {
	AuthData

	ShoppingCartId string `params:"shopping_cart_id"`
	ProductId      string `json:"productId"`
	Qty            int    `json:"qty"`
}

type DeleteCartRequest struct {
	AuthData
	ShoppingCartId string `params:"shopping_cart_id"`
}

type CheckoutCartRequest struct {
	AuthData

	ShoppingCartIds []string `json:"shoppingCartIds"`
}

type CheckoutResponse struct {
	ShoppingCartIds []string `json:"shoopingCartIds"`
	Message         string   `json:"message"`
}

func ToProductIds(req []CartItem) []string {
	result := []string{}
	for _, item := range req {
		result = append(result, item.ProductId)
	}
	return result
}

func ToShoppingCartModel(req ShopingCartCreateRequest) []domain.ShoppingCarts {
	result := []domain.ShoppingCarts{}
	for _, item := range req.Items {
		result = append(result, domain.ShoppingCarts{
			BaseModel: domain.BaseModel{
				ID:        uuid.NewString(),
				IsDeleted: objects.ToPointer(false),
				CreatedBy: &req.AuthData.UserId,
			},
			UserId:    req.AuthData.UserId,
			Status:    PENDING_CART,
			ProductId: item.ProductId,
			Qty:       item.Qty,
		})
	}
	return result
}

func ToShoopingCartRersponse(carts []domain.ShoppingCarts) []ShopingCartResponse {
	result := []ShopingCartResponse{}

	for _, cart := range carts {
		result = append(result, ShopingCartResponse{
			ShoppingCartId: cart.ID,
			ProductId:      cart.ProductId,
			Qty:            cart.Qty,
		})
	}
	return result
}

func ToCartListResponse(carts []*domain.ShoppingCarts) []ListCartResponse {
	result := []ListCartResponse{}

	for _, cart := range carts {
		result = append(result, ListCartResponse{
			ShoppingCartId: cart.ID,
			ProductId:      cart.ProductId,
			ProductName:    cart.Product.ProductName,
			ProductPrice:   cart.Product.Price,
			Qty:            cart.Qty,
		})
	}
	return result
}

func ToUpdateShopingCartResponse(cart domain.ShoppingCarts) ShopingCartResponse {
	return ShopingCartResponse{
		ShoppingCartId: cart.ID,
		ProductId:      cart.ProductId,
		Qty:            cart.Qty,
	}
}

func GetTotalPrice(carts []domain.ShoppingCarts) int64 {
	result := int64(0)
	for _, cart := range carts {
		result += cart.Product.Price * int64(cart.Qty)
	}
	return result
}
