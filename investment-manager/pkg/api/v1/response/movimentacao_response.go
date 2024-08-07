package response

import (
	"time"
)

type MovimentacaoResponse struct {
	ID            uint
	Data          time.Time
	Operacao      string
	Codigo        string
	TipoAtivo     string
	Quantidade    int
	ValorUnitario float64
	ValorTotal    float64
}

func NewMovimentacaoResponse(id uint, data time.Time, operacao string, codigo string, tipoAtivo string, quantidade int,
	valorUnitario float64, valorTotal float64) MovimentacaoResponse {
	return MovimentacaoResponse{
		ID:            id,
		Data:          data,
		Operacao:      operacao,
		Codigo:        codigo,
		TipoAtivo:     tipoAtivo,
		Quantidade:    quantidade,
		ValorUnitario: valorUnitario,
		ValorTotal:    valorTotal,
	}
}
