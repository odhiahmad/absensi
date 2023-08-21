package dto

import uuid "github.com/satori/go.uuid"

type AbsenUpdateDTO struct {
	ID       uuid.UUID
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Lat     string `json:"lat" form:"lat" binding:"required"`
	Long     string `json:"long" form:"long" binding:"required"`
	UserID   uuid.UUID `gorm:"null"`

}

type AbsenCreateDTO struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Lat     string `json:"lat" form:"lat" binding:"required"`
	Long     string `json:"long" form:"long" binding:"required"`
	UserID   uuid.UUID `gorm:"null"`
}
