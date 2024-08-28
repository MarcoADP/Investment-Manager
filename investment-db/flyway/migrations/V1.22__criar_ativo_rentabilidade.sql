create table ativo_rentabilidade (
    ativo_rentabilidade_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_calculo date not null default now(),
    codigo varchar(8),
    roe numeric(8,2),
    roa numeric(8,2)
)