package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
)

type AcaoBrComparadorService struct {
	ativoInformacaoService *AtivoInformacaoService
}

func NewAcaoBrComparadorService(ativoInformacaoService *AtivoInformacaoService) *AcaoBrComparadorService {
	return &AcaoBrComparadorService{
		ativoInformacaoService: ativoInformacaoService,
	}
}

func (s *AcaoBrComparadorService) CompareAcoes(codigos []string) (response.AcaoBrComparadorResponse, error) {
	var informacoes []response.AtivoInformacaoResponse
	var err error
	for _, value := range codigos {
		informacao, err := s.ativoInformacaoService.GetInformacaoMoreRecentByCodigo(value)
		if err == nil {
			informacoes = append(informacoes, informacao)
		}
	}

	response := s.agruparMetricas(informacoes)

	return response, err

}

func (s *AcaoBrComparadorService) agruparMetricas(informacoes []response.AtivoInformacaoResponse) response.AcaoBrComparadorResponse {

	var response response.AcaoBrComparadorResponse

	for _, informacao := range informacoes {
		response.Codigo = append(response.Codigo, informacao.Codigo)
		response.Informacao = s.agruparInformacaoMetricas(response.Informacao, informacao)
		response.Valuation = s.agruparValuationMetricas(response.Valuation, informacao.Valuation)
		response.Endividamento = s.agruparEndividamentoMetricas(response.Endividamento, informacao.Endividamento)
		response.Eficiencia = s.agruparEficienciaMetricas(response.Eficiencia, informacao.Eficiencia)
		response.Rentabilidade = s.agruparRentabilidadeMetricas(response.Rentabilidade, informacao.Rentabilidade)
		response.Dividendo = s.agruparDividendoMetricas(response.Dividendo, informacao.Dividendo)
	}

	return response

}

func (s *AcaoBrComparadorService) agruparInformacaoMetricas(responseInformacao response.ComparadorInformacaoResponse, informacao response.AtivoInformacaoResponse) response.ComparadorInformacaoResponse {
	responseInformacao.ValorFirma = append(responseInformacao.ValorFirma, map[string]float64{informacao.Codigo: informacao.ValorFirma})
	responseInformacao.LucroLiquido = append(responseInformacao.LucroLiquido, map[string]float64{informacao.Codigo: informacao.LucroLiquido})
	responseInformacao.LucroBruto = append(responseInformacao.LucroBruto, map[string]float64{informacao.Codigo: informacao.LucroBruto})
	responseInformacao.ReceitaLiquida = append(responseInformacao.ReceitaLiquida, map[string]float64{informacao.Codigo: informacao.ReceitaLiquida})
	responseInformacao.PatrimonioLiquido = append(responseInformacao.PatrimonioLiquido, map[string]float64{informacao.Codigo: informacao.PatrimonioLiquido})
	responseInformacao.AtivoTotal = append(responseInformacao.AtivoTotal, map[string]float64{informacao.Codigo: informacao.AtivoTotal})
	responseInformacao.DividaLiquida = append(responseInformacao.DividaLiquida, map[string]float64{informacao.Codigo: informacao.DividaLiquida})
	responseInformacao.Ebit = append(responseInformacao.Ebit, map[string]float64{informacao.Codigo: informacao.Ebit})
	responseInformacao.Ebitda = append(responseInformacao.Ebitda, map[string]float64{informacao.Codigo: informacao.Ebitda})
	return responseInformacao
}

func (s *AcaoBrComparadorService) agruparValuationMetricas(responseValuation response.ComparadorValuationResponse, valuation response.AtivoValuationResponse) response.ComparadorValuationResponse {
	responseValuation.LPA = append(responseValuation.LPA, map[string]float64{valuation.Codigo: valuation.LPA})
	responseValuation.PL = append(responseValuation.PL, map[string]float64{valuation.Codigo: valuation.PL})
	responseValuation.VPA = append(responseValuation.VPA, map[string]float64{valuation.Codigo: valuation.VPA})
	responseValuation.PVP = append(responseValuation.PVP, map[string]float64{valuation.Codigo: valuation.PVP})
	responseValuation.EvEbit = append(responseValuation.EvEbit, map[string]float64{valuation.Codigo: valuation.EvEbit})
	responseValuation.PEbit = append(responseValuation.PEbit, map[string]float64{valuation.Codigo: valuation.PEbit})
	responseValuation.EvEbitda = append(responseValuation.EvEbitda, map[string]float64{valuation.Codigo: valuation.EvEbitda})
	responseValuation.PEbitda = append(responseValuation.PEbitda, map[string]float64{valuation.Codigo: valuation.PEbitda})
	return responseValuation
}

func (s *AcaoBrComparadorService) agruparEndividamentoMetricas(responseEndividamento response.ComparadorEndividamentoResponse, endividamento response.AtivoEndividamentoResponse) response.ComparadorEndividamentoResponse {
	responseEndividamento.DividaPatrimonioLiquido = append(responseEndividamento.DividaPatrimonioLiquido, map[string]float64{endividamento.Codigo: endividamento.DividaPatrimonioLiquido})
	responseEndividamento.DividaEbit = append(responseEndividamento.DividaEbit, map[string]float64{endividamento.Codigo: endividamento.DividaEbit})
	responseEndividamento.DividaEbitda = append(responseEndividamento.DividaEbitda, map[string]float64{endividamento.Codigo: endividamento.DividaEbitda})
	return responseEndividamento
}

func (s *AcaoBrComparadorService) agruparEficienciaMetricas(responseEficiencia response.ComparadorEficienciaResponse, endividamento response.AtivoEficienciaResponse) response.ComparadorEficienciaResponse {
	responseEficiencia.MargemLiquida = append(responseEficiencia.MargemLiquida, map[string]float64{endividamento.Codigo: endividamento.MargemLiquida})
	responseEficiencia.MargemBruta = append(responseEficiencia.MargemBruta, map[string]float64{endividamento.Codigo: endividamento.MargemBruta})
	responseEficiencia.MargemEbit = append(responseEficiencia.MargemEbit, map[string]float64{endividamento.Codigo: endividamento.MargemEbit})
	responseEficiencia.MargemEbitda = append(responseEficiencia.MargemEbitda, map[string]float64{endividamento.Codigo: endividamento.MargemEbitda})
	return responseEficiencia
}

func (s *AcaoBrComparadorService) agruparRentabilidadeMetricas(responseRentabilidade response.ComparadorRentabilidadeResponse, rentabilidade response.AtivoRentabilidadeResponse) response.ComparadorRentabilidadeResponse {
	responseRentabilidade.ROE = append(responseRentabilidade.ROE, map[string]float64{rentabilidade.Codigo: rentabilidade.Roe})
	responseRentabilidade.ROA = append(responseRentabilidade.ROA, map[string]float64{rentabilidade.Codigo: rentabilidade.Roa})
	return responseRentabilidade
}

func (s *AcaoBrComparadorService) agruparDividendoMetricas(responseDividendo response.ComparadorDividendoResponse, dividendo response.AtivoDividendoResponse) response.ComparadorDividendoResponse {
	responseDividendo.Dividendos = append(responseDividendo.Dividendos, map[string]float64{dividendo.Codigo: dividendo.Dividendos})
	responseDividendo.DY = append(responseDividendo.DY, map[string]float64{dividendo.Codigo: dividendo.Dy})
	responseDividendo.YOC = append(responseDividendo.YOC, map[string]float64{dividendo.Codigo: dividendo.Yoc})
	return responseDividendo
}
