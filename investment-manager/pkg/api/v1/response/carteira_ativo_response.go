package response

type CarteiraAtivoResponse struct {
	ID                uint
	Codigo            string
	ProporcaoDesejada float64
}

func NewCarteiraAtivoResponse(id uint, codigo string, proporcaoDesejada float64) CarteiraAtivoResponse {
	return CarteiraAtivoResponse{
		ID:                id,
		Codigo:            codigo,
		ProporcaoDesejada: proporcaoDesejada,
	}
}
