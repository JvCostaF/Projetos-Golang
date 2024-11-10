package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllUsersHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "List All Users",
	})
}
