package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToCarteira(carteiraRequest request.CarteiraRequest) *model.Carteira {
	return model.NewCarteira(carteiraRequest.Nome, carteiraRequest.Descricao, carteiraRequest.ProporcaoDesejada)
}

func ToCarteiraResponse(carteira model.Carteira) response.CarteiraResponse {
	return response.NewCarteiraResponse(carteira.ID, carteira.Nome, carteira.Descricao, carteira.ProporcaoDesejada, []response.CarteiraAtivoResponse{})
}

func UpdateCarteira(carteira model.Carteira, carteiraRequest request.CarteiraRequest) model.Carteira {
	carteira.Nome = carteiraRequest.Nome
	carteira.Descricao = carteiraRequest.Descricao
	carteira.ProporcaoDesejada = carteiraRequest.ProporcaoDesejada
	return carteira
}
