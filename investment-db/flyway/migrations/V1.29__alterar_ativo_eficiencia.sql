ALTER TABLE ativo_eficiencia ADD ativo_informacao_id INT;

ALTER TABLE ativo_eficiencia ADD CONSTRAINT fk_ativo_eficiencia_ativo_informacao FOREIGN KEY (ativo_informacao_id) REFERENCES ativo_informacao(ativo_informacao_id)
ON DELETE CASCADE;