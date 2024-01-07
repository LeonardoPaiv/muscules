package models

type MusculoAtivo struct {
	ID               uint          `gorm:"primaryKey;column:id_musculo_ativo"`
	Nome             string        `gorm:"not null;column:nome"`
	CodClassificacao int           `gorm:"not null;column:cod_classificacao"`
	Classificacao    Classificacao `gorm:"foreignKey:CodClassificacao"`
}

func (MusculoAtivo) TableName() string {
	return "musculo_ativo"
}
