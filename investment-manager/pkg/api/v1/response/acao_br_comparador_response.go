package response

type AcaoBrComparadorResponse struct {
	Codigo        []string
	Informacao    ComparadorInformacaoResponse
	Valuation     ComparadorValuationResponse
	Endividamento ComparadorEndividamentoResponse
	Eficiencia    ComparadorEficienciaResponse
	Rentabilidade ComparadorRentabilidadeResponse
	Dividendo     ComparadorDividendoResponse
}

type ComparadorInformacaoResponse struct {
	ValorFirma        []map[string]float64
	LucroLiquido      []map[string]float64
	LucroBruto        []map[string]float64
	ReceitaLiquida    []map[string]float64
	PatrimonioLiquido []map[string]float64
	AtivoTotal        []map[string]float64
	DividaLiquida     []map[string]float64
	Ebit              []map[string]float64
	Ebitda            []map[string]float64
}

type ComparadorValuationResponse struct {
	LPA      []map[string]float64
	PL       []map[string]float64
	VPA      []map[string]float64
	PVP      []map[string]float64
	EvEbit   []map[string]float64
	PEbit    []map[string]float64
	EvEbitda []map[string]float64
	PEbitda  []map[string]float64
}

type ComparadorEndividamentoResponse struct {
	DividaPatrimonioLiquido []map[string]float64
	DividaEbit              []map[string]float64
	DividaEbitda            []map[string]float64
}

type ComparadorEficienciaResponse struct {
	MargemLiquida []map[string]float64
	MargemBruta   []map[string]float64
	MargemEbit    []map[string]float64
	MargemEbitda  []map[string]float64
}

type ComparadorRentabilidadeResponse struct {
	ROE []map[string]float64
	ROA []map[string]float64
}

type ComparadorDividendoResponse struct {
	Dividendos []map[string]float64
	DY         []map[string]float64
	YOC        []map[string]float64
}
