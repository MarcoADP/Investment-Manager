package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type GrahamFormulaService struct {
	repo *repository.GrahamFormulaRepository
}

func NewGrahamFormulaService(repo *repository.GrahamFormulaRepository) *GrahamFormulaService {
	return &GrahamFormulaService{
		repo: repo,
	}
}

func (s *GrahamFormulaService) GetGrahamFormulaByCodigo(codigo string) ([]response.GrahamFormulaResponse, error) {
	data, err := s.repo.GetGrahamFormulaByCodigo(codigo)

	if err != nil {
		return []response.GrahamFormulaResponse{}, err
	}

	return mapper.ToGrahamFormulaResponseArray(data), err
}

func (s *GrahamFormulaService) GetGrahamFormulaMoreRecentByCodigo(codigo string) (response.GrahamFormulaResponse, error) {

	data, err := s.repo.GetGrahamFormulaMoreRecentByCodigo(codigo)

	if err != nil {
		return response.GrahamFormulaResponse{}, err
	}

	return mapper.ToGrahamFormulaResponse(data), err
}

func (s *GrahamFormulaService) CreateGrahamFormula(request request.GrahamFormulaRequest) (response.GrahamFormulaResponse, error) {

	data, err := mapper.ToGrahamFormula(request)
	if err != nil {
		return response.GrahamFormulaResponse{}, err
	}

	dataPersisted, err := s.repo.CreateGrahamFormula(*data)
	if err != nil {
		return response.GrahamFormulaResponse{}, err
	}
	return mapper.ToGrahamFormulaResponse(dataPersisted), err
}

func (s *GrahamFormulaService) DeleteGrahamFormula(id uint) error {
	return s.repo.DeleteGrahamFormula(id)
}
