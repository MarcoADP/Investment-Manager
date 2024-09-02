package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoEndividamentoResponse struct {
	ID                      uint
	DataCalculo             time.Time
	Codigo                  string
	DividaPatrimonioLiquido float64
	DividaEbit              float64
	DividaEbitda            float64
}

func NewAtivoEndividamentoResponse(ativoEndividamento model.AtivoEndividamento,
) AtivoEndividamentoResponse {
	return AtivoEndividamentoResponse{
		ID:                      ativoEndividamento.ID,
		DataCalculo:             ativoEndividamento.DataCalculo,
		Codigo:                  ativoEndividamento.Codigo,
		DividaPatrimonioLiquido: ativoEndividamento.DividaPatrimonioLiquido,
		DividaEbit:              ativoEndividamento.DividaEbit,
		DividaEbitda:            ativoEndividamento.DividaEbitda,
	}
}
