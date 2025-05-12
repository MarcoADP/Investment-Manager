create table provento (
    provento_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_pagamento date not null default now(),
    tipo varchar(32),
    codigo varchar(8),
    quantidade int,
    valor_total decimal(8, 2),
    valor_medio decimal(8, 5)
)