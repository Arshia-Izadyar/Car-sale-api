package dto

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
