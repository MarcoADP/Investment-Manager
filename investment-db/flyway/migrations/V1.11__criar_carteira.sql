create table carteira (
    carteira_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    nome varchar(64),
    descricao varchar(128),
    valor_total_compra numeric(8,2),
    valor_total_atual numeric(8,2),
    proporcao_atual numeric(4,2),
    proporcao_desejada numeric(4,2),
    variacao numeric(8,5),
    data_atualizacao date not null default now()
)