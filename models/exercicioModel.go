package models

type Exercicio struct {
	ID   uint   `gorm:"primaryKey;column:id_exercicio"`
	Nome string `gorm:"not null;column:nome"`
}

func (Exercicio) TableName() string {
	return "exercicio"
}
