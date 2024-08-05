package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AcaoBrService struct {
	repo *repository.AcaoBrRepository
}

func NewAcaoBrService(repo *repository.AcaoBrRepository) *AcaoBrService {
	return &AcaoBrService{repo: repo}
}

func (s *AcaoBrService) GetAllAcaoBrs() ([]response.AcaoBrResponse, error) {
	var acoesResponse []response.AcaoBrResponse

	acoesBr, err := s.repo.GetAllAcaoBrs()

	if err != nil {
		return []response.AcaoBrResponse{}, err
	}

	for _, value := range acoesBr {
		acoesResponse = append(acoesResponse, mapper.ToAcaoBrResponse(value))
	}

	return acoesResponse, err
}

func (s *AcaoBrService) GetAcaoBrByID(id uint) (response.AcaoBrResponse, error) {
	acao, err := s.repo.GetAcaoBrByID(id)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}
	return mapper.ToAcaoBrResponse(acao), err
}

func (s *AcaoBrService) CreateAcaoBr(acaoRequest request.AcaoBrRequest) (response.AcaoBrResponse, error) {
	acao := mapper.ToAcaoBr(acaoRequest)
	acaoCreated, err := s.repo.CreateAcaoBr(*acao)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}
	return mapper.ToAcaoBrResponse(acaoCreated), err
}

func (s *AcaoBrService) UpdateAcaoBr(id uint, acaoRequest request.AcaoBrRequest) (response.AcaoBrResponse, error) {

	acao, err := s.repo.GetAcaoBrByID(id)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	acaoMapped := mapper.UpdateAcaoBr(acao, acaoRequest)
	acaoUpdated, err := s.repo.UpdateAcaoBr(acaoMapped)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	return mapper.ToAcaoBrResponse(acaoUpdated), err
}

func (s *AcaoBrService) DeleteAcaoBr(id uint) error {
	return s.repo.DeleteAcaoBr(id)
}
