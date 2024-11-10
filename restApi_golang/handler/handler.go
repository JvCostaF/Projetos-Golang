package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET User",
	})
}

func CreateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}

func DeleteUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Delete User",
	})
}

func UpdateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update User",
	})
}

func ListAllUsersHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "List All Users",
	})
}
