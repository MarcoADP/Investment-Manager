create table ativo_endividamento (
    ativo_endividamento_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_calculo date not null default now(),
    codigo varchar(8),
    divida_patrimonio_liquido numeric(8,2),
    divida_ebit numeric(8,2),
    divida_ebitda numeric(8,2)
)