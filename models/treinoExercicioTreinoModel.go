package models

type TreinoExercicioTreino struct {
	ID                 uint            `gorm:"primaryKey;column:id_treino_exercicio_treino"`
	CodTreino          int             `gorm:"not null;column:cod_treino"`
	CodExercicioTreino int             `gorm:"not null;column:cod_exercicio_treino"`
	Treino             Treino          `gorm:"foreignKey:CodTreino"`
	ExercicioTreino    ExercicioTreino `gorm:"foreignKey:CodExercicioTreino"`
}

func (TreinoExercicioTreino) TableName() string {
	return "treino_exercicio_treino "
}
