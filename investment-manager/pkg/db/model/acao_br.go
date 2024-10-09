package model

type AcaoBr struct {
	ID      uint   `gorm:"column:acao_id;primaryKey;autoIncrement"`
	Nome    string `gorm:"column:nome"`
	Codigo  string `gorm:"column:codigo"`
	Cnpj    string `gorm:"column:cnpj"`
	SetorId uint   `gorm:"column:setor_id"`
}

func (AcaoBr) TableName() string {
	return "acao_br"
}

func NewAcaoBr(nome string, codigo string, cnpj string, setorId uint) *AcaoBr {
	return &AcaoBr{
		Nome:    nome,
		Codigo:  codigo,
		Cnpj:    cnpj,
		SetorId: setorId,
	}
}
