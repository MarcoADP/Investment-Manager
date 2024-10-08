package service

import (
	"log"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoInformacaoService struct {
	repo              *repository.AtivoInformacaoRepository
	cotacaoRep        *repository.CotacaoHistoricoRepository
	valuationRep      *repository.AtivoValuationRepository
	endividamentoRep  *repository.AtivoEndividamentoRepository
	eficienciaRep     *repository.AtivoEficienciaRepository
	rentabilidadeRep  *repository.AtivoRentabilidadeRepository
	ativoDividendoRep *repository.AtivoDividendoRepository
	dividendosRep     *repository.DividendoRepository
	consolidacaoRep   *repository.ConsolidacaoRepository
}

func NewAtivoInformacaoService(repo *repository.AtivoInformacaoRepository,
	cotacaoRep *repository.CotacaoHistoricoRepository,
	valuationRep *repository.AtivoValuationRepository,
	endividamentoRep *repository.AtivoEndividamentoRepository,
	eficienciaRep *repository.AtivoEficienciaRepository,
	rentabilidadeRep *repository.AtivoRentabilidadeRepository,
	ativoDividendoRep *repository.AtivoDividendoRepository,
	dividendosRep *repository.DividendoRepository,
	consolidacaoRep *repository.ConsolidacaoRepository,
) *AtivoInformacaoService {
	return &AtivoInformacaoService{
		repo:              repo,
		cotacaoRep:        cotacaoRep,
		valuationRep:      valuationRep,
		endividamentoRep:  endividamentoRep,
		eficienciaRep:     eficienciaRep,
		rentabilidadeRep:  rentabilidadeRep,
		ativoDividendoRep: ativoDividendoRep,
		dividendosRep:     dividendosRep,
		consolidacaoRep:   consolidacaoRep,
	}
}

func (s *AtivoInformacaoService) GetInformacoesByCodigo(codigo string) ([]response.AtivoInformacaoResponse, error) {
	informacoes, err := s.repo.GetInformacoesByCodigo(codigo)

	if err != nil {
		return []response.AtivoInformacaoResponse{}, err
	}

	return mapper.ToAtivoInformacaoResponseArray(informacoes), err
}

func (s *AtivoInformacaoService) GetInformacaoMoreRecentByCodigo(codigo string) (response.AtivoInformacaoResponse, error) {

	informacao, err := s.repo.GetInformacaoMoreRecentByCodigo(codigo)

	if err != nil {
		return response.AtivoInformacaoResponse{}, err
	}

	return mapper.ToAtivoInformacaoResponse(informacao), err
}

func (s *AtivoInformacaoService) CreateInformacao(informacaoRequest request.AtivoInformacaoRequest) (response.AtivoInformacaoResponse, error) {

	informacao, err := mapper.ToAtivoInformacao(informacaoRequest)
	if err != nil {
		return response.AtivoInformacaoResponse{}, err
	}

	informacaoPersisted, err := s.repo.CreateInformacao(*informacao)
	if err != nil {
		return response.AtivoInformacaoResponse{}, err
	}
	informacaoResponse := mapper.ToAtivoInformacaoResponse(informacaoPersisted)

	cotacao, err := s.cotacaoRep.GetCotacaoMoreRecentByCodigo(informacaoPersisted.Codigo)
	if err != nil {
		return informacaoResponse, err
	}

	valuation, err := s.createValuation(informacaoPersisted, cotacao.Valor)
	if err == nil {
		informacaoResponse.Valuation = valuation
	}

	endividamento, err := s.createEndividamento(informacaoPersisted)
	if err == nil {
		informacaoResponse.Endividamento = endividamento
	}

	eficiencia, err := s.createEficiencia(informacaoPersisted)
	if err == nil {
		informacaoResponse.Eficiencia = eficiencia
	}

	rentabilidade, err := s.createRentabilidade(informacaoPersisted)
	if err == nil {
		informacaoResponse.Rentabilidade = rentabilidade
	}

	dividendo, err := s.createDividendo(informacaoPersisted, cotacao.Valor)
	if err == nil {
		informacaoResponse.Dividendo = dividendo
	}

	return informacaoResponse, err
}

func (s *AtivoInformacaoService) DeleteInformacao(id uint) error {
	return s.repo.DeleteInformacao(id)
}

func (s *AtivoInformacaoService) createValuation(informacao model.AtivoInformacao, valor float64) (response.AtivoValuationResponse, error) {
	valuation := mapper.ToAtivoValuation(informacao, valor)
	valuationPersisted, err := s.valuationRep.CreateValuation(valuation)
	if err != nil {
		return response.AtivoValuationResponse{}, err
	}

	return mapper.ToAtivoValuationResponse(valuationPersisted), err
}

func (s *AtivoInformacaoService) createEndividamento(informacao model.AtivoInformacao) (response.AtivoEndividamentoResponse, error) {
	endividamento := mapper.ToAtivoEndividamento(informacao)
	endividamentoPersisted, err := s.endividamentoRep.CreateEndividamento(endividamento)
	if err != nil {
		return response.AtivoEndividamentoResponse{}, err
	}

	return mapper.ToAtivoEndividamentoResponse(endividamentoPersisted), err
}

func (s *AtivoInformacaoService) createEficiencia(informacao model.AtivoInformacao) (response.AtivoEficienciaResponse, error) {
	eficiencia := mapper.ToAtivoEficiencia(informacao)
	eficienciaPersisted, err := s.eficienciaRep.CreateEficiencia(eficiencia)
	if err != nil {
		return response.AtivoEficienciaResponse{}, err
	}

	return mapper.ToAtivoEficienciaResponse(eficienciaPersisted), err
}

func (s *AtivoInformacaoService) createRentabilidade(informacao model.AtivoInformacao) (response.AtivoRentabilidadeResponse, error) {
	rentabilidade := mapper.ToAtivoRentabilidade(informacao)
	rentabilidadePersisted, err := s.rentabilidadeRep.CreateRentabilidade(rentabilidade)
	if err != nil {
		return response.AtivoRentabilidadeResponse{}, err
	}

	return mapper.ToAtivoRentabilidadeResponse(rentabilidadePersisted), err
}

func (s *AtivoInformacaoService) createDividendo(informacao model.AtivoInformacao, precoAtual float64) (response.AtivoDividendoResponse, error) {

	dataInicial := informacao.DataInformacao.AddDate(-1, 0, 0)
	log.Println("Data Inicial: ", dataInicial)
	dividendos, err := s.dividendosRep.GetDividendosByCodigoAndIntervalo(informacao.Codigo, dataInicial, informacao.DataInformacao)
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
	dividendoPersisted, err := s.ativoDividendoRep.CreateDividendo(dividendo)
	if err != nil {
		return response.AtivoDividendoResponse{}, err
	}

	return mapper.ToAtivoDividendoResponse(dividendoPersisted), err
}
