package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"

	"time"
)

func ToMovimentacao(movimentacaoRequest request.MovimentacaoRequest, operacao string) (*model.Movimentacao, error) {
	date, err := time.Parse("02/01/2006", movimentacaoRequest.Data)
	if err != nil {
		return &model.Movimentacao{}, err
	}

	return model.NewMovimentacao(date,
		operacao,
		movimentacaoRequest.Codigo,
		movimentacaoRequest.TipoAtivo,
		movimentacaoRequest.Quantidade,
		movimentacaoRequest.ValorUnitario,
	), nil
}

func ToMovimentacaoResponse(movimentacao model.Movimentacao) response.MovimentacaoResponse {
	return response.NewMovimentacaoResponse(
		movimentacao.ID,
		movimentacao.Data,
		movimentacao.Operacao,
		movimentacao.Codigo,
		movimentacao.TipoAtivo,
		movimentacao.Quantidade,
		movimentacao.ValorUnitario,
		movimentacao.ValorTotal,
	)
}
