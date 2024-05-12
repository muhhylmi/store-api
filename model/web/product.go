package web

import "github.com/muhhylmi/store-api/model/domain"

type ProductCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type ProductUpdateRequest struct {
	Id   string `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type ProductResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ToProductRersponse(Product domain.Product) ProductResponse {
	return ProductResponse{
		Id:   Product.Id,
		Name: Product.Name,
	}
}

func ToProductRersponses(categories []domain.Product) []ProductResponse {
	var ProductResponses []ProductResponse
	for _, Product := range categories {
		ProductResponses = append(ProductResponses, ToProductRersponse(Product))
	}

	return ProductResponses
}
