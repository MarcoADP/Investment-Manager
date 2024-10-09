package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type AcaoBrService struct {
	repo      *repository.AcaoBrRepository
	setorRepo *repository.SetorRepository
}

func NewAcaoBrService(repo *repository.AcaoBrRepository, setorRepo *repository.SetorRepository) *AcaoBrService {
	return &AcaoBrService{repo: repo, setorRepo: setorRepo}
}

func (s *AcaoBrService) GetAllAcaoBrs() ([]response.AcaoBrResponse, error) {
	acoesBr, err := s.repo.GetAllAcaoBrs()

	if err != nil {
		return []response.AcaoBrResponse{}, err
	}

	return s.toAcaoBrResponseArray(acoesBr), err
}

func (s *AcaoBrService) GetAcaoBrByCodigo(codigo string) (response.AcaoBrResponse, error) {
	acao, err := s.repo.GetAcaoBrByCodigo(codigo)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	setor, err := s.getSetor(acao.SetorId)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	return mapper.ToAcaoBrResponse(acao, s.getSetorNome(setor)), err
}

func (s *AcaoBrService) GetAcoesBySetor(setorId uint) ([]response.AcaoBrResponse, error) {
	acoes, err := s.repo.GetAcoesBySetor(setorId)
	if err != nil {
		return []response.AcaoBrResponse{}, err
	}
	return s.toAcaoBrResponseArray(acoes), err
}

func (s *AcaoBrService) CreateAcaoBr(acaoRequest request.AcaoBrRequest) (response.AcaoBrResponse, error) {
	setor, err := s.getOrCreateSetor(acaoRequest.Setor)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	acao := mapper.ToAcaoBr(acaoRequest, setor.ID)
	acaoCreated, err := s.repo.CreateAcaoBr(*acao)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}
	return mapper.ToAcaoBrResponse(acaoCreated, s.getSetorNome(setor)), err
}

func (s *AcaoBrService) UpdateAcaoBr(codigo string, acaoRequest request.AcaoBrRequest) (response.AcaoBrResponse, error) {

	acao, err := s.repo.GetAcaoBrByCodigo(codigo)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}
	var setor model.Setor
	if acao.SetorId == 0 {
		setor, err = s.getOrCreateSetor(acaoRequest.Setor)
		if err == nil {
			acao.SetorId = setor.ID
		}
	} else {
		setor, err = s.getSetor(acao.SetorId)
		if err != nil {
			return response.AcaoBrResponse{}, err
		}
	}

	acaoMapped := mapper.UpdateAcaoBr(acao, acaoRequest)
	acaoUpdated, err := s.repo.UpdateAcaoBr(acaoMapped)
	if err != nil {
		return response.AcaoBrResponse{}, err
	}

	return mapper.ToAcaoBrResponse(acaoUpdated, s.getSetorNome(setor)), err
}

func (s *AcaoBrService) DeleteAcaoBr(codigo string) error {
	return s.repo.DeleteAcaoBr(codigo)
}

func (s *AcaoBrService) getSetorNome(setor model.Setor) string {
	if setor == (model.Setor{}) {
		return "Setor não informado"
	}
	return setor.Nome
}

func (s *AcaoBrService) getSetor(setorId uint) (model.Setor, error) {
	if setorId == 0 {
		return model.Setor{}, nil
	}
	return s.setorRepo.GetSetorByID(setorId)
}

func (s *AcaoBrService) toAcaoBrResponseArray(acoesBr []model.AcaoBr) []response.AcaoBrResponse {
	var acoesResponse []response.AcaoBrResponse
	for _, value := range acoesBr {
		var setor model.Setor
		if value.SetorId != 0 {
			setor, _ = s.getSetor(value.SetorId)
		}
		acoesResponse = append(acoesResponse, mapper.ToAcaoBrResponse(value, s.getSetorNome(setor)))
	}
	return acoesResponse
}

func (s *AcaoBrService) getOrCreateSetor(setorNome string) (model.Setor, error) {
	setor, err := s.setorRepo.GetSetorByNome(setorNome)
	if err != nil {
		setor = *model.NewSetor(setorNome)
		return s.setorRepo.CreateSetor(setor)
	}
	return setor, err
}
