package dto

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type UpdatePropertyCategoryRequest struct {
	Name string `json:"name"`
}

type PropertyCategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
