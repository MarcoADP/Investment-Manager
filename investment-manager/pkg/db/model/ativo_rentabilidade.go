package model

import (
	"time"
)

type AtivoRentabilidade struct {
	ID          uint      `gorm:"column:ativo_rentabilidade_id;primaryKey;autoIncrement"`
	DataCalculo time.Time `gorm:"column:data_calculo;type:date"`
	Codigo      string    `gorm:"column:codigo"`
	Roe         float64   `gorm:"column:roe"`
	Roa         float64   `gorm:"column:roa"`
}

func (AtivoRentabilidade) TableName() string {
	return "ativo_rentabilidade"
}

func NewAtivoRentabilidade(ativoInformacao AtivoInformacao) *AtivoRentabilidade {
	return &AtivoRentabilidade{
		DataCalculo: ativoInformacao.DataInformacao,
		Codigo:      ativoInformacao.Codigo,
		Roe:         roundNumbers(ativoInformacao.LucroLiquido/ativoInformacao.PatrimonioLiquido, 4.0),
		Roa:         roundNumbers(ativoInformacao.LucroLiquido/ativoInformacao.AtivoTotal, 4.0),
	}
}
