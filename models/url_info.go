package models

import (
	"time"
)

type UrlInfo struct {
	Id         int
	LongUrl    string
	ShortUrl   string
	CreatedBy  string
	CcreatedAt time.Time
}
type CreateUrlInfoParams struct {
	longUrl   string
	shortUrl  string
	createdBy string
}

// func (urlInfo *UrlInfo) New(params *CreateUrlInfoParams) *UrlInfo {
// 	return &UrlInfo{
// 		longUrl:   params.longUrl,
// 		shortUrl:  params.shortUrl,
// 		createdBy: params.createdBy,
// 		createdAt: time.Now(),
// 	}
// }
