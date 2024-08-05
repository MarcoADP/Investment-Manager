package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToFundoImobiliario(fundoImobiliarioRequest request.FundoImobiliarioRequest) *model.FundoImobiliario {
	return model.NewFundoImobiliario(fundoImobiliarioRequest.Nome, fundoImobiliarioRequest.Codigo, fundoImobiliarioRequest.Tipo,
		fundoImobiliarioRequest.Segmento, fundoImobiliarioRequest.Cnpj)
}

func ToFundoImobiliarioResponse(fundoImobiliario model.FundoImobiliario) response.FundoImobiliarioResponse {
	return response.NewFundoImobiliarioResponse(fundoImobiliario.ID, fundoImobiliario.Nome, fundoImobiliario.Codigo, fundoImobiliario.Tipo,
		fundoImobiliario.Segmento, fundoImobiliario.Cnpj)
}

func UpdateFundoImobiliario(fundoImobiliario model.FundoImobiliario, fundoImobiliarioRequest request.FundoImobiliarioRequest) model.FundoImobiliario {
	fundoImobiliario.Codigo = fundoImobiliarioRequest.Codigo
	fundoImobiliario.Nome = fundoImobiliarioRequest.Nome
	fundoImobiliario.Cnpj = fundoImobiliarioRequest.Cnpj
	fundoImobiliario.Tipo = fundoImobiliarioRequest.Tipo
	fundoImobiliario.Segmento = fundoImobiliarioRequest.Segmento
	return fundoImobiliario
}
