package request

type AtivoInformacaoRequest struct {
	DataInformacao    string
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
}
