create table ativo_valuation (
    ativo_valuation_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_calculo date not null default now(),
    codigo varchar(8),
    lpa numeric(8,2),
    p_l numeric(8,2),
    vpa numeric(8,2),
    p_vp numeric(8,2),
    ev_ebit numeric(8,2),
    p_ebit numeric(8,2),
    ev_ebitda numeric(8,2),
    p_ebitda numeric(8,2)
)