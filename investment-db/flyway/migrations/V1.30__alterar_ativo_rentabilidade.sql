ALTER TABLE ativo_rentabilidade ADD ativo_informacao_id INT;

ALTER TABLE ativo_rentabilidade ADD CONSTRAINT fk_ativo_rentabilidade_ativo_informacao FOREIGN KEY (ativo_informacao_id) REFERENCES ativo_informacao(ativo_informacao_id)
ON DELETE CASCADE;