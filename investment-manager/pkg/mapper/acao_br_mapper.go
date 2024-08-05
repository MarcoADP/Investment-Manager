package mapper

import (
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/request"
	"github.com/MarcoADP/Investment-Manager/pkg/api/v1/response"
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
)

func ToAcaoBr(acaoBrRequest request.AcaoBrRequest) *model.AcaoBr {
	return model.NewAcaoBr(acaoBrRequest.Nome, acaoBrRequest.Codigo, acaoBrRequest.Setor, acaoBrRequest.Cnpj)
}

func ToAcaoBrResponse(acaoBr model.AcaoBr) response.AcaoBrResponse {
	return response.NewAcaoBrResponse(acaoBr.ID, acaoBr.Nome, acaoBr.Codigo, acaoBr.Setor, acaoBr.Cnpj)
}

func UpdateAcaoBr(acaoBr model.AcaoBr, acaoBrRequest request.AcaoBrRequest) model.AcaoBr {
	acaoBr.Codigo = acaoBrRequest.Codigo
	acaoBr.Nome = acaoBrRequest.Nome
	acaoBr.Cnpj = acaoBrRequest.Cnpj
	acaoBr.Setor = acaoBrRequest.Setor
	return acaoBr
}
