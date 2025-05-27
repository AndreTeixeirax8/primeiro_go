CREATE TABLE unidade(
    id varchar(20) PRIMARY KEY,
    nome varchar(100) NOT NULL, 
    cnpj varchar(20) ,
    email varchar(100) ,
    qtd_silos int NOT NULL DEFAULT 0)