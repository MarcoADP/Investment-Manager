package response

import (
	"time"
)

type CotacaoHistoricoResponse struct {
	ID     uint
	Data   time.Time
	Codigo string
	Valor  float64
}

func NewCotacaoHistoricoResponse(id uint, data time.Time, codigo string, valor float64) CotacaoHistoricoResponse {
	return CotacaoHistoricoResponse{
		ID:     id,
		Data:   data,
		Codigo: codigo,
		Valor:  valor,
	}
}
