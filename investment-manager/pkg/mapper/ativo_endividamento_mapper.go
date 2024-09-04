package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAtivoEndividamento(ativoInformacao model.AtivoInformacao) model.AtivoEndividamento {
	return *model.NewAtivoEndividamento(ativoInformacao)
}

func ToAtivoEndividamentoResponse(ativoEndividamento model.AtivoEndividamento) response.AtivoEndividamentoResponse {
	return response.NewAtivoEndividamentoResponse(ativoEndividamento)
}
