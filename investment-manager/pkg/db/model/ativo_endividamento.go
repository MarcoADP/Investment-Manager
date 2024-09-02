package model

import (
	"time"
)

type AtivoEndividamento struct {
	ID                        uint      `gorm:"column:ativo_endividamento_id;primaryKey;autoIncrement"`
	DataCalculo               time.Time `gorm:"column:data_calculo;type:date"`
	Codigo                    string    `gorm:"column:codigo"`
	Divida_Patrimonio_Liquido float64   `gorm:"column:divida_patrimonio_liquido"`
	Divida_Ebit               float64   `gorm:"column:divida_ebit"`
	Divida_Ebitda             float64   `gorm:"column:divida_ebitda"`
}

func (AtivoEndividamento) TableName() string {
	return "ativo_endividamento"
}

func NewAtivoEndividamento(ativoInformacao AtivoInformacao) *AtivoEndividamento {
	return &AtivoEndividamento{
		DataCalculo:               ativoInformacao.DataInformacao,
		Codigo:                    ativoInformacao.Codigo,
		Divida_Patrimonio_Liquido: roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.PatrimonioLiquido, 2.0),
		Divida_Ebit:               roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.Ebit, 2.0),
		Divida_Ebitda:             roundNumbers(ativoInformacao.DividaLiquida/ativoInformacao.Ebitda, 2.0),
	}
}
