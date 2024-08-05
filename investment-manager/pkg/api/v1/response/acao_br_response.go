package response

type AcaoBrResponse struct {
	ID     uint
	Nome   string
	Codigo string
	Setor  string
	Cnpj   string
}

func NewAcaoBrResponse(id uint, nome string, codigo string, setor string, cnpj string) AcaoBrResponse {
	return AcaoBrResponse{
		ID:     id,
		Nome:   nome,
		Codigo: codigo,
		Setor:  setor,
		Cnpj:   cnpj,
	}
}
