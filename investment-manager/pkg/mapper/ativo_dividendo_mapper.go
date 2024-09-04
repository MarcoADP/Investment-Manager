package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAtivoDividendo(ativoInformacao model.AtivoInformacao, dividendos float64, precoAtual float64, precoCompra float64) model.AtivoDividendo {
	return *model.NewAtivoDividendo(ativoInformacao, dividendos, precoAtual, precoCompra)
}

func ToAtivoDividendoResponse(ativoDividendo model.AtivoDividendo) response.AtivoDividendoResponse {
	return response.NewAtivoDividendoResponse(ativoDividendo)
}
