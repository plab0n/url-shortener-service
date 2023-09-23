package handlers

import (
	"fmt"
	"net/http"
	"url-shortener-service/db"
	"url-shortener-service/models"
	"url-shortener-service/utils"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var userRequest models.RegisterUserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel := getUserModel(userRequest)
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

func getUserModel(user models.RegisterUserRequest) *models.UserInfo {
	getUserTrx := db.GetRecordByValue("user_infos", "email", user.Email)
	var userModel *models.UserInfo
	if getUserTrx.Error == nil {
		tx := getUserTrx.First(&userModel)
		if tx.Error == nil {
			return userModel
		}
	}
	return nil
}
