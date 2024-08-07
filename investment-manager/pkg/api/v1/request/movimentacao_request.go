package request

type MovimentacaoRequest struct {
	Data          string
	Codigo        string
	TipoAtivo     string
	Quantidade    int
	ValorUnitario float64
}
