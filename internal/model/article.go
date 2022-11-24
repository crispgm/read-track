// Package model .
package model

import (
	"errors"
	"html"
	"net/url"
	"time"

	"gorm.io/gorm"
)

// Read types
const (
	RTRead   = "read"   // have read
	RTSkim   = "skim"   // have skimmed
	RTUnread = "unread" // unread: 1. from read to unread; 2. read it later
	RTSkip   = "skip"   // skipped and will not read
)

// Model errors
var (
	ErrDuplicatedQuery = errors.New("Duplicated record")
)

// Article definition of an article
type Article struct {
	gorm.Model

	ReadType    string `gorm:"read_type;size:16;not null" json:"Type,omitempty"`
	Title       string `gorm:"title;size:128;not null" json:"Title,omitempty" yaml:"title"`
	URL         string `gorm:"url;size:512;not null" json:"URL,omitempty"`
	Domain      string `gorm:"domain;index;size:256;not null" json:"Domain,omitempty"`
	Author      string `gorm:"author;index;size:128" json:"Author,omitempty"`
	Description string `gorm:"description;size:256" json:"Description,omitempty"`
	Device      string `gorm:"device;index;size:32" json:"Device,omitempty"`

	ReadAtText      string `gorm:"-" json:"CreatedAtText"`
	DescriptionText string `gorm:"-" json:"DescriptionText"`
}

// CreateArticle .
func CreateArticle(db *gorm.DB, a *Article) error {
	urlp, err := url.Parse(a.URL)
	if err != nil {
		return err
	}
	a.Domain = urlp.Hostname()
	var existed Article
	err = db.Where("url = ?", a.URL).Take(&existed).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result := db.Create(a)
		return result.Error
	}

	existed.ReadType = a.ReadType
	existed.Title = html.EscapeString(a.Title)
	existed.Description = html.EscapeString(a.Description)
	existed.Author = html.EscapeString(a.Author)
	existed.Device = html.EscapeString(a.Device)
	return db.Save(&existed).Error
}

// ListArticles .
func ListArticles(db *gorm.DB, loc *time.Location, year, month int) ([]Article, error) {
	var articles []Article
	yearMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc).Format("2006-01")
	err := db.
		Where("strftime('%Y-%m', updated_at, 'localtime') = ?", yearMonth).
		Order("updated_at DESC").
		Find(&articles).
		Error
	return articles, err
}

// ArticleExport for export
type ArticleExport struct {
	ReadType    string    `yaml:"read_type"`
	Title       string    `yaml:"title"`
	URL         string    `yaml:"url"`
	Domain      string    `yaml:"domain"`
	Author      string    `yaml:"author,omitempty"`
	Description string    `yaml:"description,omitempty"`
	Device      string    `yaml:"device,omitempty"`
	CreatedAt   time.Time `yaml:"created_at"`
	UpdatedAt   time.Time `yaml:"updated_at"`
}

// ExportArticles .
func ExportArticles(db *gorm.DB, loc *time.Location, year, month int) ([]Article, error) {
	var articles []Article
	yearMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc).Format("2006-01")
	err := db.
		Where("strftime('%Y-%m', updated_at, 'localtime')  = ?", yearMonth).
		Order("updated_at DESC").
		Find(&articles).
		Error
	return articles, err
}

// ArticleStat .
type ArticleStat struct {
	Name string
	All  int64
	Read int64
	Skim int64
	Skip int64
}

// ArticleStatResult .
type ArticleStatResult struct {
	ReadType  string
	ReadCount int64
}

// GetArticleStatistics .
func GetArticleStatistics(db *gorm.DB, loc *time.Location) ([]ArticleStat, error) {
	now := time.Now()

	// today
	today, err := getArticleStatistics(db, "Today", "%Y-%m-%d", now.Format("2006-01-02"))

	// this month
	thisMonth, err := getArticleStatistics(db, "This Month", "%Y-%m", now.Format("2006-01"))

	stats := []ArticleStat{today, thisMonth}
	return stats, err
}

func getArticleStatistics(db *gorm.DB, name string, dateFormat, dateCondition string) (ArticleStat, error) {
	var result []ArticleStatResult
	err := db.
		Table("Articles").
		Select("read_type", "count(id) AS read_count").
		Where("strftime(?, updated_at, 'localtime') = ?", dateFormat, dateCondition).
		Group("read_type").
		Scan(&result).
		Error

	var stat ArticleStat
	stat.Name = name
	for _, item := range result {
		stat.All += item.ReadCount
		if item.ReadType == RTRead {
			stat.Read = item.ReadCount
		} else if item.ReadType == RTSkim {
			stat.Skim = item.ReadCount
		} else if item.ReadType == RTSkip {
			stat.Skip = item.ReadCount
		}
	}
	return stat, err
}
