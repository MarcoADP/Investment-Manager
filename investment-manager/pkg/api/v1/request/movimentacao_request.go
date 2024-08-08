package request

type MovimentacaoRequest struct {
	Data          string
	Codigo        string
	TipoAtivo     string
	Quantidade    float64
	ValorUnitario float64
}
