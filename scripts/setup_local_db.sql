CREATE SCHEMA IF NOT EXISTS alura_store;
CREATE TABLE alura_store.products (
  id serial primary key,
  nome varchar,
  descricao varchar,
  preco decimal,
  quantidade integer
)