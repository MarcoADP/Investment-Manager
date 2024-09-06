ALTER TABLE ativo_endividamento ADD ativo_informacao_id INT;

ALTER TABLE ativo_endividamento ADD CONSTRAINT fk_ativo_endividamento_ativo_informacao FOREIGN KEY (ativo_informacao_id) REFERENCES ativo_informacao(ativo_informacao_id)
ON DELETE CASCADE;