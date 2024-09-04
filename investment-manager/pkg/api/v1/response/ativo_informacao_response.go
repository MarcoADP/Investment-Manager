package response

import (
	"time"
)

type AtivoInformacaoResponse struct {
	ID                uint
	DataInformacao    time.Time
	Codigo            string
	NumeroAcoes       uint64
	ValorFirma        float64
	LucroLiquido      float64
	LucroBruto        float64
	ReceitaLiquida    float64
	PatrimonioLiquido float64
	AtivoTotal        float64
	DividaLiquida     float64
	Ebit              float64
	Ebitda            float64
	Valuation         AtivoValuationResponse
	Endividamento     AtivoEndividamentoResponse
	Eficiencia        AtivoEficienciaResponse
	Rentabilidade     AtivoRentabilidadeResponse
}

func NewAtivoInformacaoResponse(id uint, data time.Time, codigo string, numeroAcoes uint64, valorFirma float64,
	lucroLiquido float64, lucroBruto float64, receitaLiquida float64, patrimonioLiquido float64,
	ativoTotal float64, dividaLiquida float64, ebit float64, ebitda float64,
) AtivoInformacaoResponse {
	return AtivoInformacaoResponse{
		ID:                id,
		DataInformacao:    data,
		Codigo:            codigo,
		NumeroAcoes:       numeroAcoes,
		ValorFirma:        valorFirma,
		LucroLiquido:      lucroLiquido,
		LucroBruto:        lucroBruto,
		ReceitaLiquida:    receitaLiquida,
		PatrimonioLiquido: patrimonioLiquido,
		AtivoTotal:        ativoTotal,
		DividaLiquida:     dividaLiquida,
		Ebit:              ebit,
		Ebitda:            ebitda,
	}
}
