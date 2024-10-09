package model

type Setor struct {
	ID   uint   `gorm:"column:setor_id;primaryKey;autoIncrement"`
	Nome string `gorm:"column:nome"`
}

func (Setor) TableName() string {
	return "setor"
}

func NewSetor(nome string) *Setor {
	return &Setor{
		Nome: nome,
	}
}
