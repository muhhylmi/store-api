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
	ShoppingCartId string     `json:"shoppingCartId"`
	Items          []CartItem `json:"items"`
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
	ShoppingCartId string             `json:"shoppingCartId"`
	Items          []CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ProductId    string `json:"productId"`
	ProductPrice int64  `json:"productPrice"`
	ProductName  string `json:"productName"`
	Qty          int    `json:"qty"`
}

func ToProductIds(req []CartItem) []string {
	result := []string{}
	for _, item := range req {
		result = append(result, item.ProductId)
	}
	return result
}

func ToShoppingCartModel(req ShopingCartCreateRequest) domain.ShoppingCarts {
	result := domain.ShoppingCarts{}
	shoppingCartId := uuid.NewString()
	result.ID = shoppingCartId
	result.IsDeleted = objects.ToPointer(false)
	result.CreatedBy = &req.AuthData.UserId
	result.UserId = req.AuthData.UserId
	result.Status = PENDING_CART
	for _, cart := range req.Items {
		result.Items = append(result.Items, domain.ShoppingCartItems{
			BaseModel: domain.BaseModel{
				ID:        uuid.NewString(),
				IsDeleted: objects.ToPointer(false),
				CreatedBy: &req.AuthData.UserId,
			},
			ShoppingCartId: shoppingCartId,
			ProductId:      cart.ProductId,
			Qty:            cart.Qty,
		})
	}
	return result
}

func ToShoopingCartRersponse(cart domain.ShoppingCarts) ShopingCartResponse {
	result := ShopingCartResponse{}
	result.ShoppingCartId = cart.ID
	for _, item := range cart.Items {
		result.Items = append(result.Items, CartItem{
			ProductId: item.ProductId,
			Qty:       item.Qty,
		})
	}
	return result
}

func ToCartListResponse(carts []*domain.ShoppingCarts) []ListCartResponse {
	result := make([]ListCartResponse, len(carts))
	for i, cart := range carts {
		result[i].ShoppingCartId = cart.ID
		for _, item := range cart.Items {
			result[i].Items = append(result[i].Items, CartItemResponse{
				ProductId:    item.ProductId,
				ProductName:  item.Product.ProductName,
				ProductPrice: item.Product.Price,
				Qty:          item.Qty,
			})
		}
	}
	return result
}
