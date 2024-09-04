package response

import (
	"math"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

type CarteiraAtivoResponse struct {
	ID                uint
	Codigo            string
	ProporcaoDesejada float64
	Quantidade        float64
	PrecoCompra       float64
	TotalCompra       float64
	PrecoAtual        float64
	TotalAtual        float64
	Saldo             float64
	Variacao          float64
	Movimento         string
}

func NewCarteiraAtivoResponse(
	id uint,
	codigo string,
	proporcaoDesejada float64,
	movimento string,
	consolidacao model.Consolidacao,
	cotacao model.CotacaoHistorico,
) CarteiraAtivoResponse {
	quantidade := consolidacao.QuantidadeEntrada - consolidacao.QuantidadeSaida
	totalAtual := quantidade * cotacao.Valor
	saldo := totalAtual - consolidacao.ValorTotalEntrada
	variacao := 0.0
	if consolidacao.ValorMedioEntrada > 0 {
		variacao = (totalAtual/consolidacao.ValorTotalEntrada - 1)
	}

	return CarteiraAtivoResponse{
		ID:                id,
		Codigo:            codigo,
		ProporcaoDesejada: proporcaoDesejada,
		Quantidade:        quantidade,
		PrecoCompra:       consolidacao.ValorMedioEntrada,
		TotalCompra:       consolidacao.ValorTotalEntrada,
		PrecoAtual:        cotacao.Valor,
		TotalAtual:        totalAtual,
		Saldo:             roundNumbers(saldo, 2.0),
		Variacao:          roundNumbers(variacao, 4.0),
		Movimento:         movimento,
	}
}

func roundNumbers(number float64, casas float64) float64 {
	n := math.Pow(10, casas)
	return math.Round(number*n) / n
}
