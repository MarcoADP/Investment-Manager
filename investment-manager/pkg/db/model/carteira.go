package model

type Carteira struct {
	ID                uint    `gorm:"column:carteira_id;primaryKey;autoIncrement"`
	Nome              string  `gorm:"column:nome"`
	Descricao         string  `gorm:"column:descricao"`
	ProporcaoDesejada float64 `gorm:"column:proporcao_desejada"`
}

func (Carteira) TableName() string {
	return "carteira"
}

func NewCarteira(nome string, descricao string, proporcaoDesejada float64) *Carteira {
	return &Carteira{
		Nome:              nome,
		Descricao:         descricao,
		ProporcaoDesejada: proporcaoDesejada,
	}
}
