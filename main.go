package main

import (
	"muscules/controllers"
	"muscules/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/classificacao", controllers.ClassificacaoRead)
	r.GET("/musculo-ativo", controllers.MusculoAtivoRead)
	r.Run()
}
