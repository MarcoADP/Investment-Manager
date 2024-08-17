package model

type CarteiraAtivo struct {
	ID                uint    `gorm:"column:carteira_ativo_id;primaryKey;autoIncrement"`
	CarteiraId        uint    `gorm:"column:carteira_id"`
	Codigo            string  `gorm:"column:codigo"`
	ProporcaoDesejada float64 `gorm:"column:proporcao_desejada"`
	Movimento         string  `gorm:"column:movimento"`
}

func (CarteiraAtivo) TableName() string {
	return "carteira_ativo"
}

func NewCarteiraAtivo(carteiraId uint, codigo string, proporcaoDesejada float64, movimento string) *CarteiraAtivo {
	return &CarteiraAtivo{
		CarteiraId:        carteiraId,
		Codigo:            codigo,
		ProporcaoDesejada: proporcaoDesejada,
		Movimento:         movimento,
	}
}
