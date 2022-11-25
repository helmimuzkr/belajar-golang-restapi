package model

type Category struct {
	ID   int
	Name string
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=25"`
}
type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=25"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
