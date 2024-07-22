create table ACAO_BR (
    acao_id bigint primary key,
    nome varchar(128),
    codigo varchar(5) UNIQUE,
    cnpj varchar(16),
    setor varchar(64)
)