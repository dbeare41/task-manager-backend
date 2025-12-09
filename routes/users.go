package routes

import (
	"my-task-manager/models"
	"my-task-manager/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Unable to parse user"})
		return
	}
	err = user.SaveUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to save user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "User signed up succesfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "unable to parse request"})
		return
	}

	err = user.VerifyUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "unable to verify user"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Could not generate token"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"Message": "logged in successfully", "token": token})

}
