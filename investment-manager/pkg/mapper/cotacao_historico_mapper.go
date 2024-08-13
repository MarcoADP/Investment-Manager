package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"

	"time"
)

func ToCotacaoHistorico(cotacaoHistoricoRequest request.CotacaoHistoricoRequest) (*model.CotacaoHistorico, error) {
	var date time.Time
	if cotacaoHistoricoRequest.Data == "" {
		date = time.Now()
	} else {
		var err error
		date, err = time.Parse("02/01/2006", cotacaoHistoricoRequest.Data)
		if err != nil {
			return &model.CotacaoHistorico{}, err
		}
	}

	return model.NewCotacaoHistorico(date, cotacaoHistoricoRequest.Codigo, cotacaoHistoricoRequest.Valor), nil
}

func ToCotacaoHistoricoResponse(cotacaoHistorico model.CotacaoHistorico) response.CotacaoHistoricoResponse {
	return response.NewCotacaoHistoricoResponse(
		cotacaoHistorico.ID,
		cotacaoHistorico.Data,
		cotacaoHistorico.Codigo,
		cotacaoHistorico.Valor,
	)
}

func ToCotacaoHistoricoResponseArray(cotacoes []model.CotacaoHistorico) []response.CotacaoHistoricoResponse {
	var cotacoesResponse []response.CotacaoHistoricoResponse
	for _, value := range cotacoes {
		cotacoesResponse = append(cotacoesResponse, ToCotacaoHistoricoResponse(value))
	}

	if cotacoesResponse == nil {
		cotacoesResponse = []response.CotacaoHistoricoResponse{}
	}
	return cotacoesResponse
}
