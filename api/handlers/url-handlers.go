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
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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
func DeleteUrlHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied!"})
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied!"})
		return
	}
	var req models.DeleteUrlRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	filters := make(map[string]interface{})
	filters["created_by"] = claims.Subject
	filters["short_url"] = req.ShortUrl
	queryParams := &models.QueryParams{
		Table:   "url_infos",
		Filters: filters,
	}
	tx := db.GetRecordByQuery(queryParams)

	if tx.Error != nil {

	}
	var urlInfo models.UrlInfo
	tx.First(&urlInfo)
	fmt.Println("Found url: ", urlInfo.ShortUrl)
	tx = db.Delete(&urlInfo)
	if tx.Error != nil {

	}
	c.JSON(http.StatusOK, gin.H{"response": "Deleted successfully"})
}
func getShortUrl(req *models.ShortenUrlRequest) string {
	hash := calculateSha256(req.LongUrl)

	len := viper.GetInt("shortUrlLength")
	sUrl := hash[:len]
	if isExist(sUrl) {
		hash = calculateSha256(req.LongUrl + uuid.New().String())
		sUrl = hash[:len]
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
	filters := make(map[string]interface{})
	filters["short_url"] = sUrl
	queryParam := &models.QueryParams{
		Table:   "url_infos",
		Filters: filters,
	}
	tx := db.GetRecordByQuery(queryParam)
	if tx.Error == gorm.ErrRecordNotFound {
		fmt.Println("Record Not found")
		return false
	}
	var urlInfo models.UrlInfo
	tx.First(&urlInfo)
	return &urlInfo != nil
}
