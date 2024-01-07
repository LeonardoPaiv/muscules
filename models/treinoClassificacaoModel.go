package models

type TreinoClassificacao struct {
	ID               uint          `gorm:"primaryKey;column:id_treino_classificacao"`
	CodTreino        int           `gorm:"not null;column:cod_treino"`
	CodClassificacao int           `gorm:"not null;column:cod_classificacao"`
	Treino           Treino        `gorm:"foreignKey:CodTreino"`
	Classificacao    Classificacao `gorm:"foreignKey:CodClassificacao"`
}

func (TreinoClassificacao) TableName() string {
	return "treino_classificacao"
}
