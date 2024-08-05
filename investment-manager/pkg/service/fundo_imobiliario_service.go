package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type FundoImobiliarioService struct {
	repo *repository.FundoImobiliarioRepository
}

func NewFundoImobiliarioService(repo *repository.FundoImobiliarioRepository) *FundoImobiliarioService {
	return &FundoImobiliarioService{repo: repo}
}

func (s *FundoImobiliarioService) GetAllFundosImobiliarios() ([]response.FundoImobiliarioResponse, error) {
	var fundosImobiliariosResponse []response.FundoImobiliarioResponse

	fundosImobiliarios, err := s.repo.GetAllFundosImobiliarios()

	if err != nil {
		return []response.FundoImobiliarioResponse{}, err
	}

	for _, value := range fundosImobiliarios {
		fundosImobiliariosResponse = append(fundosImobiliariosResponse, mapper.ToFundoImobiliarioResponse(value))
	}

	return fundosImobiliariosResponse, err
}

func (s *FundoImobiliarioService) GetFundoImobiliarioByID(id uint) (response.FundoImobiliarioResponse, error) {
	fundoImobiliario, err := s.repo.GetFundoImobiliarioByID(id)
	if err != nil {
		return response.FundoImobiliarioResponse{}, err
	}
	return mapper.ToFundoImobiliarioResponse(fundoImobiliario), err
}

func (s *FundoImobiliarioService) CreateFundoImobiliario(fundoImobiliarioRequest request.FundoImobiliarioRequest) (response.FundoImobiliarioResponse, error) {
	fundoImobiliario := mapper.ToFundoImobiliario(fundoImobiliarioRequest)
	fundoImobiliarioCreated, err := s.repo.CreateFundoImobiliario(*fundoImobiliario)
	if err != nil {
		return response.FundoImobiliarioResponse{}, err
	}
	return mapper.ToFundoImobiliarioResponse(fundoImobiliarioCreated), err
}

func (s *FundoImobiliarioService) UpdateFundoImobiliario(id uint, fundoImobiliarioRequest request.FundoImobiliarioRequest) (response.FundoImobiliarioResponse, error) {

	fundoImobiliario, err := s.repo.GetFundoImobiliarioByID(id)
	if err != nil {
		return response.FundoImobiliarioResponse{}, err
	}

	fundoImobiliarioMapped := mapper.UpdateFundoImobiliario(fundoImobiliario, fundoImobiliarioRequest)
	fundoImobiliarioUpdated, err := s.repo.UpdateFundoImobiliario(fundoImobiliarioMapped)
	if err != nil {
		return response.FundoImobiliarioResponse{}, err
	}

	return mapper.ToFundoImobiliarioResponse(fundoImobiliarioUpdated), err
}

func (s *FundoImobiliarioService) DeleteFundoImobiliario(id uint) error {
	return s.repo.DeleteFundoImobiliario(id)
}
