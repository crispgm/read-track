package model

import (
	"database/sql"
	"errors"
	"net/url"
	"time"

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

	Title       string       `gorm:"title;size:128;not null" json:"Title,omitempty" yaml:"title"`
	URL         string       `gorm:"url;size:512;not null" json:"URL,omitempty"`
	Domain      string       `gorm:"domain;index;size:256;not null" json:"Domain,omitempty"`
	Author      string       `gorm:"author;index;size:128;not null" json:"Author,omitempty"`
	Description string       `gorm:"description;size:256;not null" json:"Description,omitempty"`
	ReadType    string       `gorm:"read_type;size:16;not null" json:"Type,omitempty"`
	ReadAt      sql.NullTime `gorm:"read_at" json:"ReadAt,omitempty"`
	Status      uint8        `gorm:"status;not null" json:"Status"`
}

// ArticleExport for export
type ArticleExport struct {
	Title       string     `yaml:"title"`
	URL         string     `yaml:"url"`
	Domain      string     `yaml:"domain"`
	Author      string     `yaml:"author,omitempty"`
	Description string     `yaml:"description,omitempty"`
	CreatedAt   time.Time  `yaml:"created_at"`
	ReadType    string     `yaml:"read_type"`
	ReadAt      *time.Time `yaml:"read_at"`
}

// CreateArticle .
func CreateArticle(db *gorm.DB, a *Article) error {
	urlp, err := url.Parse(a.URL)
	if err != nil {
		return err
	}
	a.Domain = urlp.Hostname()
	a.ReadAt = sql.NullTime{Time: time.Now(), Valid: true}
	err = db.Where("url = ? AND status = ?", a.URL, StatusOK).Take(&Article{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.Status = StatusOK
		result := db.Create(a)
		return result.Error
	}

	return ErrDuplicatedQuery
}

// ExportArticles .
func ExportArticles(db *gorm.DB, loc *time.Location, year, month int) ([]Article, error) {
	var articles []Article
	fromTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	toTime := fromTime.AddDate(0, 1, 0).Add(time.Second * -1)
	err := db.Where("created_at BETWEEN (? AND ?) AND status = ?", fromTime, toTime, StatusOK).Find(&articles).Error
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
		Where("strftime(?, created_at, 'localtime') = ? AND status = ?", dateFormat, dateCondition, StatusOK).
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
