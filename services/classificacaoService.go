package services

import (
	"muscules/models"
	"muscules/repositories"
)

func ReadALL() []models.MusculoAtivo {
	return repositories.ReadALL()
}
