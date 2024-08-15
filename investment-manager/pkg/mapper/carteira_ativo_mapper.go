package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToCarteiraAtivo(carteiraId uint, ativoRequest request.CarteiraAtivoRequest) *model.CarteiraAtivo {
	return model.NewCarteiraAtivo(carteiraId, ativoRequest.Codigo, ativoRequest.ProporcaoDesejada)
}

func ToCarteiraAtivoResponse(ativo model.CarteiraAtivo) response.CarteiraAtivoResponse {
	return response.NewCarteiraAtivoResponse(ativo.ID, ativo.Codigo, ativo.ProporcaoDesejada)
}

func UpdateCarteiraAtivo(ativo model.CarteiraAtivo, ativoRequest request.CarteiraAtivoRequest) model.CarteiraAtivo {
	ativo.Codigo = ativoRequest.Codigo
	ativo.ProporcaoDesejada = ativoRequest.ProporcaoDesejada
	return ativo
}
