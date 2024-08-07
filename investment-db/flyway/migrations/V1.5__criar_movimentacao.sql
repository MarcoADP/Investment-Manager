create table movimentacao (
    movimentacao_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data date,
    operacao varchar(16),
    codigo varchar(5),
    tipo_ativo varchar(16),
    quantidade int,
    valor_unitario decimal(8, 2),
    valor_total decimal(8, 2)
)