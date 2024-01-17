package routers

import (
	"muscules/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersClassificacao(r *gin.RouterGroup) {
	r.GET("", controllers.ClassificacaoRead)
}
