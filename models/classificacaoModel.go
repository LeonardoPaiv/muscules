package models

type Classificacao struct {
	ID   uint   `gorm:"primaryKey;column:id_classificacao"`
	Nome string `gorm:"not null;column:nome"`
}

func (Classificacao) TableName() string {
	return "classificacao"
}
