package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAtivoRentabilidade(ativoInformacao model.AtivoInformacao) model.AtivoRentabilidade {
	return *model.NewAtivoRentabilidade(ativoInformacao)
}

func ToAtivoRentabilidadeResponse(ativoRentabilidade model.AtivoRentabilidade) response.AtivoRentabilidadeResponse {
	return response.NewAtivoRentabilidadeResponse(ativoRentabilidade)
}
