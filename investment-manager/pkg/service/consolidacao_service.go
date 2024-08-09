package service

import (
	"strings"

	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"github.com/MarcoADP/Investment-Manager/pkg/db/repository"
	"github.com/MarcoADP/Investment-Manager/pkg/mapper"
)

type ConsolidacaoService struct {
	repo                *repository.ConsolidacaoRepository
	movimentacaoService *MovimentacaoService
}

func NewConsolidacaoService(repo *repository.ConsolidacaoRepository, movimentoService *MovimentacaoService) *ConsolidacaoService {
	return &ConsolidacaoService{repo: repo, movimentacaoService: movimentoService}
}

func (s *ConsolidacaoService) GetAllConsolidacoes() ([]response.ConsolidacaoResponse, error) {
	var consolidacoesResponse []response.ConsolidacaoResponse

	consolidacoes, err := s.repo.GetAllConsolidacoes()

	if err != nil {
		return []response.ConsolidacaoResponse{}, err
	}

	for _, value := range consolidacoes {
		consolidacoesResponse = append(consolidacoesResponse, mapper.ToConsolidacaoResponse(value))
	}

	if consolidacoesResponse == nil {
		consolidacoesResponse = []response.ConsolidacaoResponse{}
	}

	return consolidacoesResponse, err
}

func (s *ConsolidacaoService) GetConsolidacaoByCodigo(codigo string) (response.ConsolidacaoResponse, error) {
	consolidacao, err := s.repo.GetConsolidacaoByCodigo(strings.ToUpper(codigo))

	if err != nil {
		return response.ConsolidacaoResponse{}, err
	}

	return mapper.ToConsolidacaoResponse(consolidacao), err
}

func (s *ConsolidacaoService) GenerateConsolidacoes() ([]response.ConsolidacaoResponse, error) {
	var consolidacoesResponse []response.ConsolidacaoResponse

	movimentacoesAcao, err := s.movimentacaoService.GetMovimentacoesGroupedByCodigo()
	if err != nil {
		return []response.ConsolidacaoResponse{}, err
	}

	for codigo, movimentacoes := range movimentacoesAcao {
		qtdEntrada := 0.0
		qtdSaida := 0.0
		valorEntrada := 0.0
		valorSaida := 0.0
		var tipoAtivo string
		for _, movimentacao := range movimentacoes {
			if movimentacao.Operacao == "ENTRADA" {
				qtdEntrada = qtdEntrada + movimentacao.Quantidade
				valorEntrada = valorEntrada + movimentacao.ValorTotal
			} else {
				qtdSaida = qtdSaida + movimentacao.Quantidade
				valorSaida = valorSaida + movimentacao.ValorTotal
			}
		}

		consolidacao := model.NewConsolidacao(codigo, tipoAtivo, qtdEntrada, valorEntrada, qtdSaida, valorSaida)
		consolidacaoExist, err := s.GetConsolidacaoByCodigo(codigo)

		var consolidacaoPersisted model.Consolidacao
		if err == nil {
			consolidacao.ID = consolidacaoExist.ID
			consolidacaoPersisted, err = s.repo.UpdateConsolidacao(*consolidacao)
		} else {
			consolidacaoPersisted, err = s.repo.CreateConsolidacao(*consolidacao)
		}

		if err != nil {
			return []response.ConsolidacaoResponse{}, err
		}

		consolidacoesResponse = append(consolidacoesResponse, mapper.ToConsolidacaoResponse(consolidacaoPersisted))

	}

	return consolidacoesResponse, err
}
