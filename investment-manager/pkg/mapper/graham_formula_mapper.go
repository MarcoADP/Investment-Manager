package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"

	"time"
)

func ToGrahamFormula(request request.GrahamFormulaRequest) (*model.GrahamFormula, error) {
	var date time.Time
	if request.Data == "" {
		date = time.Now()
	} else {
		var err error
		date, err = time.Parse("02/01/2006", request.Data)
		if err != nil {
			return &model.GrahamFormula{}, err
		}
	}

	return model.NewGrahamFormula(date, request.Codigo, request.PrecoAtual, request.Lpa, request.Vpa, request.PlEsperado, request.PvpEsperado), nil
}

func ToGrahamFormulaResponse(data model.GrahamFormula) response.GrahamFormulaResponse {
	return response.NewGrahamFormulaResponse(
		data.ID,
		data.Data,
		data.Codigo,
		data.PrecoAtual,
		data.Lpa,
		data.Vpa,
		data.PlEsperado,
		data.PvpEsperado,
		data.PrecoJusto,
		data.MargemSeguranca,
	)
}

func ToGrahamFormulaResponseArray(data []model.GrahamFormula) []response.GrahamFormulaResponse {
	var responses []response.GrahamFormulaResponse
	for _, value := range data {
		responses = append(responses, ToGrahamFormulaResponse(value))
	}

	if responses == nil {
		responses = []response.GrahamFormulaResponse{}
	}
	return responses
}
