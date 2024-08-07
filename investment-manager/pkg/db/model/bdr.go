package model

type Bdr struct {
	ID     uint   `gorm:"column:bdr_id;primaryKey;autoIncrement"`
	Nome   string `gorm:"column:nome"`
	Codigo string `gorm:"column:codigo"`
	Setor  string `gorm:"column:setor"`
	Cnpj   string `gorm:"column:cnpj"`
}

func (Bdr) TableName() string {
	return "brazilian_depositary_receipts"
}

func NewBdr(nome string, codigo string, setor string, cnpj string) *Bdr {
	return &Bdr{
		Nome:   nome,
		Codigo: codigo,
		Setor:  setor,
		Cnpj:   cnpj,
	}
}
