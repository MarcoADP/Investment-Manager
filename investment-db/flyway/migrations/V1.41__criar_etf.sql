create table ETF (
    etf_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    nome varchar(128),
    codigo varchar(5) UNIQUE,
    cnpj varchar(16),
    tipo varchar(64)
)