ALTER TABLE fundo_imobiliario RENAME COLUMN setor TO tipo;

ALTER TABLE fundo_imobiliario ADD segmento varchar(64) NOT NULL;
