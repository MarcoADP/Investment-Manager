package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type BdrService struct {
	repo *repository.BdrRepository
}

func NewBdrService(repo *repository.BdrRepository) *BdrService {
	return &BdrService{repo: repo}
}

func (s *BdrService) GetAllBdrs() ([]response.BdrResponse, error) {
	var bdrsResponse []response.BdrResponse

	bdrs, err := s.repo.GetAllBdrs()

	if err != nil {
		return []response.BdrResponse{}, err
	}

	for _, value := range bdrs {
		bdrsResponse = append(bdrsResponse, mapper.ToBdrResponse(value))
	}

	if bdrsResponse == nil {
		bdrsResponse = []response.BdrResponse{}
	}

	return bdrsResponse, err
}

func (s *BdrService) GetBdrByID(id uint) (response.BdrResponse, error) {
	bdr, err := s.repo.GetBdrByID(id)
	if err != nil {
		return response.BdrResponse{}, err
	}
	return mapper.ToBdrResponse(bdr), err
}

func (s *BdrService) CreateBdr(bdrRequest request.BdrRequest) (response.BdrResponse, error) {
	bdr := mapper.ToBdr(bdrRequest)
	bdrCreated, err := s.repo.CreateBdr(*bdr)
	if err != nil {
		return response.BdrResponse{}, err
	}
	return mapper.ToBdrResponse(bdrCreated), err
}

func (s *BdrService) UpdateBdr(id uint, bdrRequest request.BdrRequest) (response.BdrResponse, error) {

	bdr, err := s.repo.GetBdrByID(id)
	if err != nil {
		return response.BdrResponse{}, err
	}

	bdrMapped := mapper.UpdateBdr(bdr, bdrRequest)
	bdrUpdated, err := s.repo.UpdateBdr(bdrMapped)
	if err != nil {
		return response.BdrResponse{}, err
	}

	return mapper.ToBdrResponse(bdrUpdated), err
}

func (s *BdrService) DeleteBdr(id uint) error {
	return s.repo.DeleteBdr(id)
}
