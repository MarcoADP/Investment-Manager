package model

import (
	"time"
)

type CotacaoHistorico struct {
	ID     uint      `gorm:"column:cotacao_historico_id;primaryKey;autoIncrement"`
	Data   time.Time `gorm:"column:data_preco;type:date"`
	Codigo string    `gorm:"column:codigo"`
	Valor  float64   `gorm:"column:valor"`
}

func (CotacaoHistorico) TableName() string {
	return "cotacao_historico"
}

func NewCotacaoHistorico(data time.Time, codigo string, valor float64) *CotacaoHistorico {
	return &CotacaoHistorico{
		Data:   data,
		Codigo: codigo,
		Valor:  valor,
	}
}
