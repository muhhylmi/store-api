package web

import "github.com/muhhylmi/store-api/model/domain"

type ProductCreateRequest struct {
	AuthData

	Name       string `validate:"required,min=1,max=100" json:"name"`
	CategoryId string `validate:"required,uuid4" json:"categoryId"`
	Price      int64  `validate:"required,min=1000" json:"price"`
}

type ProductUpdateRequest struct {
	Id   string `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type ProductResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type ProductListRequest struct {
	Keyword    string `query:"q"`
	CategoryId string `query:"category_id"`
}

func ToProductRersponse(Product domain.Product) ProductResponse {
	return ProductResponse{
		Id:    Product.BaseModel.ID,
		Name:  Product.ProductName,
		Price: Product.Price,
	}
}

func ToProductRersponses(categories []*domain.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range categories {
		productResponses = append(productResponses, ToProductRersponse(*product))
	}

	return productResponses
}
