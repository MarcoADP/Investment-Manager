package model

import (
	"time"
)

type AtivoDividendo struct {
	ID                uint      `gorm:"column:ativo_dividendo_id;primaryKey;autoIncrement"`
	AtivoInformacaoId uint      `gorm:"column:ativo_informacao_id"`
	DataCalculo       time.Time `gorm:"column:data_calculo;type:date"`
	Codigo            string    `gorm:"column:codigo"`
	Dividendos        float64   `gorm:"column:dividendos"`
	Dy                float64   `gorm:"column:dy"`
	Yoc               float64   `gorm:"column:yoc"`
}

func (AtivoDividendo) TableName() string {
	return "ativo_dividendo"
}

func NewAtivoDividendo(ativoInformacao AtivoInformacao, dividendos float64, precoAtual float64, precoCompra float64) *AtivoDividendo {
	dy := 0.0
	yoc := 0.0
	if precoAtual > 0 {
		dy = dividendos / precoAtual
	}

	if precoCompra > 0 {
		yoc = dividendos / precoCompra
	}

	return &AtivoDividendo{
		AtivoInformacaoId: ativoInformacao.ID,
		DataCalculo:       ativoInformacao.DataInformacao,
		Codigo:            ativoInformacao.Codigo,
		Dividendos:        dividendos,
		Dy:                roundNumbers(dy, 4.0),
		Yoc:               roundNumbers(yoc, 4.0),
	}
}
