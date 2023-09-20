package main

import (
	"time"
)

type UrlInfo struct {
	id        int
	longUrl   string
	shortUrl  string
	createdBy string
	createdAt time.Time
}
type CreateUrlInfoParams struct {
	longUrl   string
	shortUrl  string
	createdBy string
}

func (urlInfo *UrlInfo) create(params *CreateUrlInfoParams) *UrlInfo {
	return &UrlInfo{
		longUrl:   params.longUrl,
		shortUrl:  params.shortUrl,
		createdBy: params.createdBy,
		createdAt: time.Now(),
	}
}
