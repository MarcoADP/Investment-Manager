package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToConsolidacaoResponse(consolidacao model.Consolidacao) response.ConsolidacaoResponse {
	return response.NewConsolidacaoResponse(
		consolidacao.ID,
		consolidacao.Codigo,
		consolidacao.TipoAtivo,
		consolidacao.QuantidadeEntrada,
		consolidacao.ValorMedioEntrada,
		consolidacao.ValorTotalEntrada,
		consolidacao.QuantidadeSaida,
		consolidacao.ValorMedioSaida,
		consolidacao.ValorTotalSaida,
		consolidacao.LucroMedio,
		consolidacao.LucroProporcao,
	)
}
