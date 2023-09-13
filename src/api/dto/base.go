package dto

import "mime/multipart"

type CreateColorRequest struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type UpdateColorRequest struct {
	Name string `json:"name,omitempty"`
	Hex  string `json:"hex,omitempty"`
}

type ColorResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type CreatePersianYearRequest struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type UpdatePersianYearRequest struct {
	Name string `json:"name"`
}

type PersianYearResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type CreateCountryRequest struct {
	Name string `json:"name"`
}

type UpdateCountryRequest struct {
	Name string `json:"name"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities"`
}

type CreateCityRequest struct {
	Name      string `json:"name" binding:"required"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityRequest struct {
	Name string `json:"name" binding:"required"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
}

type FileFormRequest struct {
	File *multipart.FileHeader `form:"file" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MineType    string `json:"mineType"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MineType    string `json:"mineType"`
}

type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"max=15"`
	CountryID int    `json:"countryId" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name string `json:"name" binding:"max=15"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

type CreateGearboxRequest struct {
	Name string `json:"name"`
}

type UpdateGearboxRequest struct {
	Name string `json:"name"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
