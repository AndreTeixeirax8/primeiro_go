CREATE TABLE contato(
    id varchar(20) PRIMARY KEY,
    nome varchar(100) NOT NULL, 
    email varchar(100) ,
    unidade_id varchar(20) varchar not null 
    create_at TIMESTAMP,
    update_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (unidade_id) REFERENCES unidade(id) ON UPDATE CASCADE ON DELETE CASCADE
