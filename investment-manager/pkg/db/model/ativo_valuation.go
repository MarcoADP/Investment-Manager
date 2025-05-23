package model

import (
	"math"
	"time"
)

type AtivoValuation struct {
	ID                uint      `gorm:"column:ativo_valuation_id;primaryKey;autoIncrement"`
	AtivoInformacaoId uint      `gorm:"column:ativo_informacao_id"`
	DataCalculo       time.Time `gorm:"column:data_calculo;type:date"`
	Codigo            string    `gorm:"column:codigo"`
	LPA               float64   `gorm:"column:lpa"`
	PL                float64   `gorm:"column:p_l"`
	VPA               float64   `gorm:"column:vpa"`
	PVP               float64   `gorm:"column:p_vp"`
	EvEbit            float64   `gorm:"column:ev_ebit"`
	PEbit             float64   `gorm:"column:p_ebit"`
	EvEbitda          float64   `gorm:"column:ev_ebitda"`
	PEbitda           float64   `gorm:"column:p_ebitda"`
}

func (AtivoValuation) TableName() string {
	return "ativo_valuation"
}

func NewAtivoValuation(ativoInformacao AtivoInformacao, preco float64) *AtivoValuation {
	lpa := ativoInformacao.LucroLiquido / float64(ativoInformacao.NumeroAcoes)
	vpa := ativoInformacao.PatrimonioLiquido / float64(ativoInformacao.NumeroAcoes)
	ebitPerAcao := ativoInformacao.Ebit / float64(ativoInformacao.NumeroAcoes)
	evEbitda := 0.00
	pEbitda := 0.0
	if ativoInformacao.Ebitda != 0 {
		ebitdaPerAcao := ativoInformacao.Ebitda / float64(ativoInformacao.NumeroAcoes)
		evEbitda = ativoInformacao.ValorFirma / ativoInformacao.Ebitda
		pEbitda = preco / ebitdaPerAcao
	}

	return &AtivoValuation{
		AtivoInformacaoId: ativoInformacao.ID,
		DataCalculo:       ativoInformacao.DataInformacao,
		Codigo:            ativoInformacao.Codigo,
		LPA:               roundNumbers(lpa, 2.0),
		PL:                roundNumbers(preco/lpa, 2.0),
		VPA:               roundNumbers(vpa, 2.0),
		PVP:               roundNumbers(preco/vpa, 2.0),
		EvEbit:            roundNumbers(ativoInformacao.ValorFirma/ativoInformacao.Ebit, 2.0),
		PEbit:             roundNumbers(preco/ebitPerAcao, 2.0),
		EvEbitda:          roundNumbers(evEbitda, 2.0),
		PEbitda:           roundNumbers(pEbitda, 2.0),
	}
}

func roundNumbers(number float64, casas float64) float64 {
	n := math.Pow(10, casas)
	return math.Round(number*n) / n
}
