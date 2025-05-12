create table consolidacao_provento (
    consolidacao_provento_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    codigo varchar(8),
    tipo_ativo varchar(32),
    ano int,
    valor_total decimal(8, 2),
    valor_medio decimal(8, 5)
)