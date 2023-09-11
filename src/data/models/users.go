package models

import "database/sql"

type User struct {
	BaseModel
	FirstName   string         `gorm:"size:20;type:string;not null"`
	LastName    sql.NullString `gorm:"size:30;type:string;null"`
	Username    string         `gorm:"size:50;type:string;not null;unique"`
	Email       sql.NullString `gorm:"size:150;type:string;null;unique"`
	Password    string         `gorm:"size:64;type:string;not null"`
	Enable      bool           `gorm:"default:true"`
	PhoneNumber sql.NullString `gorm:"size:11;type:string;null;unique;default:null"`
	UserRoles   []UserRole
}

type Role struct {
	BaseModel
	Name      string `gorm:"size:20;type:string;not null;unique"`
	UserRoles []UserRole
}

type UserRole struct {
	User   User `gorm:"foreignKey:UserId"`
	Role   Role `gorm:"foreignKey:RoleId"`
	UserId int
	RoleId int
}
