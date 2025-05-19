ALTER TABLE public.cotacao_historico ADD quantidade int;

ALTER TABLE public.cotacao_historico ADD valor_total numeric(8, 2);

ALTER TABLE public.cotacao_historico ADD variacao numeric(8, 5);
