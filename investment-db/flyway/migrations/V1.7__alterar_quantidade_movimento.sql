ALTER TABLE movimentacao ALTER COLUMN quantidade TYPE numeric(8, 2) USING quantidade::numeric(8, 2);
