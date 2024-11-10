package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update User",
	})
}
