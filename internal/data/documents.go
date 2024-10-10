package data

import (
	"time"
)

type Document struct {
	ID             uint64 `gorm:"primarykey;"`
	Title          string
	Author         string
	YearPublished  time.Time
	ISBN           uint64
	LibraryID      uint64 `gorm:"foreignKey:LibraryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DocumentTypeID uint   `gorm:"foreignKey:DocumentTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
