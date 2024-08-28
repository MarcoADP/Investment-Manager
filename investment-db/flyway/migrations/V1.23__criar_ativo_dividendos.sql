create table ativo_dividendo (
    ativo_dividendo_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_calculo date not null default now(),
    codigo varchar(8),
    dividendos numeric(8,4),
    dy numeric(8,2),
    yoc numeric(8,2)
)