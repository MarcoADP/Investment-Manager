ALTER TABLE acao_br ALTER COLUMN codigo TYPE varchar(8) USING codigo::varchar(8);

ALTER TABLE fundo_imobiliario ALTER COLUMN codigo TYPE varchar(8) USING codigo::varchar(8);

ALTER TABLE BRAZILIAN_DEPOSITARY_RECEIPTS ALTER COLUMN codigo TYPE varchar(8) USING codigo::varchar(8);

ALTER TABLE movimentacao ALTER COLUMN codigo TYPE varchar(8) USING codigo::varchar(8);
