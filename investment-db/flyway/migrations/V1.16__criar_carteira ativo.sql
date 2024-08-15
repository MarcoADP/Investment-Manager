create table carteira_ativo (
    carteira_ativo_id bigint GENERATED ALWAYS AS IDENTITY primary key,
    carteira_id bigint,
    codigo varchar(8),
    proporcao_desejada numeric(4,2),
    CONSTRAINT fk_carteira_id
        FOREIGN KEY (carteira_id) 
        REFERENCES carteira (carteira_id)
)