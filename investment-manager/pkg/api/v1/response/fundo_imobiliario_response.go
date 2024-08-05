package response

type FundoImobiliarioResponse struct {
	ID       uint
	Nome     string
	Codigo   string
	Tipo     string
	Segmento string
	Cnpj     string
}

func NewFundoImobiliarioResponse(id uint, nome string, codigo string, tipo string, segmento string, cnpj string) FundoImobiliarioResponse {
	return FundoImobiliarioResponse{
		ID:       id,
		Nome:     nome,
		Codigo:   codigo,
		Tipo:     tipo,
		Segmento: segmento,
		Cnpj:     cnpj,
	}
}
