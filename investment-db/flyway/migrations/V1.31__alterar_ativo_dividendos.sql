ALTER TABLE ativo_dividendo ADD ativo_informacao_id INT;

ALTER TABLE ativo_dividendo ADD CONSTRAINT fk_ativo_dividendo_ativo_informacao FOREIGN KEY (ativo_informacao_id) REFERENCES ativo_informacao(ativo_informacao_id)
ON DELETE CASCADE;