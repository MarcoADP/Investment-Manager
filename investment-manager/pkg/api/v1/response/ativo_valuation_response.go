package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoValuationResponse struct {
	ID          uint
	DataCalculo time.Time
	Codigo      string
	LPA         float64
	PL          float64
	VPA         float64
	PVP         float64
	EvEbit      float64
	PEbit       float64
	EvEbitda    float64
	PEbitda     float64
}

func NewAtivoValuationResponse(ativoValuation model.AtivoValuation,
) AtivoValuationResponse {
	return AtivoValuationResponse{
		ID:          ativoValuation.ID,
		DataCalculo: ativoValuation.DataCalculo,
		Codigo:      ativoValuation.Codigo,
		LPA:         ativoValuation.LPA,
		PL:          ativoValuation.PL,
		VPA:         ativoValuation.VPA,
		PVP:         ativoValuation.PVP,
		EvEbit:      ativoValuation.EvEbit,
		PEbit:       ativoValuation.PEbit,
		EvEbitda:    ativoValuation.EvEbitda,
		PEbitda:     ativoValuation.PEbitda,
	}
}
