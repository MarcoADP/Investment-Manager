create table dividendo (
    dividendo_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    data_com date not null default now(),
    data_pagamento date not null default now(),
    tipo varchar(32),
    codigo varchar(8),
    valor numeric(12,8)
)