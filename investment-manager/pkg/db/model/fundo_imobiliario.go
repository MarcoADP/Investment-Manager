package model

type FundoImobiliario struct {
	ID       uint   `gorm:"column:fundo_imobiliario_id;primaryKey"`
	Nome     string `gorm:"column:nome"`
	Codigo   string `gorm:"column:codigo"`
	Tipo     string `gorm:"column:tipo"`
	Segmento string `gorm:"column:segmento"`
	Cnpj     string `gorm:"column:cnpj"`
}

func (FundoImobiliario) TableName() string {
	return "fundo_imobiliario"
}

func NewFundoImobiliario(nome string, codigo string, tipo string, segmento string, cnpj string) *FundoImobiliario {
	return &FundoImobiliario{
		Nome:     nome,
		Codigo:   codigo,
		Tipo:     tipo,
		Segmento: segmento,
		Cnpj:     cnpj,
	}
}
