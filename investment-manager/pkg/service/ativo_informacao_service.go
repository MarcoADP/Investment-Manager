package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoInformacaoService struct {
	repo                      *repository.AtivoInformacaoRepository
	cotacaoRep                *repository.CotacaoHistoricoRepository
	valuationRep              *repository.AtivoValuationRepository
	ativoEndividamentoService *AtivoEndividamentoService
	ativoEficienciaService    *AtivoEficienciaService
	ativoRentabilidadeService *AtivoRentabilidadeService
	ativoDividendoService     *AtivoDividendoService
}

func NewAtivoInformacaoService(repo *repository.AtivoInformacaoRepository,
	cotacaoRep *repository.CotacaoHistoricoRepository,
	valuationRep *repository.AtivoValuationRepository,
	ativoEndividamentoService *AtivoEndividamentoService,
	ativoEficienciaService *AtivoEficienciaService,
	ativoRentabilidadeService *AtivoRentabilidadeService,
	ativoDividendoService *AtivoDividendoService,
) *AtivoInformacaoService {
	return &AtivoInformacaoService{
		repo:                      repo,
		cotacaoRep:                cotacaoRep,
		valuationRep:              valuationRep,
		ativoEndividamentoService: ativoEndividamentoService,
		ativoEficienciaService:    ativoEficienciaService,
		ativoRentabilidadeService: ativoRentabilidadeService,
		ativoDividendoService:     ativoDividendoService,
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

	informacaoResponse.Endividamento, _ = s.ativoEndividamentoService.CreateAtivoEndividamento(informacaoPersisted)
	informacaoResponse.Eficiencia, _ = s.ativoEficienciaService.CreateAtivoEficiencia(informacaoPersisted)
	informacaoResponse.Rentabilidade, _ = s.ativoRentabilidadeService.CreateAtivoRentabilidade(informacaoPersisted)
	informacaoResponse.Dividendo, _ = s.ativoDividendoService.CreateAtivoDividendo(informacaoPersisted, cotacao.Valor)

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
