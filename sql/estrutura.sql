-- public.titular definition

-- Drop table

DROP TABLE if exists public.titular;
CREATE TABLE if exists public.titular (
	id serial NOT NULL,
	nome varchar(255) NULL,
	cpf varchar(11) NULL,
	email varchar(255) NULL,
	nascimento date NULL,
	nomepai varchar(255) NULL,
	nomemae varchar(255) NULL,
	cidade text NULL,
	estado text NULL,
	CONSTRAINT titular_cpf_key UNIQUE (cpf),
	CONSTRAINT titular_pkey PRIMARY KEY (id)
);

-- public.contabancaria definition

-- Drop table

DROP TABLE if exists public.contabancaria;
CREATE TABLE if not exists public.contabancaria (
	id serial NOT NULL,
	codigobanco int4 NULL,
	agencia varchar(10) NULL,
	conta varchar(10) NULL,
	digito varchar(10) NULL,
	titular_id int4 NULL,
	CONSTRAINT contabancaria_pkey PRIMARY KEY (id)
);

-- public.contabancaria foreign keys

ALTER TABLE public.contabancaria ADD CONSTRAINT contabancaria_codigobanco_fkey FOREIGN KEY (codigobanco) REFERENCES public.bancos(cod);

-- public.saldo definition

-- Drop table

DROP TABLE if exists public.saldo;
CREATE TABLE if not exists public.saldo (
	titular_id int4 NULL,
	conta_id int4 NULL,
	saldo numeric(10, 4) NULL,
	created timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated timestamp NULL
);


-- public.saldo foreign keys

ALTER TABLE public.saldo ADD CONSTRAINT saldo_conta_id_fkey FOREIGN KEY (conta_id) REFERENCES public.contabancaria(id);
ALTER TABLE public.saldo ADD CONSTRAINT saldo_titular_id_fkey FOREIGN KEY (titular_id) REFERENCES public.titular(id);

-- public.registrosaldo definition

-- Drop table

DROP TABLE if exists public.registrosaldo;
CREATE TABLE if not exists public.registrosaldo (
	id serial NOT NULL,
	titular_id int4 NULL,
	conta_id int4 NULL,
	sinal int4 NULL,
	valor numeric(10, 4) NULL,
	created timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT registrosaldo_pkey PRIMARY KEY (id)
);


-- public.registrosaldo foreign keys

ALTER TABLE public.registrosaldo ADD CONSTRAINT registrosaldo_conta_id_fkey FOREIGN KEY (conta_id) REFERENCES public.contabancaria(id);
ALTER TABLE public.registrosaldo ADD CONSTRAINT registrosaldo_titular_id_fkey FOREIGN KEY (titular_id) REFERENCES public.titular(id);

-- public.transferencia definition

-- Drop table

DROP TABLE if exists public.transferencia;
CREATE TABLE if not exists public.transferencia (
	contaorigem_id int4 NULL,
	contadestino_id int4 NULL,
	valor numeric(10, 4) NULL,
	created timestamp NULL DEFAULT CURRENT_TIMESTAMP
);


-- public.transferencia foreign keys

ALTER TABLE public.transferencia ADD CONSTRAINT transferencia_contadestino_id_fkey FOREIGN KEY (contadestino_id) REFERENCES public.contabancaria(id);
ALTER TABLE public.transferencia ADD CONSTRAINT transferencia_contaorigem_id_fkey FOREIGN KEY (contaorigem_id) REFERENCES public.contabancaria(id);