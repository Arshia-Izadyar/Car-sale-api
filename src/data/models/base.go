package models

type Country struct {
	BaseModel
	Name   string `gorm:"size:30;not null;type:string;unique"`
	Cities []City
}

type City struct {
	BaseModel
	Name      string `gorm:"type:string;not null;unique;size:40"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

type PersianYear struct {
	BaseModel
	Name string `gorm:"type:string;not null;unique;size:15"`
	Year int    `gorm:"not null;unique;size:5"`
}

type Color struct {
	BaseModel
	Name string `gorm:"type:string;size:30;not null;unique"`
	Hex  string `gorm:"size:7;type:string;not null"`
}
