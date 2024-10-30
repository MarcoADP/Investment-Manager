package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoValuationService struct {
	repo *repository.AtivoValuationRepository
}

func NewAtivoValuationService(repo *repository.AtivoValuationRepository) *AtivoValuationService {
	return &AtivoValuationService{repo: repo}
}

func (s *AtivoValuationService) CreateAtivoValuation(informacao model.AtivoInformacao, valor float64) (response.AtivoValuationResponse, error) {
	valuation := mapper.ToAtivoValuation(informacao, valor)
	valuationPersisted, err := s.repo.CreateValuation(valuation)
	if err != nil {
		return response.AtivoValuationResponse{}, err
	}

	return mapper.ToAtivoValuationResponse(valuationPersisted), err
}

func (s *AtivoValuationService) GetAtivoValuationByInformacaoID(ativoInformacaoId uint) (response.AtivoValuationResponse, error) {
	valuation, err := s.repo.GetValuationByAtivoInformacao(ativoInformacaoId)
	if err != nil {
		return response.AtivoValuationResponse{}, err
	}

	return mapper.ToAtivoValuationResponse(valuation), err
}
