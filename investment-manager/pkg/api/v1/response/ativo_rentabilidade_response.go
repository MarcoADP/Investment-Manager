package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoRentabilidadeResponse struct {
	ID          uint
	DataCalculo time.Time
	Codigo      string
	Roe         float64
	Roa         float64
}

func NewAtivoRentabilidadeResponse(ativoRentabilidade model.AtivoRentabilidade) AtivoRentabilidadeResponse {
	return AtivoRentabilidadeResponse{
		ID:          ativoRentabilidade.ID,
		DataCalculo: ativoRentabilidade.DataCalculo,
		Codigo:      ativoRentabilidade.Codigo,
		Roe:         ativoRentabilidade.Roe,
		Roa:         ativoRentabilidade.Roa,
	}
}
