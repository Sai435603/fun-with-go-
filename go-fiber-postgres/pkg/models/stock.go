package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Stock struct {
	StockId string  `gorm:"primary_key;column:stock_id" json:"stock-id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price,string"`
	Company string  `json:"company"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (s *Stock) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("StockId", uuid.New().String())
}