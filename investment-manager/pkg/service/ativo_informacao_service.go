package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoInformacaoService struct {
	repo                      *repository.AtivoInformacaoRepository
	cotacaoRep                *repository.CotacaoHistoricoRepository
	ativoValuationService     *AtivoValuationService
	ativoEndividamentoService *AtivoEndividamentoService
	ativoEficienciaService    *AtivoEficienciaService
	ativoRentabilidadeService *AtivoRentabilidadeService
	ativoDividendoService     *AtivoDividendoService
}

func NewAtivoInformacaoService(repo *repository.AtivoInformacaoRepository,
	cotacaoRep *repository.CotacaoHistoricoRepository,
	ativoValuationService *AtivoValuationService,
	ativoEndividamentoService *AtivoEndividamentoService,
	ativoEficienciaService *AtivoEficienciaService,
	ativoRentabilidadeService *AtivoRentabilidadeService,
	ativoDividendoService *AtivoDividendoService,
) *AtivoInformacaoService {
	return &AtivoInformacaoService{
		repo:                      repo,
		cotacaoRep:                cotacaoRep,
		ativoValuationService:     ativoValuationService,
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

	var responses []response.AtivoInformacaoResponse
	for _, value := range informacoes {
		response := mapper.ToAtivoInformacaoResponse(value)
		response = s.GetIndicadores(response)
		responses = append(responses, response)
	}

	if responses == nil {
		responses = []response.AtivoInformacaoResponse{}
	}

	return responses, err

}

func (s *AtivoInformacaoService) GetInformacaoMoreRecentByCodigo(codigo string) (response.AtivoInformacaoResponse, error) {

	informacao, err := s.repo.GetInformacaoMoreRecentByCodigo(codigo)

	if err != nil {
		return response.AtivoInformacaoResponse{}, err
	}

	response := mapper.ToAtivoInformacaoResponse(informacao)
	response = s.GetIndicadores(response)

	return response, err
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

	informacaoResponse.Valuation, _ = s.ativoValuationService.CreateAtivoValuation(informacaoPersisted, cotacao.Valor)
	informacaoResponse.Endividamento, _ = s.ativoEndividamentoService.CreateAtivoEndividamento(informacaoPersisted)
	informacaoResponse.Eficiencia, _ = s.ativoEficienciaService.CreateAtivoEficiencia(informacaoPersisted)
	informacaoResponse.Rentabilidade, _ = s.ativoRentabilidadeService.CreateAtivoRentabilidade(informacaoPersisted)
	informacaoResponse.Dividendo, _ = s.ativoDividendoService.CreateAtivoDividendo(informacaoPersisted, cotacao.Valor)

	return informacaoResponse, err
}

func (s *AtivoInformacaoService) DeleteInformacao(id uint) error {
	return s.repo.DeleteInformacao(id)
}

func (s *AtivoInformacaoService) GetIndicadores(ativoInformacaoResponse response.AtivoInformacaoResponse) response.AtivoInformacaoResponse {
	ativoInformacaoResponse.Valuation, _ = s.ativoValuationService.GetAtivoValuationByInformacaoID(ativoInformacaoResponse.ID)
	return ativoInformacaoResponse
}
