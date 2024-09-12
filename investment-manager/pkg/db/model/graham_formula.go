package model

import (
	"math"
	"time"
)

type GrahamFormula struct {
	ID              uint      `gorm:"column:graham_formula_id;primaryKey;autoIncrement"`
	Data            time.Time `gorm:"column:data_calculo;type:date"`
	Codigo          string    `gorm:"column:codigo"`
	PrecoAtual      float64   `gorm:"column:preco_atual"`
	Lpa             float64   `gorm:"column:lpa"`
	Vpa             float64   `gorm:"column:vpa"`
	PlEsperado      float64   `gorm:"column:pl_esperado"`
	PvpEsperado     float64   `gorm:"column:pvp_esperado"`
	PrecoJusto      float64   `gorm:"column:preco_justo"`
	MargemSeguranca float64   `gorm:"column:margem_seguranca"`
}

func (GrahamFormula) TableName() string {
	return "graham_formula"
}

func NewGrahamFormula(data time.Time, codigo string, precoAtual float64, lpa float64, vpa float64, pl float64, pvp float64) *GrahamFormula {
	precoJusto := math.Sqrt(lpa * vpa * pl * pvp)
	margem := ((precoJusto / precoAtual) - 1) * 100
	return &GrahamFormula{
		Data:            data,
		Codigo:          codigo,
		PrecoAtual:      precoAtual,
		Lpa:             lpa,
		Vpa:             vpa,
		PlEsperado:      pl,
		PvpEsperado:     pvp,
		PrecoJusto:      precoJusto,
		MargemSeguranca: margem,
	}
}
