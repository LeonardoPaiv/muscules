package models

type Metodo struct {
	ID        uint   `gorm:"primaryKey;column:id_metodo"`
	Nome      string `gorm:"not null;column:nome"`
	Descricao string `gorm:"not null;column:descricao"`
}

func (Metodo) TableName() string {
	return "metodo"
}
