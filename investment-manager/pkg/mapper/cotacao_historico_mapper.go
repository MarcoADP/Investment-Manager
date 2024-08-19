package mapper

import (
	"github.com/MarcoADP/Investment-Manager/internal/brapi"
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

func ToCotacaoHistoricoFromBrapi(cotacaoResponse brapi.BrapiCotacaoResponse) (*model.CotacaoHistorico, error) {

	if len(cotacaoResponse.Results) == 0 {
		return &model.CotacaoHistorico{}, nil
	}

	cotacao := cotacaoResponse.Results[0]
	var err error
	date, err := convertToDate(cotacao.RegularMarketTime)
	if err != nil {
		return &model.CotacaoHistorico{}, err
	}

	return model.NewCotacaoHistorico(date, cotacao.Symbol, cotacao.RegularMarketPrice), nil
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

func convertToDate(input string) (time.Time, error) {
	// Parse o tempo no formato RFC3339
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Time{}, err
	}

	// Formate o tempo para "02/01/2006" como time.Time
	formattedDate, err := time.Parse("02/01/2006", t.Format("02/01/2006"))
	if err != nil {
		return time.Time{}, err
	}

	return formattedDate, nil
}
