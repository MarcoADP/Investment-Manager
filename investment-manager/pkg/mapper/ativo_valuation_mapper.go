package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAtivoValuation(ativoInformacao model.AtivoInformacao, preco float64) model.AtivoValuation {
	return *model.NewAtivoValuation(ativoInformacao, preco)
}

func ToAtivoValuationResponse(ativoValuation model.AtivoValuation) response.AtivoValuationResponse {
	return response.NewAtivoValuationResponse(ativoValuation)
}
