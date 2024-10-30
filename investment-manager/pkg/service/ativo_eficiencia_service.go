package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoEficienciaService struct {
	repo *repository.AtivoEficienciaRepository
}

func NewAtivoEficienciaService(repo *repository.AtivoEficienciaRepository) *AtivoEficienciaService {
	return &AtivoEficienciaService{repo: repo}
}

func (s *AtivoEficienciaService) CreateAtivoEficiencia(informacao model.AtivoInformacao) (response.AtivoEficienciaResponse, error) {
	eficiencia := mapper.ToAtivoEficiencia(informacao)
	eficienciaPersisted, err := s.repo.CreateEficiencia(eficiencia)
	if err != nil {
		return response.AtivoEficienciaResponse{}, err
	}

	return mapper.ToAtivoEficienciaResponse(eficienciaPersisted), err
}
