create table ativo_informacao (
    ativo_informacao_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_informacao date not null default now(),
    codigo varchar(8),
    numero_acoes numeric(16,0),
    valor_firma numeric(16,2),
    lucro_liquido numeric(16,2),
    lucro_bruto numeric(16,2),
    receita_liquida numeric(16,2),
    patrimonio_liquido numeric(16,2),
    ativo_total numeric(16,2),
    divida_liquida numeric(16,2),
    ebit numeric(16,2),
    ebitda numeric(16,2)
)