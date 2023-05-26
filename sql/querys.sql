CREATE DATABASE "Wallet"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Spanish_Spain.1252'
    LC_CTYPE = 'Spanish_Spain.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    dni integer NOT NULL,
    countryId varchar(100) NOT NULL,
    created date NOT NULL
);

CREATE TABLE logs (
	id SERIAL PRIMARY KEY,
	dni integer NOT NULL,
	petition_date date NOT NULL,
	status integer NOT NULL
);

ALTER TABLE logs
ADD COLUMN wallet_id integer

ALTER TABLE wallet
ADD COLUMN balance integer
