package model

import (
	"database/sql"
	"errors"
	"net/url"

	"gorm.io/gorm"
)

// Read types
const (
	RTRead = "read"
	RTSkim = "skim"
	RTSkip = "skip"

	StatusOK = 0
)

// Model errors
var (
	ErrDuplicatedQuery = errors.New("Duplicated record")
)

// Article definition of an article
type Article struct {
	gorm.Model

	Title       string         `gorm:"title;size:128;not null" json:"Title,omitempty"`
	Description string         `gorm:"description;size:256;not null" json:"Description,omitempty"`
	URL         string         `gorm:"url;size:512;not null" json:"URL,omitempty"`
	Domain      string         `gorm:"domain;index;size:256;not null" json:"Domain,omitempty"`
	Author      sql.NullString `gorm:"author;index;size:128;not null" json:"Author,omitempty"`
	ReadType    string         `gorm:"read_type;size:16;not null" json:"Type,omitempty"`
	ReadAt      sql.NullTime   `gorm:"read_at" json:"ReadAt,omitempty"`
	Status      uint8          `gorm:"status;not null" json:"Status"`
}

// CreateArticle .
func CreateArticle(db *gorm.DB, a *Article) error {
	urlp, err := url.Parse(a.URL)
	if err != nil {
		return err
	}
	a.Domain = urlp.Hostname()
	err = db.Where("url = ? AND status = ?", a.URL, StatusOK).Take(&Article{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.Status = StatusOK
		result := db.Create(a)
		return result.Error
	}

	return ErrDuplicatedQuery
}
