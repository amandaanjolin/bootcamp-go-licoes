-- resolva as seguintes consultas:

-- Mostrar todos os registros da tabela de filmes.
select * from movies
-- Exibir o primeiro nome, o sobrenome e a classificação de todos os atores.
select first_name, last_name from actors
-- Exibir o título de todas as séries e usar aliases para que o nome da tabela e o campo estejam em inglês.
select title as Title from series
-- Exibir o primeiro nome e o sobrenome dos atores cuja classificação seja superior a 7,5.
select first_name, last_name from actors where rating > 7.5
-- Exibir o título do filme, a classificação e os prêmios dos filmes com classificação superior a 7,5 e com mais de dois prêmios.
select title, rating, awards from movies where rating > 7.5 and awards >2
-- Exibir o título do filme e a classificação ordenados por classificação em ordem crescente.
select title, rating from movies order by rating asc
--  Exibir os títulos dos filmes e a classificação ordenados por classificação.
select title, rating from movies order by rating 
-- Exibir os títulos dos três primeiros filmes no banco de dados.
select *  from movies  limit 3
-- Exibir os 5 filmes mais bem classificados.
select * from movies order by rating desc limit 5
-- Listar os 10 primeiros atores.
select * from actors  limit 10
-- Exibir o título e a classificação de todos os filmes cujo título é Toy Story.
select title, rating from movies  where title lIKE '%toy story%'
-- Exibir todos os atores cujos nomes começam com Sam.
select * from actors where first_name lIKE '%sam%'
-- Exibir o título dos filmes lançados entre 2004 e 2008.
select * from movies where year(release_date) >= 2004 and year(release_date) <= 2008
-- Exibir o título dos filmes com classificação superior a 3, com mais de 1 prêmio e com data de lançamento entre 1988 e 2009.
select title from movies where year(release_date) >= 1988 and year(release_date) <= 2009 and rating >3 and awards > 1
-- Ordenar os resultados por classificação.
select title, rating from movies where year(release_date) >= 1988 and year(release_date) <= 2009 and rating >3 and awards > 1 order by rating desc