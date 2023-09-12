package models

import "time"

type Gearbox struct {
	BaseModel
	Name string `gorm:"size:15;type:string;not null"`
}

type CarType struct {
	BaseModel
	Name      string `gorm:"size:50;type:string;not null"`
	CarModels []CarModel
}

type Company struct {
	BaseModel
	Name      string  `gorm:"size:15;type:string;not null"`
	Country   Country `gorm:"foreignKey:CountryID"`
	CountryID int
	CarModels []CarModel
}

type CarModel struct {
	BaseModel
	Name               string  `gorm:"size:15;type:string;not null;unique"`
	Company            Company `gorm:"foreignKey:CompanyId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CompanyId          int
	CarType            CarType `gorm:"foreignKey:CarTypeId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CarTypeId          int
	Gearbox            Gearbox `gorm:"foreignKey:GearboxId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	GearboxId          int
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelProperties []CarModelProperty
	CarModelFiles      []CarModelFile
	CarModelComments   []CarModelComment
}

type CarModelColor struct { // many to many
	BaseModel
	Color      Color `gorm:"foreignKey:ColorId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	ColorId    int
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CarModelId int
}

type CarModelYear struct {
	BaseModel
	CarModel      CarModel    `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId    int         `gorm:"uniqueIndex:idx_CarModelId_PersianYearId"`
	PersianYear   PersianYear `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PersianYearId int         `gorm:"uniqueIndex:idx_CarModelId_PersianYearId"`
	CarModelPrice []CarModelPrice
}

type CarModelFile struct {
	BaseModel
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CarModelId  int
	File        File `gorm:"foreignKey:FileId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	FileId      int
	IsMainImage bool
}

type CarModelPrice struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CarModelId int
	Message    string `gorm:"size:1000;type:string;not null"`
	User       User   `gorm:"foreignKey:UserId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	UserId     int
}
