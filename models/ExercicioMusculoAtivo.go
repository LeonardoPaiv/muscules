package models

type ExercicioMusculoAtivo struct {
	ID              uint         `gorm:"primaryKey;column:id_exercicio_musculo_ativo"`
	CodExercicio    int          `gorm:"not null;column:cod_exercicio"`
	CodMusculoAtivo int          `gorm:"not null;column:cod_musculo_ativo"`
	Exercicio       Exercicio    `gorm:"foreignKey:CodExercicio"`
	MusculoAtivo    MusculoAtivo `gorm:"foreignKey:CodMusculoAtivo"`
}

func (ExercicioMusculoAtivo) TableName() string {
	return "exercicio_musculo_ativo"
}
