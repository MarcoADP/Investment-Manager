package model

import (
	"time"
)

type AtivoInformacao struct {
	ID                uint      `gorm:"column:ativo_informacao_id;primaryKey;autoIncrement"`
	DataInformacao    time.Time `gorm:"column:data_informacao;type:date"`
	Codigo            string    `gorm:"column:codigo"`
	NumeroAcoes       uint64    `gorm:"column:numero_acoes"`
	ValorFirma        float64   `gorm:"column:valor_firma"`
	LucroLiquido      float64   `gorm:"column:lucro_liquido"`
	LucroBruto        float64   `gorm:"column:lucro_bruto"`
	ReceitaLiquida    float64   `gorm:"column:receita_liquida"`
	PatrimonioLiquido float64   `gorm:"column:patrimonio_liquido"`
	AtivoTotal        float64   `gorm:"column:ativo_total"`
	DividaLiquida     float64   `gorm:"column:divida_liquida"`
	Ebit              float64   `gorm:"column:ebit"`
	Ebitda            float64   `gorm:"column:ebitda"`
}

func (AtivoInformacao) TableName() string {
	return "ativo_informacao"
}

func NewAtivoInformacao(data time.Time, codigo string, numeroAcoes uint64, valorFirma float64,
	lucroLiquido float64, lucroBruto float64, receitaLiquida float64, patrimonioLiquido float64,
	ativoTotal float64, dividaLiquida float64, ebit float64, ebitda float64,
) *AtivoInformacao {
	return &AtivoInformacao{
		DataInformacao:    data,
		Codigo:            codigo,
		NumeroAcoes:       numeroAcoes,
		ValorFirma:        valorFirma,
		LucroLiquido:      lucroLiquido,
		LucroBruto:        lucroBruto,
		ReceitaLiquida:    receitaLiquida,
		PatrimonioLiquido: patrimonioLiquido,
		AtivoTotal:        ativoTotal,
		DividaLiquida:     dividaLiquida,
		Ebit:              ebit,
		Ebitda:            ebitda,
	}
}
