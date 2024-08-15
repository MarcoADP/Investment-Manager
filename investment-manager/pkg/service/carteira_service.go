package service

import (
	"errors"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type CarteiraService struct {
	repo      *repository.CarteiraRepository
	ativoRepo *repository.CarteiraAtivoRepository
}

func NewCarteiraService(repo *repository.CarteiraRepository, ativoRepo *repository.CarteiraAtivoRepository) *CarteiraService {
	return &CarteiraService{repo: repo, ativoRepo: ativoRepo}
}

func (s *CarteiraService) getAtivosByCarteiraID(carteiraID uint) []model.CarteiraAtivo {
	ativos, err := s.ativoRepo.GetAtivosByCarteiraID(carteiraID)
	if err != nil {
		ativos = []model.CarteiraAtivo{}
	}
	return ativos
}

func (s *CarteiraService) GetAllCarteiras() ([]response.CarteiraResponse, error) {
	var carteirasResponse []response.CarteiraResponse

	bdrs, err := s.repo.GetAllCarteiras()

	if err != nil {
		return []response.CarteiraResponse{}, err
	}

	for _, value := range bdrs {
		ativos := s.getAtivosByCarteiraID(value.ID)
		carteirasResponse = append(carteirasResponse, mapper.ToCarteiraResponse(value, ativos))
	}

	if carteirasResponse == nil {
		carteirasResponse = []response.CarteiraResponse{}
	}

	return carteirasResponse, err
}

func (s *CarteiraService) GetCarteiraByID(id uint) (response.CarteiraResponse, error) {
	carteira, err := s.repo.GetCarteiraByID(id)
	if err != nil {
		return response.CarteiraResponse{}, err
	}
	ativos := s.getAtivosByCarteiraID(carteira.ID)
	return mapper.ToCarteiraResponse(carteira, ativos), err
}

func (s *CarteiraService) CreateCarteira(carteiraRequest request.CarteiraRequest) (response.CarteiraResponse, error) {
	carteira := mapper.ToCarteira(carteiraRequest)
	carteiraCreated, err := s.repo.CreateCarteira(*carteira)
	if err != nil {
		return response.CarteiraResponse{}, err
	}
	return mapper.ToCarteiraResponse(carteiraCreated, []model.CarteiraAtivo{}), err
}

func (s *CarteiraService) UpdateCarteira(id uint, carteiraRequest request.CarteiraRequest) (response.CarteiraResponse, error) {

	carteira, err := s.repo.GetCarteiraByID(id)
	if err != nil {
		return response.CarteiraResponse{}, err
	}

	carteiraMapped := mapper.UpdateCarteira(carteira, carteiraRequest)
	carteiraUpdated, err := s.repo.UpdateCarteira(carteiraMapped)
	if err != nil {
		return response.CarteiraResponse{}, err
	}

	ativos := s.getAtivosByCarteiraID(carteiraUpdated.ID)
	return mapper.ToCarteiraResponse(carteiraUpdated, ativos), err
}

func (s *CarteiraService) DeleteCarteira(id uint) error {
	ativos := s.getAtivosByCarteiraID(id)
	for _, ativo := range ativos {
		s.ativoRepo.DeleteCarteiraAtivo(ativo.ID)
	}
	return s.repo.DeleteCarteira(id)
}

func (s *CarteiraService) AddAtivo(carteiraId uint, ativoRequest request.CarteiraAtivoRequest) (response.CarteiraAtivoResponse, error) {

	_, err := s.ativoRepo.GetAtivoByCodigoAndCarteiraID(ativoRequest.Codigo, carteiraId)
	if err == nil {
		return response.CarteiraAtivoResponse{}, errors.New("ativo já foi adicionado nessa carteira")
	}

	ativo := mapper.ToCarteiraAtivo(carteiraId, ativoRequest)
	ativoCreated, err := s.ativoRepo.CreateCarteiraAtivo(*ativo)
	if err != nil {
		return response.CarteiraAtivoResponse{}, err
	}
	return mapper.ToCarteiraAtivoResponse(ativoCreated), err

}

func (s *CarteiraService) RemoverAtivo(carteiraId uint, codigo string) error {

	ativo, err := s.ativoRepo.GetAtivoByCodigoAndCarteiraID(codigo, carteiraId)
	if err != nil {
		return err
	}

	return s.ativoRepo.DeleteCarteiraAtivo(ativo.ID)

}
