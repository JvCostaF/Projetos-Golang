package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}
