package response

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type AtivoDividendoResponse struct {
	ID          uint
	DataCalculo time.Time
	Codigo      string
	Dividendos  float64
	Dy          float64
	Yoc         float64
}

func NewAtivoDividendoResponse(ativoDividendo model.AtivoDividendo) AtivoDividendoResponse {
	return AtivoDividendoResponse{
		ID:          ativoDividendo.ID,
		DataCalculo: ativoDividendo.DataCalculo,
		Codigo:      ativoDividendo.Codigo,
		Dividendos:  ativoDividendo.Dividendos,
		Dy:          ativoDividendo.Dy,
		Yoc:         ativoDividendo.Yoc,
	}
}
