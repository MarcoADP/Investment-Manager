package service

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type MovimentacaoService struct {
	repo *repository.MovimentacaoRepository
}

func NewMovimentacaoService(repo *repository.MovimentacaoRepository) *MovimentacaoService {
	return &MovimentacaoService{repo: repo}
}

func (s *MovimentacaoService) GetAllMovimentacoes() ([]response.MovimentacaoResponse, error) {
	var movimentacoesResponse []response.MovimentacaoResponse

	movimentacoes, err := s.repo.GetAllMovimentacoes()

	if err != nil {
		return []response.MovimentacaoResponse{}, err
	}

	for _, value := range movimentacoes {
		movimentacoesResponse = append(movimentacoesResponse, mapper.ToMovimentacaoResponse(value))
	}

	if movimentacoesResponse == nil {
		movimentacoesResponse = []response.MovimentacaoResponse{}
	}

	return movimentacoesResponse, err
}

func (s *MovimentacaoService) GetAllMovimentacaoByCodigo(codigo string) ([]response.MovimentacaoResponse, error) {
	var movimentacoesResponse []response.MovimentacaoResponse

	movimentacoes, err := s.repo.GetMovimentacaoByCodigo(codigo)

	if err != nil {
		return []response.MovimentacaoResponse{}, err
	}

	for _, value := range movimentacoes {
		movimentacoesResponse = append(movimentacoesResponse, mapper.ToMovimentacaoResponse(value))
	}

	if movimentacoesResponse == nil {
		movimentacoesResponse = []response.MovimentacaoResponse{}
	}

	return movimentacoesResponse, err
}

func (s *MovimentacaoService) GetMovimentacaoByID(id uint) (response.MovimentacaoResponse, error) {
	movimentacao, err := s.repo.GetMovimentacaoByID(id)
	if err != nil {
		return response.MovimentacaoResponse{}, err
	}
	return mapper.ToMovimentacaoResponse(movimentacao), err
}

func (s *MovimentacaoService) CreateMovimentacao(movimentacaoRequest request.MovimentacaoRequest, operacao string) (response.MovimentacaoResponse, error) {
	movimentacao, err := mapper.ToMovimentacao(movimentacaoRequest, operacao)
	if err != nil {
		return response.MovimentacaoResponse{}, err
	}

	movimentacaoCreated, err := s.repo.CreateMovimentacao(*movimentacao)
	if err != nil {
		return response.MovimentacaoResponse{}, err
	}
	return mapper.ToMovimentacaoResponse(movimentacaoCreated), err
}

func (s *MovimentacaoService) DeleteMovimentacao(id uint) error {
	return s.repo.DeleteMovimentacao(id)
}

func (s *MovimentacaoService) GetMovimentacoesGroupedByCodigo() (map[string][]response.MovimentacaoResponse, error) {
	movimentacoesAgrupados := make(map[string][]response.MovimentacaoResponse)
	movimentacoes, err := s.GetAllMovimentacoes()

	if err != nil {
		return movimentacoesAgrupados, err
	}

	for _, movimentacao := range movimentacoes {
		movimentacoesAgrupados[movimentacao.Codigo] = append(movimentacoesAgrupados[movimentacao.Codigo], movimentacao)
	}

	return movimentacoesAgrupados, err
}
