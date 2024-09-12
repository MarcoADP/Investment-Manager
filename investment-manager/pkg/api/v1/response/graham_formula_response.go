package response

import (
	"time"
)

type GrahamFormulaResponse struct {
	ID              uint
	Data            time.Time
	Codigo          string
	PrecoAtual      float64
	Lpa             float64
	Vpa             float64
	PlEsperado      float64
	PvpEsperado     float64
	PrecoJusto      float64
	MargemSeguranca float64
}

func NewGrahamFormulaResponse(id uint, data time.Time, codigo string, precoAtual float64, lpa float64, vpa float64, pl float64, pvp float64,
	precoJusto float64, margem float64,
) GrahamFormulaResponse {
	return GrahamFormulaResponse{
		ID:              id,
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
