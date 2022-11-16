package model

import (
	"time"

	"gorm.io/gorm"
)

// Read types
const (
	RTRead = "read"
	RTSkim = "skim"
	RTSkip = "skip"
)

// Article definition of an article
type Article struct {
	gorm.Model

	Title       string     `json:"Title,omitempty"`
	Description string     `json:"Description,omitempty"`
	URL         string     `json:"URL,omitempty"`
	Domain      string     `json:"Domain,omitempty"`
	Author      string     `json:"Author,omitempty"`
	Type        string     `json:"Type,omitempty"`
	ReadAt      *time.Time `json:"ReadAt,omitempty"`
}

// CreateArticle .
func CreateArticle(db *gorm.DB, a *Article) error {
	result := db.Create(a)

	return result.Error
}
