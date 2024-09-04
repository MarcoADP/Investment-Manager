package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type DividendoService struct {
	repo *repository.DividendoRepository
}

func NewDividendoService(repo *repository.DividendoRepository) *DividendoService {
	return &DividendoService{repo: repo}
}

func (s *DividendoService) GetAllDividendos() ([]response.DividendoResponse, error) {
	var dividendosResponse []response.DividendoResponse

	dividendos, err := s.repo.GetAllDividendos()

	if err != nil {
		return []response.DividendoResponse{}, err
	}

	for _, value := range dividendos {
		dividendosResponse = append(dividendosResponse, mapper.ToDividendoResponse(value))
	}

	if dividendosResponse == nil {
		dividendosResponse = []response.DividendoResponse{}
	}

	return dividendosResponse, err
}

func (s *DividendoService) GetDividendoByID(id uint) (response.DividendoResponse, error) {
	dividendo, err := s.repo.GetDividendoByID(id)
	if err != nil {
		return response.DividendoResponse{}, err
	}
	return mapper.ToDividendoResponse(dividendo), err
}

func (s *DividendoService) CreateDividendo(dividendoRequest request.DividendoRequest) (response.DividendoResponse, error) {
	dividendo, err := mapper.ToDividendo(dividendoRequest)
	if err != nil {
		return response.DividendoResponse{}, err
	}

	dividendoCreated, err := s.repo.CreateDividendo(*dividendo)
	if err != nil {
		return response.DividendoResponse{}, err
	}
	return mapper.ToDividendoResponse(dividendoCreated), err
}

func (s *DividendoService) DeleteDividendo(id uint) error {
	return s.repo.DeleteDividendo(id)
}

func (s *DividendoService) GetDividendosByCodigo(codigo string) ([]response.DividendoResponse, error) {
	dividendos, err := s.repo.GetDividendosByCodigo(codigo)

	if err != nil {
		return []response.DividendoResponse{}, err
	}

	return mapper.ToDividendoResponseArray(dividendos), err
}
