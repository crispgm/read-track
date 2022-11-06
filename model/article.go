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

	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	Domain      string     `json:"domain,omitempty"`
	Author      string     `json:"author,omitempty"`
	Type        string     `json:"type,omitempty"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
}
