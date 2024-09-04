ALTER TABLE ativo_rentabilidade ALTER COLUMN roe TYPE numeric(8, 4) USING roe::numeric(8, 4);
ALTER TABLE ativo_rentabilidade ALTER COLUMN roa TYPE numeric(8, 4) USING roa::numeric(8, 4);
