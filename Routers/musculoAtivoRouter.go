package routers

import (
	"muscules/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersMusculoAtivo(r *gin.RouterGroup) {
	r.GET("", controllers.MusculoAtivoRead)
}
