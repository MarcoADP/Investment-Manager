package model

import (
	"time"
)

type AtivoEndividamento struct {
	ID                      uint      `gorm:"column:ativo_endividamento_id;primaryKey;autoIncrement"`
	AtivoInformacaoId       uint      `gorm:"column:ativo_informacao_id"`
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
	dividaEbitda := 0.0
	if ativoInformacao.Ebitda != 0 {
		dividaEbitda = ativoInformacao.DividaLiquida / ativoInformacao.Ebitda
	}
	return &AtivoEndividamento{
		AtivoInformacaoId:       ativoInformacao.ID,
		DataCalculo:             ativoInformacao.DataInformacao,
		Codigo:                  ativoInformacao.Codigo,
		DividaPatrimonioLiquido: roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.PatrimonioLiquido, 2.0),
		DividaEbit:              roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.Ebit, 2.0),
		DividaEbitda:            roundNumbers(float64(dividaEbitda), 2.0),
	}
}
