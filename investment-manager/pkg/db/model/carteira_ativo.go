package model

type CarteiraAtivo struct {
	ID                uint    `gorm:"column:carteira_ativo_id;primaryKey;autoIncrement"`
	CarteiraId        uint    `gorm:"column:carteira_id"`
	Codigo            string  `gorm:"column:codigo"`
	ProporcaoDesejada float64 `gorm:"column:proporcao_desejada"`
}

func (CarteiraAtivo) TableName() string {
	return "carteira_ativo"
}

func NewCarteiraAtivo(carteiraId uint, codigo string, proporcaoDesejada float64) *CarteiraAtivo {
	return &CarteiraAtivo{
		CarteiraId:        carteiraId,
		Codigo:            codigo,
		ProporcaoDesejada: proporcaoDesejada,
	}
}
