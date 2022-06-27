-- Table: public.produtos

-- DROP TABLE IF EXISTS public.produtos;

CREATE TABLE IF NOT EXISTS produtos
(
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
)
