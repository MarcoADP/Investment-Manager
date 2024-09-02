package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoEndividamentoResponse struct {
	ID                        uint
	DataCalculo               time.Time
	Codigo                    string
	Divida_Patrimonio_Liquido float64
	Divida_Ebit               float64
	Divida_Ebitda             float64
}

func NewAtivoEndividamentoResponse(ativoEndividamento model.AtivoEndividamento,
) AtivoEndividamentoResponse {
	return AtivoEndividamentoResponse{
		ID:                        ativoEndividamento.ID,
		DataCalculo:               ativoEndividamento.DataCalculo,
		Codigo:                    ativoEndividamento.Codigo,
		Divida_Patrimonio_Liquido: ativoEndividamento.Divida_Patrimonio_Liquido,
		Divida_Ebit:               ativoEndividamento.Divida_Ebit,
		Divida_Ebitda:             ativoEndividamento.Divida_Ebitda,
	}
}
