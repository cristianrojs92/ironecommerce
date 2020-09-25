--
-- Script para la creacion de las tablas iniciales del sistema
--

--
--USERS (usuarios)
--

CREATE TABLE IF NOT EXISTS users (

  id                            serial PRIMARY KEY,  
  username                      varchar(50) not null,
  name                          varchar(50) NOT NULL,
  surname                       varchar(50) NOT NULL,         
  email                         varchar(50),
  password                      varchar(500) NOT NULL,
  delete                        boolean NOT NULL default false,

  CONSTRAINT users_unique UNIQUE (username)
);
CREATE INDEX IF NOT EXISTS users_index_1 ON users USING btree(username, password);

--
-- AUTH (autenticaciones)
--

CREATE TABLE IF NOT EXISTS auth (
  user_id       integer REFERENCES users(id) ON DELETE CASCADE,
  token         varchar(3000) not null,
  
  CONSTRAINT auth_unique UNIQUE (user_id)
);

--
--CATEGORIES (categorias)
--

CREATE TABLE IF NOT EXISTS categories (

  id                            serial PRIMARY KEY,  
  main                          boolean NOT NULL default false,
  name                          varchar(100) NOT NULL,
  icon_name                     varchar(50) NOT NULL
);

--
--CATEGORIES_RELATIONS (relaciones entre categorias)
--

CREATE TABLE IF NOT EXISTS categories_relations (

  id_main                          integer NOT NULL,
  id_child                         integer NOT NULL,

  CONSTRAINT categories_relations_unique UNIQUE (id_main,id_child)
);

CREATE INDEX IF NOT EXISTS categories_relations_index_1 ON categories_relations USING btree(id_main, id_child);

--
--PRODUCTS (productos)
--

CREATE TABLE IF NOT EXISTS products (

   id                            serial PRIMARY KEY, 
   id_category                   integer NOT NULL,
   name                          varchar(100) NOT NULL,
   description                   varchar(1000) NOT NULL,  
   price                         integer NOT NULL,
   stock                         integer NOT NULL,
   image                         varchar(100) NOT NULL
);

CREATE INDEX IF NOT EXISTS products_index_1 ON products USING btree(id_category);

--
--OFFERS (ofertas)
--

CREATE TABLE IF NOT EXISTS offers (

   id_product                    integer PRIMARY KEY,
   discount                      integer NOT NULL,
   expiration                    varchar(10) NOT NULL,
   stock                         integer NOT NULL
);

