package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToBdr(bdrRequest request.BdrRequest) *model.Bdr {
	return model.NewBdr(bdrRequest.Nome, bdrRequest.Codigo, bdrRequest.Setor, bdrRequest.Cnpj)
}

func ToBdrResponse(bdr model.Bdr) response.BdrResponse {
	return response.NewBdrResponse(bdr.ID, bdr.Nome, bdr.Codigo, bdr.Setor, bdr.Cnpj)
}

func UpdateBdr(bdr model.Bdr, bdrRequest request.BdrRequest) model.Bdr {
	bdr.Codigo = bdrRequest.Codigo
	bdr.Nome = bdrRequest.Nome
	bdr.Cnpj = bdrRequest.Cnpj
	bdr.Setor = bdrRequest.Setor
	return bdr
}
