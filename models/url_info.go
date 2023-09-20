package main

type UrlInfo struct {
	id        int
	longUrl   string
	shortUrl  string
	createdBy string
	createdAt string
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
	}
}
