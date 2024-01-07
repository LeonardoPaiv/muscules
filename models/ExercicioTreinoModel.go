package models

type ExercicioTreino struct {
	ID           uint      `gorm:"primaryKey;column:id_exercicio_treino"`
	CodExercicio int       `gorm:"not null;column:cod_exercicio"`
	CodMetodo    int       `gorm:"not null;column:cod_metodo"`
	Repeticoes   int       `gorm:"not null;column:repeticoes"`
	Descanso     int       `gorm:"not null;column:descanso"`
	Exercicio    Exercicio `gorm:"foreignKey:CodExercicio"`
	Metodo       Metodo    `gorm:"foreignKey:CodMetodo"`
}

func (ExercicioTreino) TableName() string {
	return "exercicio_treino"
}
