package handlers

import (
	"fmt"
	"net/http"
	"time"
	"url-shortener-service/db"
	"url-shortener-service/models"
	"url-shortener-service/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterHandler(c *gin.Context) {
	var userRequest models.RegisterUserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel := getUserModel(userRequest.Email)
	if userModel != nil {
		c.JSON(http.StatusOK, gin.H{"response": "User already registered. Please login."})
		return
	}
	userModel = createUserModel(userModel, &userRequest)
	insertUserTrx := db.Insert(userModel)
	if insertUserTrx.Error != nil {
		fmt.Println("LoginHandler: ", insertUserTrx.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "User registered successfully"})
}
func TokenHandler(c *gin.Context) {
	{
		email := c.Query("email")
		pass := c.Query("password")
		fmt.Println(email)
		fmt.Println(pass)
		user := getUserModel(email)
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found. Please register."})
			return
		}

		errHash := utils.CompareHashPassword(pass, user.Password)

		if !errHash {
			c.JSON(400, gin.H{"error": "invalid password"})
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			StandardClaims: jwt.StandardClaims{
				Subject:   user.Email,
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		jwtKey := viper.GetString("jwtKey")
		tokenString, err := token.SignedString([]byte(jwtKey))

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{"error": "could not generate token"})
			return
		}
		c.JSON(200, gin.H{"Token": tokenString})
	}
}
func createUserModel(userModel *models.UserInfo, userRequest *models.RegisterUserRequest) *models.UserInfo {
	passwordHash, err := utils.GenerateHashPassword(userRequest.Password)
	if err != nil {
		panic(err.Error())
	}
	userModel = &models.UserInfo{
		Email:    userRequest.Email,
		Password: passwordHash,
	}
	return userModel
}

func getUserModel(email string) *models.UserInfo {
	getUserTrx := db.GetRecordByValue("user_infos", "email", email)
	var userModel *models.UserInfo
	if getUserTrx.Error == nil {
		tx := getUserTrx.First(&userModel)
		if tx.Error == nil {
			return userModel
		}
	}
	return nil
}
