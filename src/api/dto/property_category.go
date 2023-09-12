package dto

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type UpdatePropertyCategoryRequest struct {
	Name string `json:"name"`
}

type PropertyCategoryResponse struct {
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Icon       string             `json:"icon"`
	Properties []PropertyResponse `json:"properties,omitempty"`
}

type CreatePropertyRequest struct {
	Description string `json:"description"`
	Name        string `json:"name" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Unit        string `json:"unit"`
	DataType    string `json:"dataType" binding:"required"`
	CategoryId  int    `json:"categoryId" binding:"required"`
}

type UpdatePropertyRequest struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Unit        string `json:"unit,omitempty"`
	DataType    string `json:"dataType,omitempty"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name,omitempty" `
	Icon        string                   `json:"icon,omitempty"`
	Description string                   `json:"description,omitempty"`
	DataType    string                   `json:"data_type,omitempty"`
	Unit        string                   `json:"unit,omitempty"`
	Category    PropertyCategoryResponse `json:"category,omitempty"`
}
