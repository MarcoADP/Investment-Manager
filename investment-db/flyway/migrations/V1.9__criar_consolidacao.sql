create table consolidacao (
    consolidacao_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    codigo varchar(8),
    tipo_ativo varchar(32),
    quantidade_entrada decimal(8, 2),
    valor_medio_entrada decimal(8, 2),
    valor_total_entrada decimal(8, 2),
    quantidade_saida decimal(8, 2) default 0,
    valor_medio_saida decimal(8, 2) default 0,
    valor_total_saida decimal(8, 2) default 0,
    lucro_medio decimal(8, 2) default 0
)