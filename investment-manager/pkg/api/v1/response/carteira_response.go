package response

type CarteiraResponse struct {
	ID                uint
	Nome              string
	Descricao         string
	ProporcaoDesejada float64
	Ativos            []CarteiraAtivoResponse
}

func NewCarteiraResponse(id uint, nome string, descricao string, proporcaoDesejada float64, ativos []CarteiraAtivoResponse) CarteiraResponse {
	return CarteiraResponse{
		ID:                id,
		Nome:              nome,
		Descricao:         descricao,
		ProporcaoDesejada: proporcaoDesejada,
		Ativos:            ativos,
	}
}
