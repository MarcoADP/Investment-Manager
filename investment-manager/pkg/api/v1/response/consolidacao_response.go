package response

type ConsolidacaoResponse struct {
	ID                uint
	Codigo            string
	TipoAtivo         string
	QuantidadeEntrada float64
	ValorMedioEntrada float64
	ValorTotalEntrada float64
	QuantidadeSaida   float64
	ValorMedioSaida   float64
	ValorTotalSaida   float64
	LucroMedio        float64
	LucroProporcao    float64
}

func NewConsolidacaoResponse(id uint, codigo string, tipoAtivo string,
	quantidadeEntrada float64, valorMedioEntrada float64, valorTotalEntrada float64,
	quantidadeSaida float64, valorMedioSaida float64, valorTotalSaida float64,
	lucroMedio float64, lucroProporcao float64,
) ConsolidacaoResponse {
	return ConsolidacaoResponse{
		ID:                id,
		Codigo:            codigo,
		TipoAtivo:         tipoAtivo,
		QuantidadeEntrada: quantidadeEntrada,
		ValorMedioEntrada: valorMedioEntrada,
		ValorTotalEntrada: valorTotalEntrada,
		QuantidadeSaida:   quantidadeSaida,
		ValorMedioSaida:   valorMedioSaida,
		ValorTotalSaida:   valorTotalSaida,
		LucroMedio:        lucroMedio,
		LucroProporcao:    lucroProporcao,
	}
}
