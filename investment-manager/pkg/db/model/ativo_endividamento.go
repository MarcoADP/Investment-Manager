package model

import (
	"time"
)

type AtivoEndividamento struct {
	ID                      uint      `gorm:"column:ativo_endividamento_id;primaryKey;autoIncrement"`
	DataCalculo             time.Time `gorm:"column:data_calculo;type:date"`
	Codigo                  string    `gorm:"column:codigo"`
	DividaPatrimonioLiquido float64   `gorm:"column:divida_patrimonio_liquido"`
	DividaEbit              float64   `gorm:"column:divida_ebit"`
	DividaEbitda            float64   `gorm:"column:divida_ebitda"`
}

func (AtivoEndividamento) TableName() string {
	return "ativo_endividamento"
}

func NewAtivoEndividamento(ativoInformacao AtivoInformacao) *AtivoEndividamento {
	return &AtivoEndividamento{
		DataCalculo:             ativoInformacao.DataInformacao,
		Codigo:                  ativoInformacao.Codigo,
		DividaPatrimonioLiquido: roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.PatrimonioLiquido, 2.0),
		DividaEbit:              roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.Ebit, 2.0),
		DividaEbitda:            roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.Ebitda, 2.0),
	}
}
