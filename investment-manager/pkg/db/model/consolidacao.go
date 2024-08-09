package model

type Consolidacao struct {
	ID                uint    `gorm:"column:consolidacao_id;primaryKey;autoIncrement"`
	Codigo            string  `gorm:"column:codigo"`
	TipoAtivo         string  `gorm:"column:tipo_ativo"`
	QuantidadeEntrada float64 `gorm:"column:quantidade_entrada"`
	ValorMedioEntrada float64 `gorm:"column:valor_medio_entrada"`
	ValorTotalEntrada float64 `gorm:"column:valor_total_entrada"`
	QuantidadeSaida   float64 `gorm:"column:quantidade_saida"`
	ValorMedioSaida   float64 `gorm:"column:valor_medio_saida"`
	ValorTotalSaida   float64 `gorm:"column:valor_total_saida"`
	LucroMedio        float64 `gorm:"column:lucro_medio"`
	LucroProporcao    float64 `gorm:"column:lucro_proporcao"`
}

func (Consolidacao) TableName() string {
	return "consolidacao"
}

func NewConsolidacao(codigo string, tipoAtivo string,
	quantidadeEntrada float64, valorTotalEntrada float64,
	quantidadeSaida float64, valorTotalSaida float64,
) *Consolidacao {
	lucroMedio := 0.0
	lucroProporcao := 0.0
	valorMedioEntrada := valorTotalEntrada / quantidadeEntrada
	valorMedioSaida := 0.0
	if quantidadeSaida > 0 {
		valorMedioSaida = valorTotalSaida / quantidadeSaida
		lucroMedio = valorMedioSaida - valorMedioEntrada
		lucroProporcao = (valorMedioSaida / valorMedioEntrada) - 1
	}
	return &Consolidacao{
		Codigo:            codigo,
		TipoAtivo:         tipoAtivo,
		QuantidadeEntrada: quantidadeEntrada,
		ValorMedioEntrada: valorMedioEntrada,
		ValorTotalEntrada: valorTotalEntrada,
		QuantidadeSaida:   quantidadeSaida,
		ValorMedioSaida:   valorMedioSaida,
		ValorTotalSaida:   valorTotalSaida,
		LucroMedio:        lucroMedio,
		LucroProporcao:    lucroProporcao,
	}
}
