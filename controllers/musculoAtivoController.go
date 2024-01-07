package controllers

import (
	"muscules/initializers"
	"muscules/models"

	"github.com/gin-gonic/gin"
)

func MusculoAtivoRead(c *gin.Context) {
	var musculo_ativo []models.MusculoAtivo
	initializers.DB.Preload("Classificacao").Find(&musculo_ativo)

	c.JSON(200, gin.H{
		"musculo_ativo": musculo_ativo,
	})

}
