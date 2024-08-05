package model

type AcaoBr struct {
	ID     uint   `gorm:"column:acao_id;primaryKey;autoIncrement"`
	Nome   string `gorm:"column:nome"`
	Codigo string `gorm:"column:codigo"`
	Setor  string `gorm:"column:setor"`
	Cnpj   string `gorm:"column:cnpj"`
}

func (AcaoBr) TableName() string {
	return "acao_br"
}

func NewAcaoBr(nome string, codigo string, setor string, cnpj string) *AcaoBr {
	return &AcaoBr{
		Nome:   nome,
		Codigo: codigo,
		Setor:  setor,
		Cnpj:   cnpj,
	}
}
