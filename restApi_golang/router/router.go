package router

import "github.com/gin-gonic/gin" //Na pratica, vamos utilizar o pacote com o nome do ultimo diretorio do Path

func Initialize() { //Qualquer entidade que comeca com letra maiuscula sera exportada do package!
	//Inicializa o router com as configuracoes Default do gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
