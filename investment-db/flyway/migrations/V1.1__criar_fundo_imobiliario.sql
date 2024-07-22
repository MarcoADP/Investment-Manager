create table FUNDO_IMOBILIARIO (
    fundo_imobiliario_id bigint primary key,
    nome varchar(128),
    codigo varchar(6) UNIQUE,
    cnpj varchar(16),
    setor varchar(64)
)