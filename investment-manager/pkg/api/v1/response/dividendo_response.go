package response

import (
	"time"
)

type DividendoResponse struct {
	ID            uint
	DataCom       time.Time
	DataPagamento time.Time
	Tipo          string
	Codigo        string
	Valor         float64
}

func NewDividendoResponse(id uint, dataCom time.Time, dataPagamento time.Time, tipo string, codigo string, valor float64) DividendoResponse {
	return DividendoResponse{
		ID:            id,
		DataCom:       dataCom,
		DataPagamento: dataPagamento,
		Codigo:        codigo,
		Tipo:          tipo,
		Valor:         valor,
	}
}
