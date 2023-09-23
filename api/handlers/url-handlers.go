package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"url-shortener-service/db"
	"url-shortener-service/models"
	"url-shortener-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CreateShortUrl(c *gin.Context) {
	var shortenUrlRequest models.ShortenUrlRequest

	if err := c.Bind(&shortenUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := getShortUrl(&shortenUrlRequest)
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)

	//TODO: handle exception
	urlInfo := &models.UrlInfo{
		LongUrl:  shortenUrlRequest.LongUrl,
		ShortUrl: shortUrl,
	}
	if err == nil {
		urlInfo.CreatedBy = claims.Subject
	}
	tx := db.Insert(urlInfo)
	if tx.Error != nil {
		fmt.Println("Data Insert failed ", tx.Error)
	}
	host := viper.GetString("host")
	urlWithHostName := host + shortUrl
	c.JSON(http.StatusOK, gin.H{"shortUrl": urlWithHostName})
	return
}
func RedirectToLongUrl(c *gin.Context) {
	sUrl := c.Param("shortUrl")
	fmt.Println("Found ", sUrl)
	var urlInfo models.UrlInfo
	db.GetItemByValue("short_url", sUrl).First(&urlInfo)
	fmt.Println("Redirecting to ", urlInfo.LongUrl)
	location := url.URL{Path: urlInfo.LongUrl}
	c.Redirect(http.StatusFound, location.RequestURI())
}
func getShortUrl(req *models.ShortenUrlRequest) string {
	hash := calculateSha256(req.LongUrl)

	len := viper.GetInt("shortUrlLength")
	sUrl := hash[:len]
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
