create table ativo_historico_preco (
    ativo_historico_preco_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    codigo varchar(8),
    valor numeric(8,2),
    data_preco date not null default now()
)