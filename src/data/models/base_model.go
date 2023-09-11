package models

import (
	"database/sql"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int           `gorm:"primaryKey"`
	CreatedAt time.Time     `gorm:"type:TIMESTAMP with time zone;not null"`
	UpdatedAt sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	CreatedBy int           `gorm:"not null"`
	UpdatedBy sql.NullInt64 `gorm:"null"`
	DeletedBy sql.NullInt64 `gorm:"null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	var UserId = -1
	if value != nil {
		UserId = int(value.(float64))
	}
	b.CreatedBy = UserId
	b.CreatedAt = time.Now()
	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	var UserId = &sql.NullInt64{
		Valid: false,
	}
	if value != "" {
		UserId = &sql.NullInt64{
			Valid: true,
			Int64: int64(value.(float64)),
		}
	}
	b.UpdatedBy = *UserId
	b.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	return nil
}
