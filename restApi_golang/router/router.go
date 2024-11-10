package router

import "github.com/gin-gonic/gin" //Na pratica, vamos utilizar o pacote com o nome do ultimo diretorio do Path

func Initialize() { //Qualquer entidade que comeca com letra maiuscula sera exportada do package!
	//Initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
