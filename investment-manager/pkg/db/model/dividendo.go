package model

import (
	"time"
)

type Dividendo struct {
	ID            uint      `gorm:"column:dividendo_id;primaryKey;autoIncrement"`
	DataCom       time.Time `gorm:"column:data_com;type:date"`
	DataPagamento time.Time `gorm:"column:data_pagamento;type:date"`
	Tipo          string    `gorm:"column:tipo"`
	Codigo        string    `gorm:"column:codigo"`
	Valor         float64   `gorm:"column:valor"`
}

func (Dividendo) TableName() string {
	return "dividendo"
}

func NewDividendo(dataCom time.Time, dataPagamento time.Time, tipo string, codigo string, valor float64) *Dividendo {
	return &Dividendo{
		DataCom:       dataCom,
		DataPagamento: dataPagamento,
		Codigo:        codigo,
		Tipo:          tipo,
		Valor:         valor,
	}
}
