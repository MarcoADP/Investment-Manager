create table BRAZILIAN_DEPOSITARY_RECEIPTS (
    bdr_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    nome varchar(128),
    codigo varchar(5) UNIQUE,
    cnpj varchar(16),
    setor varchar(64)
)