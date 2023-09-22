package models

import (
	"gorm.io/gorm"
)

type UrlInfo struct {
	*gorm.Model
	LongUrl   string
	ShortUrl  string `gorm:"index"`
	CreatedBy string
}

// type CreateUrlInfoParams struct {
// 	longUrl   string
// 	shortUrl  string
// 	createdBy string
// }

// func (urlInfo *UrlInfo) New(params *CreateUrlInfoParams) *UrlInfo {
// 	return &UrlInfo{
// 		longUrl:   params.longUrl,
// 		shortUrl:  params.shortUrl,
// 		createdBy: params.createdBy,
// 		createdAt: time.Now(),
// 	}
// }
