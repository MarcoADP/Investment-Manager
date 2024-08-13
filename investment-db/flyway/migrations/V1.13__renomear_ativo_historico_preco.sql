ALTER TABLE ativo_historico_preco RENAME TO cotacao_historico;

ALTER TABLE cotacao_historico RENAME COLUMN ativo_historico_preco_id TO cotacao_historico_id;
