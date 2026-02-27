package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModels struct{
	Id int `gorm:"primaryKey"`

	CreateAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeleteAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreateBy int `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeleteBy *sql.NullInt64 `gorm:"null"`
}

func(h *BaseModels)BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	var userId = -1

	if value != nil{
		userId = int(value.(float64))
	}

	h.CreateAt = time.Now().UTC()
	h.CreateBy = userId
	return
}

func(h *BaseModels)BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	var userId = sql.NullInt64{Valid: false}

	if value != nil{
		userId = sql.NullInt64{Valid: true, Int64: int64(value.(float64))} 
	}
	
	h.ModifiedAt = sql.NullTime{
		Time: time.Now().UTC(),
		Valid: true,
	}
	h.ModifiedBy = &userId
	return
}

func(h *BaseModels)BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	var userId = sql.NullInt64{Valid: false}

	if value != nil{
		userId = sql.NullInt64{Valid: true, Int64: int64(value.(float64))} 
	}
	
	h.DeleteAt = sql.NullTime{
		Time: time.Now().UTC(),
		Valid: true,
	}
	h.DeleteBy = &userId
	return
}