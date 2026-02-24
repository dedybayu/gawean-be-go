package models

import "time"

type Level struct {
	ID        uint      `gorm:"primaryKey"`
	NamaLevel string    `gorm:"size:50;not null"`
	Kode      string    `gorm:"size:10;uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
