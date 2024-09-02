package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"

	"time"
)

func ToAtivoInformacao(informacaoRequest request.AtivoInformacaoRequest) (*model.AtivoInformacao, error) {
	var date time.Time
	if informacaoRequest.DataInformacao == "" {
		date = time.Now()
	} else {
		var err error
		date, err = time.Parse("02/01/2006", informacaoRequest.DataInformacao)
		if err != nil {
			return &model.AtivoInformacao{}, err
		}
	}

	return model.NewAtivoInformacao(date, informacaoRequest.Codigo, informacaoRequest.NumeroAcoes, informacaoRequest.ValorFirma,
		informacaoRequest.LucroLiquido, informacaoRequest.LucroBruto, informacaoRequest.ReceitaLiquida, informacaoRequest.PatrimonioLiquido,
		informacaoRequest.AtivoTotal, informacaoRequest.DividaLiquida, informacaoRequest.Ebit, informacaoRequest.Ebitda), nil
}

func ToAtivoInformacaoResponse(informacao model.AtivoInformacao) response.AtivoInformacaoResponse {
	return response.NewAtivoInformacaoResponse(
		informacao.ID,
		informacao.DataInformacao,
		informacao.Codigo,
		informacao.NumeroAcoes,
		informacao.ValorFirma,
		informacao.LucroLiquido,
		informacao.LucroBruto,
		informacao.ReceitaLiquida,
		informacao.PatrimonioLiquido,
		informacao.AtivoTotal,
		informacao.DividaLiquida,
		informacao.Ebit,
		informacao.Ebitda,
	)
}

func ToAtivoInformacaoResponseArray(cotacoes []model.AtivoInformacao) []response.AtivoInformacaoResponse {
	var cotacoesResponse []response.AtivoInformacaoResponse
	for _, value := range cotacoes {
		cotacoesResponse = append(cotacoesResponse, ToAtivoInformacaoResponse(value))
	}

	if cotacoesResponse == nil {
		cotacoesResponse = []response.AtivoInformacaoResponse{}
	}
	return cotacoesResponse
}
