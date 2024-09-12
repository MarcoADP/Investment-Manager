create table graham_formula (
    graham_formula_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    codigo varchar(8),
    data_calculo date not null default now(),
    preco_atual numeric(8,2),
    lpa numeric(8,2),
    vpa numeric(8,2),
    pl_esperado numeric(8,2),
    pvp_esperado numeric(8,2),
    preco_justo numeric(8,2),
    margem_seguranca numeric(8,2)
)