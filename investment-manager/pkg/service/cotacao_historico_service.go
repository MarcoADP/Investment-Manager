package service

import (
	"log"
	"time"

	"github.com/MarcoADP/Investment-Manager/internal/brapi"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type CotacaoHistoricoService struct {
	repo *repository.CotacaoHistoricoRepository
}

func NewCotacaoHistoricoService(repo *repository.CotacaoHistoricoRepository) *CotacaoHistoricoService {
	return &CotacaoHistoricoService{repo: repo}
}

func (s *CotacaoHistoricoService) GetAllCotacoes() ([]response.CotacaoHistoricoResponse, error) {
	cotacoes, err := s.repo.GetAllCotacoes()

	if err != nil {
		return []response.CotacaoHistoricoResponse{}, err
	}

	return mapper.ToCotacaoHistoricoResponseArray(cotacoes), err
}

func (s *CotacaoHistoricoService) GetCotacoesByCodigo(codigo string) ([]response.CotacaoHistoricoResponse, error) {
	cotacoes, err := s.repo.GetCotacoesByCodigo(codigo)

	if err != nil {
		return []response.CotacaoHistoricoResponse{}, err
	}

	return mapper.ToCotacaoHistoricoResponseArray(cotacoes), err
}

func (s *CotacaoHistoricoService) GetCotacaoByCodigoAndData(codigo string, data string) (response.CotacaoHistoricoResponse, error) {

	date, err := time.Parse("02/01/2006", data)
	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}

	cotacao, err := s.repo.GetCotacoesByCodigoAndData(codigo, date)

	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}

	return mapper.ToCotacaoHistoricoResponse(cotacao), err
}

func (s *CotacaoHistoricoService) GetCotacaoMoreRecentByCodigo(codigo string) (response.CotacaoHistoricoResponse, error) {

	cotacao, err := s.repo.GetCotacaoMoreRecentByCodigo(codigo)

	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}

	return mapper.ToCotacaoHistoricoResponse(cotacao), err
}

func (s *CotacaoHistoricoService) CreateCotacao(cotacaoRequest request.CotacaoHistoricoRequest) (response.CotacaoHistoricoResponse, error) {
	cotacao, err := mapper.ToCotacaoHistorico(cotacaoRequest)
	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}

	cotacaoExisted, err := s.repo.GetCotacoesByCodigoAndData(cotacao.Codigo, cotacao.Data)
	var cotacaoPersisted model.CotacaoHistorico
	if err == nil {
		cotacao.ID = cotacaoExisted.ID
		cotacaoPersisted, err = s.repo.UpdateCotacaoHistorico(*cotacao)
	} else {
		cotacaoPersisted, err = s.repo.CreateCotacao(*cotacao)
	}

	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}
	return mapper.ToCotacaoHistoricoResponse(cotacaoPersisted), err
}

func (s *CotacaoHistoricoService) DeleteCotacao(id uint) error {
	return s.repo.DeleteCotacaoHistorico(id)
}

func (s *CotacaoHistoricoService) GetCotacaoExterno(codigo string) (response.CotacaoHistoricoResponse, error) {

	cotacaoResponse, err := brapi.GetCotacaoBrapi(codigo, "18ftFmyA1zQzEzhwQJmync")
	if err != nil {
		log.Fatalf("Failed to get quote: %v", err)
		return response.CotacaoHistoricoResponse{}, err
	}

	cotacao, err := mapper.ToCotacaoHistoricoFromBrapi(*cotacaoResponse)
	if err != nil {
		log.Fatalf("Failed to get quote: %v", err)
		return response.CotacaoHistoricoResponse{}, err
	}

	cotacaoExisted, err := s.repo.GetCotacoesByCodigoAndData(cotacao.Codigo, cotacao.Data)
	var cotacaoPersisted model.CotacaoHistorico
	if err == nil {
		cotacao.ID = cotacaoExisted.ID
		cotacaoPersisted, err = s.repo.UpdateCotacaoHistorico(*cotacao)
	} else {
		cotacaoPersisted, err = s.repo.CreateCotacao(*cotacao)
	}

	if err != nil {
		return response.CotacaoHistoricoResponse{}, err
	}
	return mapper.ToCotacaoHistoricoResponse(cotacaoPersisted), err
}
