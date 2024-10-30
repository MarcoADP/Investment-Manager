package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AtivoRentabilidadeService struct {
	repo *repository.AtivoRentabilidadeRepository
}

func NewAtivoRentabilidadeService(repo *repository.AtivoRentabilidadeRepository) *AtivoRentabilidadeService {
	return &AtivoRentabilidadeService{repo: repo}
}

func (s *AtivoRentabilidadeService) CreateAtivoRentabilidade(informacao model.AtivoInformacao) (response.AtivoRentabilidadeResponse, error) {
	rentabilidade := mapper.ToAtivoRentabilidade(informacao)
	rentabilidadePersisted, err := s.repo.CreateRentabilidade(rentabilidade)
	if err != nil {
		return response.AtivoRentabilidadeResponse{}, err
	}

	return mapper.ToAtivoRentabilidadeResponse(rentabilidadePersisted), err
}
