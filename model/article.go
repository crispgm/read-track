package model

import (
	"time"
)

// Read types
const (
	RTRead = "read"
	RTSkim = "skim"
	RTSkip = "skip"
)

// Article definition of an article
type Article struct {
	Title       string
	Description string
	URL         string
	Domain      string
	Author      string
	Type        string
	ReadAt      *time.Time
	FinishedAt  *time.Time
}
