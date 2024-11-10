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
