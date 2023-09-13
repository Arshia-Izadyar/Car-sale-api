package dto

import "time"

type CreateCarTypeRequest struct {
	Name string `json:"name"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	Company            CompanyResponse            `json:"company"`
	CarType            CarTypeResponse            `json:"carType"`
	Gearbox            GearboxResponse            `json:"gearbox"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors"`
	CarModelFiles      []CarModelFileResponse     `json:"carModelFiles"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments"`
}

type CreateCarModelColorRequest struct {
	ColorId    int `json:"colorId" binding:"required"`
	CarModelId int `json:"carModelId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	ColorId    int `json:"colorId,omitempty"`
	CarModelId int `json:"carModelId,omitempty"`
}

type CarModelColorResponse struct {
	Id         int `json:"id"`
	ColorId    int `json:"colorId,omitempty"`
	CarModelId int `json:"carModelId,omitempty"`
}

type CreateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId"`
	PersianYearId int `json:"persianYearId"`
}

type UpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId,omitempty"`
	PersianYearId int `json:"persianYearId,omitempty"`
}

type CarModelYearResponse struct {
	Id            int                     `json:"id"`
	CarModelId    int                     `json:"carModelId"`
	PersianYear   PersianYearResponse     `json:"persianYear"`
	CarModelPrice []CarModelPriceResponse `json:"carModelPrice"`
}

type CreateCarModelPriceRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type UpdateCarModelPriceRequest struct {
	Price   float64   `json:"price,omitempty"`
	PriceAt time.Time `json:"priceAt,omitempty"`
}

type CarModelPriceResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type CreateCarModelFileRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	FileId      int  `json:"fileId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type UpdateCarModelFileRequest struct {
	IsMainImage bool `json:"isMainImage"`
}

type CarModelFileResponse struct {
	Id          int          `json:"id"`
	CarModelId  int          `json:"carModelId"`
	File        FileResponse `json:"file"`
	IsMainImage bool         `json:"isMainImage"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value"`
}

type CarModelPropertyResponse struct {
	CarModelId int              `json:"carModelId"`
	Property   PropertyResponse `json:"property"`
	Value      string           `json:"value"`
}

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId"`
	Message    string `json:"message" binding:"required"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"message"`
}

type CarModelCommentResponse struct {
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
