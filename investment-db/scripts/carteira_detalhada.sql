WITH cotacao_recente AS (
    SELECT 
        codigo,
        valor,
        data_preco,
        ROW_NUMBER() OVER (PARTITION BY codigo ORDER BY data_preco DESC) AS rn
    FROM 
        cotacao_historico
),
carteira_preenchida as (
	select c.carteira_id, c.nome, coalesce(ab.codigo, coalesce(fi.codigo, bdr.codigo)) as codigo, 
	case when fi.codigo is not null then fi.tipo when ab.codigo is not null then 'Ação' else 'BDR' end as tipo, coalesce(s.nome, coalesce(fi.segmento, bdr.setor)) as setor, 
	co.quantidade_entrada - co.quantidade_saida as quantidade, co.valor_medio_entrada as valor_compra, 
	(co.quantidade_entrada - co.quantidade_saida) * co.valor_medio_entrada as total_compra,
	ch.valor as valor_atual, (co.quantidade_entrada - co.quantidade_saida) * ch.valor as total_atual,
	(co.quantidade_entrada - co.quantidade_saida) * (ch.valor - co.valor_medio_entrada) as saldo,
	(ch.valor / co.valor_medio_entrada - 1) * 100 as variacao,
	ca.proporcao_desejada, ca.movimento, ch.data_preco 
from carteira_ativo ca
left join carteira c on c.carteira_id = ca.carteira_id 
left join acao_br ab on ab.codigo  = ca.codigo
left join fundo_imobiliario fi on fi.codigo  = ca.codigo 
left join brazilian_depositary_receipts bdr on bdr.codigo  = ca.codigo 
left join setor s on s.setor_id  = ab.setor_id
left join consolidacao co on co.codigo  = ca.codigo 
left join cotacao_recente ch on ch.codigo = ca.codigo and ch.rn = 1
where c.carteira_id = 2
),
carteira_total AS (
    SELECT 
    	cp.carteira_id,
        SUM(cp.total_compra) AS total_compra,
        SUM(cp.total_atual) as total_atual
    FROM 
        carteira_preenchida cp
    group by (cp.carteira_id)
)
select c.carteira_id, c.nome, c.codigo, c.tipo, c.setor, 
	c.quantidade, c.valor_compra, c.total_compra,
	c.valor_atual, c.total_atual, c.saldo, 
	ROUND(c.variacao, 2) as variacao,
	c.proporcao_desejada, 
	ROUND((c.total_compra / ct.total_compra)*100, 2) as proporcao_compra,
	ROUND((c.total_atual / ct.total_atual)*100, 2) as proporcao_atual,	
	c.movimento,
	c.data_preco
from carteira_preenchida c
left join carteira_total ct on ct.carteira_id = c.carteira_id
order by tipo, setor;