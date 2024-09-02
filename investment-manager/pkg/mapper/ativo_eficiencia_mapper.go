package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAtivoEficiencia(ativoInformacao model.AtivoInformacao) model.AtivoEficiencia {
	return *model.NewAtivoEficiencia(ativoInformacao)
}

func ToAtivoEficienciaResponse(ativoValuation model.AtivoEficiencia) response.AtivoEficienciaResponse {
	return response.NewAtivoEficienciaResponse(ativoValuation)
}
