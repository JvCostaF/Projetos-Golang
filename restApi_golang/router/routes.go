package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{ //Converte conteudo dentro do bloco em JSON
				"name":   "Joao Victor da Costa Farias",
				"age":    "24",
				"status": "Single",
			})
		})

		v1.POST("/createUser", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{ //Converte conteudo dentro do bloco em JSON
				"name":   "Joao Victor da Costa Farias",
				"age":    "24",
				"status": "Single",
			})
		})

		v1.DELETE("/deleteUser", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{ //Converte conteudo dentro do bloco em JSON
				"name":   "Joao Victor da Costa Farias",
				"age":    "24",
				"status": "Single",
			})
		})

		v1.PUT("/updateUser", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{ //Converte conteudo dentro do bloco em JSON
				"name":   "Joao Victor da Costa Farias",
				"age":    "24",
				"status": "Single",
			})
		})

		v1.GET("/users", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{ //Converte conteudo dentro do bloco em JSON
				"name":   "Joao Victor da Costa Farias",
				"age":    "24",
				"status": "Single",
			})
		})
	}
}
