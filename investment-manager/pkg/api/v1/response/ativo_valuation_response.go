package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoValuationResponse struct {
	ID          uint      `gorm:"column:ativo_valuation;primaryKey;autoIncrement"`
	DataCalculo time.Time `gorm:"column:data_calculo;type:date"`
	Codigo      string    `gorm:"column:codigo"`
	LPA         float64   `gorm:"column:lpa"`
	PL          float64   `gorm:"column:p_l"`
	VPA         float64   `gorm:"column:vpa"`
	PVP         float64   `gorm:"column:p_vp"`
	EvEbit      float64   `gorm:"column:ev_ebit"`
	PEbit       float64   `gorm:"column:p_ebit"`
	EvEbitda    float64   `gorm:"column:ev_ebitda"`
	PEbitda     float64   `gorm:"column:p_ebit"`
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
