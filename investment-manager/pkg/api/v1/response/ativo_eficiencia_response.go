package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoEficienciaResponse struct {
	ID            uint
	DataCalculo   time.Time
	Codigo        string
	MargemLiquida float64
	MargemBruta   float64
	MargemEbit    float64
	MargemEbitda  float64
}

func NewAtivoEficienciaResponse(ativoEficiencia model.AtivoEficiencia) AtivoEficienciaResponse {
	return AtivoEficienciaResponse{
		ID:            ativoEficiencia.ID,
		DataCalculo:   ativoEficiencia.DataCalculo,
		Codigo:        ativoEficiencia.Codigo,
		MargemLiquida: ativoEficiencia.MargemLiquida,
		MargemBruta:   ativoEficiencia.MargemBruta,
		MargemEbit:    ativoEficiencia.MargemEbit,
		MargemEbitda:  ativoEficiencia.MargemEbitda,
	}
}
