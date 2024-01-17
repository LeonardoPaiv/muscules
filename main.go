package main

import (
	routers "muscules/Routers"
	"muscules/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	classificacao := r.Group("classificacao")
	musculesAtivo := r.Group("musculo-ativo")
	routers.RoutersClassificacao(classificacao)
	routers.RoutersMusculoAtivo(musculesAtivo)
	r.Run()
}
