package router

import (
	"github.com/JvCostaF/Projetos-Golang/tree/main/restApi_golang/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.ShowUserHandler)

		v1.POST("/createUser", handler.CreateUserHandler)

		v1.DELETE("/deleteUser", handler.DeleteUserHandler)

		v1.PUT("/updateUser", handler.UpdateUserHandler)

		v1.GET("/users", handler.ListAllUsersHandler)
	}
}
