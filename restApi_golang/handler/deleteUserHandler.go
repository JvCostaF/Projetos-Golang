package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Delete User",
	})
}
