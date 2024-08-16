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
	repo            *repository.CarteiraRepository
	ativoRepo       *repository.CarteiraAtivoRepository
	consolidacaoRep *repository.ConsolidacaoRepository
	cotacaoRep      *repository.CotacaoHistoricoRepository
}

func NewCarteiraService(
	repo *repository.CarteiraRepository,
	ativoRepo *repository.CarteiraAtivoRepository,
	consolidacaoRepo *repository.ConsolidacaoRepository,
	cotacaoRep *repository.CotacaoHistoricoRepository,
) *CarteiraService {
	return &CarteiraService{
		repo:            repo,
		ativoRepo:       ativoRepo,
		consolidacaoRep: consolidacaoRepo,
		cotacaoRep:      cotacaoRep,
	}
}

func (s *CarteiraService) getAtivosByCarteiraID(carteiraID uint) []response.CarteiraAtivoResponse {
	ativos, err := s.ativoRepo.GetAtivosByCarteiraID(carteiraID)
	if err != nil {
		ativos = []model.CarteiraAtivo{}
	}
	var ativosResponse []response.CarteiraAtivoResponse
	for _, ativo := range ativos {
		consolidacao, _ := s.consolidacaoRep.GetConsolidacaoByCodigo(ativo.Codigo)
		cotacao, _ := s.cotacaoRep.GetCotacaoMoreRecentByCodigo(ativo.Codigo)
		ativosResponse = append(ativosResponse, mapper.ToCarteiraAtivoResponse(ativo, consolidacao, cotacao))
	}
	return ativosResponse
}

func (s *CarteiraService) GetAllCarteiras() ([]response.CarteiraResponse, error) {
	var carteirasResponse []response.CarteiraResponse

	carteiras, err := s.repo.GetAllCarteiras()

	if err != nil {
		return []response.CarteiraResponse{}, err
	}

	for _, value := range carteiras {
		carteiraResponse := mapper.ToCarteiraResponse(value)
		carteiraResponse.Ativos = s.getAtivosByCarteiraID(value.ID)
		carteirasResponse = append(carteirasResponse, carteiraResponse)
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
	carteiraResponse := mapper.ToCarteiraResponse(carteira)
	carteiraResponse.Ativos = s.getAtivosByCarteiraID(carteira.ID)
	return carteiraResponse, err
}

func (s *CarteiraService) CreateCarteira(carteiraRequest request.CarteiraRequest) (response.CarteiraResponse, error) {
	carteira := mapper.ToCarteira(carteiraRequest)
	carteiraCreated, err := s.repo.CreateCarteira(*carteira)
	if err != nil {
		return response.CarteiraResponse{}, err
	}
	return mapper.ToCarteiraResponse(carteiraCreated), err
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

	carteiraResponse := mapper.ToCarteiraResponse(carteiraUpdated)
	carteiraResponse.Ativos = s.getAtivosByCarteiraID(carteiraUpdated.ID)
	return carteiraResponse, err
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
	return mapper.ToCarteiraAtivoSimpleResponse(ativoCreated), err

}

func (s *CarteiraService) RemoverAtivo(carteiraId uint, codigo string) error {

	ativo, err := s.ativoRepo.GetAtivoByCodigoAndCarteiraID(codigo, carteiraId)
	if err != nil {
		return err
	}

	return s.ativoRepo.DeleteCarteiraAtivo(ativo.ID)

}
