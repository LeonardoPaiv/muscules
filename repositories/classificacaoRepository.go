package repositories

import (
	"muscules/initializers"
	"muscules/models"
)

func ReadALL() []models.MusculoAtivo {
	var musculo_ativo []models.MusculoAtivo
	initializers.DB.Preload("Classificacao").Find(&musculo_ativo)
	return musculo_ativo
}
