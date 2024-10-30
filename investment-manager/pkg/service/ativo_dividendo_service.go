package service

import (
	"log"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoDividendoService struct {
	repo            *repository.AtivoDividendoRepository
	dividendoRepo   *repository.DividendoRepository
	consolidacaoRep *repository.ConsolidacaoRepository
}

func NewAtivoDividendoService(
	repo *repository.AtivoDividendoRepository,
	dividendoRepo *repository.DividendoRepository,
	consolidacaoRep *repository.ConsolidacaoRepository,
) *AtivoDividendoService {
	return &AtivoDividendoService{
		repo:            repo,
		dividendoRepo:   dividendoRepo,
		consolidacaoRep: consolidacaoRep,
	}
}

func (s *AtivoDividendoService) CreateAtivoDividendo(informacao model.AtivoInformacao, precoAtual float64) (response.AtivoDividendoResponse, error) {
	dataInicial := informacao.DataInformacao.AddDate(-1, 0, 0)
	log.Println("Data Inicial: ", dataInicial)
	dividendos, err := s.dividendoRepo.GetDividendosByCodigoAndIntervalo(informacao.Codigo, dataInicial, informacao.DataInformacao)
	somaDividendos := 0.0
	if err == nil {
		for _, value := range dividendos {
			somaDividendos = somaDividendos + value.Valor
		}
	}
	log.Println("Dividendos: ", somaDividendos)

	consolidacao, err := s.consolidacaoRep.GetConsolidacaoByCodigo(informacao.Codigo)
	precoCompra := 0.0
	if err == nil {
		precoCompra = consolidacao.ValorMedioEntrada
	}
	log.Println("Compra: ", precoCompra)
	log.Println("Atual: ", precoAtual)

	dividendo := mapper.ToAtivoDividendo(informacao, somaDividendos, precoAtual, precoCompra)
	dividendoPersisted, err := s.repo.CreateDividendo(dividendo)
	if err != nil {
		return response.AtivoDividendoResponse{}, err
	}

	return mapper.ToAtivoDividendoResponse(dividendoPersisted), err
}

func (s *AtivoDividendoService) GetAtivoDividendoByInformacaoID(ativoInformacaoId uint) (response.AtivoDividendoResponse, error) {
	ativoDividendo, err := s.repo.GetDividendoByAtivoInformacao(ativoInformacaoId)
	if err != nil {
		return response.AtivoDividendoResponse{}, err
	}

	return mapper.ToAtivoDividendoResponse(ativoDividendo), err
}
