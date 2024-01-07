package controllers

import (
	"muscules/services"

	"github.com/gin-gonic/gin"
)

func ClassificacaoRead(c *gin.Context) {
	c.JSON(200, gin.H{
		"Classificacao": services.ReadALL(),
	})

}
