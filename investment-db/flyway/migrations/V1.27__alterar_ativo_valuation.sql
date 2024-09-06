ALTER TABLE ativo_valuation ADD ativo_informacao_id INT;

ALTER TABLE ativo_valuation ADD CONSTRAINT fk_ativo_valuation_ativo_informacao FOREIGN KEY (ativo_informacao_id) REFERENCES ativo_informacao(ativo_informacao_id)
ON DELETE CASCADE;