package response

type BdrResponse struct {
	ID     uint
	Nome   string
	Codigo string
	Setor  string
	Cnpj   string
}

func NewBdrResponse(id uint, nome string, codigo string, setor string, cnpj string) BdrResponse {
	return BdrResponse{
		ID:     id,
		Nome:   nome,
		Codigo: codigo,
		Setor:  setor,
		Cnpj:   cnpj,
	}
}
