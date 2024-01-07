package models

type Treino struct {
	ID uint `gorm:"primaryKey;column:id_treino"`
}

func (Treino) TableName() string {
	return "treino"
}
