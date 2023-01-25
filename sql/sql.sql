CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nome varchar(255) NOT NULL,
    nick varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    senha varchar(255) NOT NULL,
    Criadoem datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
