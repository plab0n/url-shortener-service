package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"url-shortener-service/models"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {
	var shortenUrlRequest models.ShortenUrlRequest

	if err := c.Bind(&shortenUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := getShortUrl(&shortenUrlRequest)
	//TODO save in db
	c.JSON(http.StatusOK, gin.H{"shortUrl": shortUrl})
	return
}

func getShortUrl(req *models.ShortenUrlRequest) string {
	hash := calculateSha256(req.LongUrl)

	len := 6
	//TODO: Length can be configurable
	host := "http://localhost:8080/"
	//TODO: HotName can be configurable
	sUrl := host + hash[:len]
	if isExist(sUrl) {
		//TODO: re-calculate hash with a salt and return the sUrl
	}

	return sUrl
}

func calculateSha256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)
	return hashHex
}
func isExist(sUrl string) bool {
	//TODO: check in db if sUrl exists
	return false
}
