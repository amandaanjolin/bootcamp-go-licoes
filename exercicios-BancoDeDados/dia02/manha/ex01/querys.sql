-- São necessárias as seguintes consultas:

-- Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
SELECT f.nome, f.posto, d.localidad
FROM FUNCIONARIO f
JOIN DEPARTAMENTO d ON f.depto_nro = d.depto_nro
WHERE f.posto = 'Vendedor';

-- Mostra os departamentos com mais de cinco empregados.
SELECT d.nombre_depto, COUNT(*) AS total_empregados
FROM FUNCIONARIO f
JOIN DEPARTAMENTO d ON f.depto_nro = d.depto_nro
GROUP BY d.depto_nro, d.nombre_depto
HAVING COUNT(*) > 5;

-- Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".
SELECT f.nome, f.salário, d.nombre_depto
FROM FUNCIONARIO f
JOIN DEPARTAMENTO d ON f.depto_nro = d.depto_nro
WHERE f.posto = (
    SELECT posto FROM FUNCIONARIO
    WHERE nome = 'Mito' AND sobrenome = 'Barchuk'
);

-- Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.
SELECT f.*
FROM FUNCIONARIO f
JOIN DEPARTAMENTO d ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidade'
ORDER BY f.nome;

-- Mostra o nome do empregado com o salário mais baixo.
SELECT nome, sobrenome, salário
FROM FUNCIONARIO
ORDER BY salário ASC
LIMIT 1;

-- Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT f.*
FROM FUNCIONARIO f
JOIN DEPARTAMENTO d ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Vendas'
ORDER BY f.salário DESC
LIMIT 1;
