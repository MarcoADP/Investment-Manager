package model

import (
	"time"
)

type Movimentacao struct {
	ID            uint      `gorm:"column:movimentacao_id;primaryKey;autoIncrement"`
	Data          time.Time `gorm:"column:data;type:date"`
	Operacao      string    `gorm:"column:operacao"`
	Codigo        string    `gorm:"column:codigo"`
	TipoAtivo     string    `gorm:"column:tipo_ativo"`
	Quantidade    float64   `gorm:"column:quantidade"`
	ValorUnitario float64   `gorm:"column:valor_unitario"`
	ValorTotal    float64   `gorm:"column:valor_total"`
}

func (Movimentacao) TableName() string {
	return "movimentacao"
}

func NewMovimentacao(data time.Time, operacao string, codigo string, tipoAtivo string, quantidade float64, valorUnitario float64) *Movimentacao {
	return &Movimentacao{
		Data:          data,
		Operacao:      operacao,
		Codigo:        codigo,
		TipoAtivo:     tipoAtivo,
		Quantidade:    quantidade,
		ValorUnitario: valorUnitario,
		ValorTotal:    valorUnitario * float64(quantidade),
	}
}
