package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)

type Base struct{
	ID string `gorm: "type:uuid:primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB)(err error){
	b.ID = uuid.New().String()
	return
}
