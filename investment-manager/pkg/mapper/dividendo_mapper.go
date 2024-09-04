package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"

	"time"
)

func ToDividendo(dividendoRequest request.DividendoRequest) (*model.Dividendo, error) {
	var dataCom time.Time
	if dividendoRequest.DataCom == "" {
		dataCom = time.Now()
	} else {
		var err error
		dataCom, err = time.Parse("02/01/2006", dividendoRequest.DataCom)
		if err != nil {
			return &model.Dividendo{}, err
		}
	}

	var dataPagamento time.Time
	if dividendoRequest.DataPagamento == "" {
		dataPagamento = time.Now()
	} else {
		var err error
		dataPagamento, err = time.Parse("02/01/2006", dividendoRequest.DataPagamento)
		if err != nil {
			return &model.Dividendo{}, err
		}
	}

	return model.NewDividendo(dataCom, dataPagamento, dividendoRequest.Tipo, dividendoRequest.Codigo, dividendoRequest.Valor), nil
}

func ToDividendoResponse(dividendo model.Dividendo) response.DividendoResponse {
	return response.NewDividendoResponse(
		dividendo.ID,
		dividendo.DataCom,
		dividendo.DataPagamento,
		dividendo.Tipo,
		dividendo.Codigo,
		dividendo.Valor,
	)
}

func ToDividendoResponseArray(cotacoes []model.Dividendo) []response.DividendoResponse {
	var cotacoesResponse []response.DividendoResponse
	for _, value := range cotacoes {
		cotacoesResponse = append(cotacoesResponse, ToDividendoResponse(value))
	}

	if cotacoesResponse == nil {
		cotacoesResponse = []response.DividendoResponse{}
	}
	return cotacoesResponse
}
