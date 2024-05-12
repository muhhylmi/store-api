package web

import "github.com/muhhylmi/store-api/model/domain"

type CategoryCreateRequest struct {
	CategoryName string `validate:"required,min=1,max=100" json:"categoryName"`

	AuthData
}

type CategoryResponse struct {
	Id           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

func ToCategoryRersponse(category domain.Categories) CategoryResponse {
	return CategoryResponse{
		Id:           category.BaseModel.ID,
		CategoryName: category.CategoryName,
	}
}

func ToCategoryRersponses(categories []*domain.Categories) []CategoryResponse {
	var categorieResponses []CategoryResponse
	for _, category := range categories {
		categorieResponses = append(categorieResponses, ToCategoryRersponse(*category))
	}

	return categorieResponses
}
