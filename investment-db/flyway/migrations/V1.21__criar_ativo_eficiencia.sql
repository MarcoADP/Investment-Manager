create table ativo_eficiencia (
    ativo_eficiencia_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_calculo date not null default now(),
    codigo varchar(8),
    margem_bruta numeric(8,2),
    margem_liquida numeric(8,2),
    margem_ebit numeric(8,2),
    margem_ebitda numeric(8,2)
)