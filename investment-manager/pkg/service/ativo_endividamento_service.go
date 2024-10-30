package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoEndividamentoService struct {
	repo *repository.AtivoEndividamentoRepository
}

func NewAtivoEndividamentoService(repo *repository.AtivoEndividamentoRepository) *AtivoEndividamentoService {
	return &AtivoEndividamentoService{repo: repo}
}

func (s *AtivoEndividamentoService) CreateAtivoEndividamento(informacao model.AtivoInformacao) (response.AtivoEndividamentoResponse, error) {
	endividamento := mapper.ToAtivoEndividamento(informacao)
	endividamentoPersisted, err := s.repo.CreateEndividamento(endividamento)
	if err != nil {
		return response.AtivoEndividamentoResponse{}, err
	}

	return mapper.ToAtivoEndividamentoResponse(endividamentoPersisted), err
}
