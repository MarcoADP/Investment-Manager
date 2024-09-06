package model

import (
	"time"
)

type AtivoEficiencia struct {
	ID                uint      `gorm:"column:ativo_eficiencia_id;primaryKey;autoIncrement"`
	AtivoInformacaoId uint      `gorm:"column:ativo_informacao_id"`
	DataCalculo       time.Time `gorm:"column:data_calculo;type:date"`
	Codigo            string    `gorm:"column:codigo"`
	MargemLiquida     float64   `gorm:"column:margem_liquida"`
	MargemBruta       float64   `gorm:"column:margem_bruta"`
	MargemEbit        float64   `gorm:"column:margem_ebit"`
	MargemEbitda      float64   `gorm:"column:margem_ebitda"`
}

func (AtivoEficiencia) TableName() string {
	return "ativo_eficiencia"
}

func NewAtivoEficiencia(ativoInformacao AtivoInformacao) *AtivoEficiencia {
	return &AtivoEficiencia{
		AtivoInformacaoId: ativoInformacao.ID,
		DataCalculo:       ativoInformacao.DataInformacao,
		Codigo:            ativoInformacao.Codigo,
		MargemLiquida:     roundNumbers(ativoInformacao.LucroLiquido/ativoInformacao.ReceitaLiquida, 4.0),
		MargemBruta:       roundNumbers(ativoInformacao.LucroBruto/ativoInformacao.ReceitaLiquida, 4.0),
		MargemEbit:        roundNumbers(ativoInformacao.Ebit/ativoInformacao.ReceitaLiquida, 4.0),
		MargemEbitda:      roundNumbers(ativoInformacao.Ebitda/ativoInformacao.ReceitaLiquida, 4.0),
	}
}
