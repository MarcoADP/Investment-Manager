ALTER TABLE ativo_eficiencia ALTER COLUMN margem_liquida TYPE numeric(8, 4) USING margem_liquida::numeric(8, 4);
ALTER TABLE ativo_eficiencia ALTER COLUMN margem_bruta TYPE numeric(8, 4) USING margem_bruta::numeric(8, 4);
ALTER TABLE ativo_eficiencia ALTER COLUMN margem_ebit TYPE numeric(8, 4) USING margem_ebit::numeric(8, 4);
ALTER TABLE ativo_eficiencia ALTER COLUMN margem_ebitda TYPE numeric(8, 4) USING margem_ebitda::numeric(8, 4);
