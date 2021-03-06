CREATE DATABASE inbolso;

\c inbolso;

-- public.bancos definition

drop table if exists public.bancos;
CREATE TABLE if not exists public.bancos (
	cod serial NOT NULL, 
	banco varchar(255) NULL,
	CONSTRAINT cod_pkey PRIMARY KEY (cod)
);

INSERT INTO public.bancos (cod, banco) VALUES (001,'001 - BANCO DO BRASIL S/A');
INSERT INTO public.bancos (cod, banco) VALUES (002,'002 - BANCO CENTRAL DO BRASIL');
INSERT INTO public.bancos (cod, banco) VALUES (003,'003 - BANCO DA AMAZONIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (004,'004 - BANCO DO NORDESTE DO BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (007,'007 - BANCO NAC DESENV. ECO. SOCIAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (008,'008 - BANCO MERIDIONAL DO BRASIL');
INSERT INTO public.bancos (cod, banco) VALUES (020,'020 - BANCO DO ESTADO DE ALAGOAS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (021,'021 - BANCO DO ESTADO DO ESPIRITO SANTO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (022,'022 - BANCO DE CREDITO REAL DE MINAS GERAIS SA');
INSERT INTO public.bancos (cod, banco) VALUES (024,'024 - BANCO DO ESTADO DE PERNAMBUCO');
INSERT INTO public.bancos (cod, banco) VALUES (025,'025 - BANCO ALFA S/A');
INSERT INTO public.bancos (cod, banco) VALUES (026,'026 - BANCO DO ESTADO DO ACRE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (027,'027 - BANCO DO ESTADO DE SANTA CATARINA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (028,'028 - BANCO DO ESTADO DA BAHIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (029,'029 - BANCO DO ESTADO DO RIO DE JANEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (030,'030 - BANCO DO ESTADO DA PARAIBA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (031,'031 - BANCO DO ESTADO DE GOIAS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (032,'032 - BANCO DO ESTADO DO MATO GROSSO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (033,'033 - BANCO DO ESTADO DE SAO PAULO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (034,'034 - BANCO DO ESADO DO AMAZONAS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (035,'035 - BANCO DO ESTADO DO CEARA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (036,'036 - BANCO DO ESTADO DO MARANHAO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (037,'037 - BANCO DO ESTADO DO PARA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (038,'038 - BANCO DO ESTADO DO PARANA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (039,'039 - BANCO DO ESTADO DO PIAUI S.A');
INSERT INTO public.bancos (cod, banco) VALUES (041,'041 - BANCO DO ESTADO DO RIO GRANDE DO SUL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (047,'047 - BANCO DO ESTADO DE SERGIPE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (048,'048 - BANCO DO ESTADO DE MINAS GERAIS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (059,'059 - BANCO DO ESTADO DE RONDONIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (066,'066 - BANCO MORGAN STANLEY S.A');
INSERT INTO public.bancos (cod, banco) VALUES (070,'070 - BANCO DE BRASILIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (077,'077 - BANCO DE INTER S.A');
INSERT INTO public.bancos (cod, banco) VALUES (104,'104 - CAIXA ECONOMICA FEDERAL');
INSERT INTO public.bancos (cod, banco) VALUES (106,'106 - BANCO ITABANCO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (107,'107 - BANCO BBM S.A');
INSERT INTO public.bancos (cod, banco) VALUES (109,'109 - BANCO CREDIBANCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (116,'116 - BANCO B.N.L DO BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (129,'129 - UBS BRASIL BANCO DE INVESTIMENTO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (148,'148 - MULTI BANCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (151,'151 - CAIXA ECONOMICA DO ESTADO DE SAO PAULO');
INSERT INTO public.bancos (cod, banco) VALUES (153,'153 - CAIXA ECONOMICA DO ESTADO DO R.G.SUL');
INSERT INTO public.bancos (cod, banco) VALUES (165,'165 - BANCO NORCHEM S.A');
INSERT INTO public.bancos (cod, banco) VALUES (166,'166 - BANCO INTER-ATLANTICO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (168,'168 - BANCO C.C.F. BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (175,'175 - CONTINENTAL BANCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (184,'184 - BBA - CREDITANSTALT S.A');
INSERT INTO public.bancos (cod, banco) VALUES (199,'199 - BANCO FINANCIAL PORTUGUES');
INSERT INTO public.bancos (cod, banco) VALUES (200,'200 - BANCO FRICRISA AXELRUD S.A');
INSERT INTO public.bancos (cod, banco) VALUES (201,'201 - BANCO AUGUSTA INDUSTRIA E COMERCIAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (204,'204 - BANCO S.R.L S.A');
INSERT INTO public.bancos (cod, banco) VALUES (205,'205 - BANCO SUL AMERICA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (206,'206 - BANCO MARTINELLI S.A');
INSERT INTO public.bancos (cod, banco) VALUES (208,'208 - BANCO PACTUAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (210,'210 - DEUTSCH SUDAMERIKANICHE BANK AG');
INSERT INTO public.bancos (cod, banco) VALUES (211,'211 - BANCO SISTEMA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (212,'212 - BANCO ORIGINAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (213,'213 - BANCO ARBI S.A');
INSERT INTO public.bancos (cod, banco) VALUES (214,'214 - BANCO DIBENS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (215,'215 - BANCO AMERICA DO SUL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (216,'216 - BANCO REGIONAL MALCON S.A');
INSERT INTO public.bancos (cod, banco) VALUES (217,'217 - BANCO AGROINVEST S.A');
INSERT INTO public.bancos (cod, banco) VALUES (218,'218 - BBS - BANCO BONSUCESSO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (219,'219 - BANCO DE CREDITO DE SAO PAULO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (220,'220 - BANCO CREFISUL');
INSERT INTO public.bancos (cod, banco) VALUES (221,'221 - BANCO GRAPHUS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (222,'222 - BANCO AGF BRASIL S. A.');
INSERT INTO public.bancos (cod, banco) VALUES (223,'223 - BANCO INTERUNION S.A');
INSERT INTO public.bancos (cod, banco) VALUES (224,'224 - BANCO FIBRA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (225,'225 - BANCO BRASCAN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (228,'228 - BANCO ICATU S.A');
INSERT INTO public.bancos (cod, banco) VALUES (229,'229 - BANCO CRUZEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (230,'230 - BANCO BANDEIRANTES S.A');
INSERT INTO public.bancos (cod, banco) VALUES (231,'231 - BANCO BOAVISTA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (232,'232 - BANCO INTERPART S.A');
INSERT INTO public.bancos (cod, banco) VALUES (233,'233 - BANCO MAPPIN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (234,'234 - BANCO LAVRA S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (235,'235 - BANCO LIBERAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (236,'236 - BANCO CAMBIAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (237,'237 - BANCO BRADESCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (239,'239 - BANCO BANCRED S.A');
INSERT INTO public.bancos (cod, banco) VALUES (240,'240 - BANCO DE CREDITO REAL DE MINAS GERAIS S.');
INSERT INTO public.bancos (cod, banco) VALUES (241,'241 - BANCO CLASSICO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (242,'242 - BANCO EUROINVEST S.A');
INSERT INTO public.bancos (cod, banco) VALUES (243,'243 - BANCO STOCK S.A');
INSERT INTO public.bancos (cod, banco) VALUES (244,'244 - BANCO CIDADE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (245,'245 - BANCO EMPRESARIAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (246,'246 - BANCO ABC ROMA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (247,'247 - BANCO OMEGA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (249,'249 - BANCO INVESTCRED S.A');
INSERT INTO public.bancos (cod, banco) VALUES (250,'250 - BANCO SCHAHIN CURY S.A');
INSERT INTO public.bancos (cod, banco) VALUES (251,'251 - BANCO SAO JORGE S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (252,'252 - BANCO FININVEST S.A');
INSERT INTO public.bancos (cod, banco) VALUES (254,'254 - BANCO PARANA BANCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (255,'255 - MILBANCO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (256,'256 - BANCO GULVINVEST S.A');
INSERT INTO public.bancos (cod, banco) VALUES (258,'258 - BANCO INDUSCRED S.A');
INSERT INTO public.bancos (cod, banco) VALUES (260,'260 - NU PAGAMENTOS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (261,'261 - BANCO VARIG S.A');
INSERT INTO public.bancos (cod, banco) VALUES (262,'262 - BANCO BOREAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (263,'263 - BANCO CACIQUE');
INSERT INTO public.bancos (cod, banco) VALUES (264,'264 - BANCO PERFORMANCE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (265,'265 - BANCO FATOR S.A');
INSERT INTO public.bancos (cod, banco) VALUES (266,'266 - BANCO CEDULA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (267,'267 - BANCO BBM-COM.C.IMOB.CFI S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (275,'275 - BANCO REAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (277,'277 - BANCO PLANIBANC S.A');
INSERT INTO public.bancos (cod, banco) VALUES (282,'282 - BANCO BRASILEIRO COMERCIAL');
INSERT INTO public.bancos (cod, banco) VALUES (291,'291 - BANCO DE CREDITO NACIONAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (294,'294 - BCR - BANCO DE CREDITO REAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (295,'295 - BANCO CREDIPLAN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (300,'300 - BANCO DE LA NACION ARGENTINA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (302,'302 - BANCO DO PROGRESSO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (303,'303 - BANCO HNF S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (304,'304 - BANCO PONTUAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (308,'308 - BANCO COMERCIAL BANCESA S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (318,'318 - BANCO B.M.G. S.A');
INSERT INTO public.bancos (cod, banco) VALUES (320,'320 - BANCO INDUSTRIAL E COMERCIAL');
INSERT INTO public.bancos (cod, banco) VALUES (341,'341 - BANCO ITAU S.A');
INSERT INTO public.bancos (cod, banco) VALUES (346,'346 - BANCO FRANCES E BRASILEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (347,'347 - BANCO SUDAMERIS BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (351,'351 - BANCO BOZANO SIMONSEN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (353,'353 - BANCO GERAL DO COMERCIO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (356,'356 - ABN AMRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (366,'366 - BANCO SOGERAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (369,'369 - PONTUAL');
INSERT INTO public.bancos (cod, banco) VALUES (370,'370 - BEAL - BANCO EUROPEU PARA AMERICA LATINA');
INSERT INTO public.bancos (cod, banco) VALUES (372,'372 - BANCO ITAMARATI S.A');
INSERT INTO public.bancos (cod, banco) VALUES (375,'375 - BANCO FENICIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (376,'376 - CHASE MANHATTAN BANK S.A');
INSERT INTO public.bancos (cod, banco) VALUES (388,'388 - BANCO MERCANTIL DE DESCONTOS S/A');
INSERT INTO public.bancos (cod, banco) VALUES (389,'389 - BANCO MERCANTIL DO BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (392,'392 - BANCO MERCANTIL DE SAO PAULO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (394,'394 - BANCO B.M.C. S.A');
INSERT INTO public.bancos (cod, banco) VALUES (399,'399 - HSBC BANK BRASIL S.A. ??? BANCO M??LTIPLO');
INSERT INTO public.bancos (cod, banco) VALUES (409,'409 - UNIBANCO - UNIAO DOS public.bancos BRASILEIROS');
INSERT INTO public.bancos (cod, banco) VALUES (412,'412 - BANCO NACIONAL DA BAHIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (415,'415 - BANCO NACIONAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (420,'420 - BANCO NACIONAL DO NORTE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (422,'422 - BANCO SAFRA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (424,'424 - BANCO NOROESTE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (434,'434 - BANCO FORTALEZA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (453,'453 - BANCO RURAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (456,'456 - BANCO TOKIO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (464,'464 - BANCO SUMITOMO BRASILEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (466,'466 - BANCO MITSUBISHI BRASILEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (472,'472 - LLOYDS BANK PLC');
INSERT INTO public.bancos (cod, banco) VALUES (473,'473 - BANCO FINANCIAL PORTUGUES S.A');
INSERT INTO public.bancos (cod, banco) VALUES (477,'477 - CITIBANK N.A');
INSERT INTO public.bancos (cod, banco) VALUES (479,'479 - BANCO DE BOSTON S.A');
INSERT INTO public.bancos (cod, banco) VALUES (480,'480 - BANCO PORTUGUES DO ATLANTICO-BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (483,'483 - BANCO AGRIMISA S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (487,'487 - DEUTSCHE BANK S.A - BANCO ALEMAO');
INSERT INTO public.bancos (cod, banco) VALUES (488,'488 - BANCO J. P. MORGAN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (489,'489 - BANESTO BANCO URUGAUAY S.A');
INSERT INTO public.bancos (cod, banco) VALUES (492,'492 - INTERNATIONALE NEDERLANDEN BANK N.V.');
INSERT INTO public.bancos (cod, banco) VALUES (493,'493 - BANCO UNION S.A.C.A');
INSERT INTO public.bancos (cod, banco) VALUES (494,'494 - BANCO LA REP. ORIENTAL DEL URUGUAY');
INSERT INTO public.bancos (cod, banco) VALUES (495,'495 - BANCO LA PROVINCIA DE BUENOS AIRES');
INSERT INTO public.bancos (cod, banco) VALUES (496,'496 - BANCO EXTERIOR DE ESPANA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (498,'498 - CENTRO HISPANO BANCO');
INSERT INTO public.bancos (cod, banco) VALUES (499,'499 - BANCO IOCHPE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (501,'501 - BANCO BRASILEIRO IRAQUIANO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (502,'502 - BANCO SANTANDER S.A');
INSERT INTO public.bancos (cod, banco) VALUES (504,'504 - BANCO MULTIPLIC S.A');
INSERT INTO public.bancos (cod, banco) VALUES (505,'505 - BANCO GARANTIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (600,'600 - BANCO LUSO BRASILEIRO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (601,'601 - BFC BANCO S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (602,'602 - BANCO PATENTE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (604,'604 - BANCO INDUSTRIAL DO BRASIL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (607,'607 - BANCO SANTOS NEVES S.A');
INSERT INTO public.bancos (cod, banco) VALUES (608,'608 - BANCO OPEN S.A');
INSERT INTO public.bancos (cod, banco) VALUES (610,'610 - BANCO V.R. S.A');
INSERT INTO public.bancos (cod, banco) VALUES (611,'611 - BANCO PAULISTA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (612,'612 - BANCO GUANABARA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (613,'613 - BANCO PECUNIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (616,'616 - BANCO INTERPACIFICO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (617,'617 - BANCO INVESTOR S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (618,'618 - BANCO TENDENCIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (621,'621 - BANCO APLICAP S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (622,'622 - BANCO DRACMA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (623,'623 - BANCO PANAMERICANO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (624,'624 - BANCO GENERAL MOTORS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (625,'625 - BANCO ARAUCARIA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (626,'626 - BANCO FICSA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (627,'627 - BANCO DESTAK S.A');
INSERT INTO public.bancos (cod, banco) VALUES (628,'628 - BANCO CRITERIUM S.A');
INSERT INTO public.bancos (cod, banco) VALUES (629,'629 - BANCORP BANCO COML. E. DE INVESTMENTO');
INSERT INTO public.bancos (cod, banco) VALUES (630,'630 - BANCO INTERCAP S.A');
INSERT INTO public.bancos (cod, banco) VALUES (633,'633 - BANCO REDIMENTO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (634,'634 - BANCO TRIANGULO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (635,'635 - BANCO DO ESTADO DO AMAPA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (637,'637 - BANCO SOFISA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (638,'638 - BANCO PROSPER S.A');
INSERT INTO public.bancos (cod, banco) VALUES (639,'639 - BIG S.A. - BANCO IRMAOS GUIMARAES');
INSERT INTO public.bancos (cod, banco) VALUES (640,'640 - BANCO DE CREDITO METROPOLITANO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (641,'641 - BANCO EXCEL ECONOMICO S/A');
INSERT INTO public.bancos (cod, banco) VALUES (643,'643 - BANCO SEGMENTO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (645,'645 - BANCO DO ESTADO DE RORAIMA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (647,'647 - BANCO MARKA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (648,'648 - BANCO ATLANTIS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (649,'649 - BANCO DIMENSAO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (650,'650 - BANCO PEBB S.A');
INSERT INTO public.bancos (cod, banco) VALUES (652,'652 - ITA?? UNIBANCO HOLDING S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (653,'653 - BANCO INDUSVAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (654,'654 - BANCO A. J. RENNER S.A');
INSERT INTO public.bancos (cod, banco) VALUES (655,'655 - BANCO VOTORANTIM S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (656,'656 - BANCO MATRIX S.A');
INSERT INTO public.bancos (cod, banco) VALUES (657,'657 - BANCO TECNICORP S.A');
INSERT INTO public.bancos (cod, banco) VALUES (658,'658 - BANCO PORTO REAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (702,'702 - BANCO SANTOS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (705,'705 - BANCO INVESTCORP S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (707,'707 - BANCO DAYCOVAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (711,'711 - BANCO VETOR S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (713,'713 - BANCO CINDAM S.A');
INSERT INTO public.bancos (cod, banco) VALUES (715,'715 - BANCO VEGA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (718,'718 - BANCO OPERADOR S.A');
INSERT INTO public.bancos (cod, banco) VALUES (719,'719 - BANCO PRIMUS S.A');
INSERT INTO public.bancos (cod, banco) VALUES (720,'720 - BANCO MAXINVEST S.A');
INSERT INTO public.bancos (cod, banco) VALUES (721,'721 - BANCO CREDIBEL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (722,'722 - BANCO INTERIOR DE SAO PAULO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (724,'724 - BANCO PORTO SEGURO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (725,'725 - BANCO FINABANCO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (726,'726 - BANCO UNIVERSAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (728,'728 - BANCO FITAL S.A');
INSERT INTO public.bancos (cod, banco) VALUES (729,'729 - BANCO FONTE S.A');
INSERT INTO public.bancos (cod, banco) VALUES (730,'730 - BANCO COMERCIAL PARAGUAYO S.A');
INSERT INTO public.bancos (cod, banco) VALUES (731,'731 - BANCO GNPP S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (732,'732 - BANCO PREMIER S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (733,'733 - BANCO NACOES S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (734,'734 - BANCO GERDAU S.A');
INSERT INTO public.bancos (cod, banco) VALUES (735,'735 - BACO POTENCIAL');
INSERT INTO public.bancos (cod, banco) VALUES (736,'736 - BANCO UNITED S.A');
INSERT INTO public.bancos (cod, banco) VALUES (737,'737 - THECA');
INSERT INTO public.bancos (cod, banco) VALUES (738,'738 - MARADA');
INSERT INTO public.bancos (cod, banco) VALUES (739,'739 - BGN');
INSERT INTO public.bancos (cod, banco) VALUES (740,'740 - BCN BARCLAYS');
INSERT INTO public.bancos (cod, banco) VALUES (741,'741 - BRP');
INSERT INTO public.bancos (cod, banco) VALUES (742,'742 - EQUATORIAL');
INSERT INTO public.bancos (cod, banco) VALUES (743,'743 - BANCO EMBLEMA S.A');
INSERT INTO public.bancos (cod, banco) VALUES (744,'744 - THE FIRST NATIONAL BANK OF BOSTON');
INSERT INTO public.bancos (cod, banco) VALUES (745,'745 - CITIBAN N.A.');
INSERT INTO public.bancos (cod, banco) VALUES (746,'746 - MODAL S\A');
INSERT INTO public.bancos (cod, banco) VALUES (747,'747 - RAIBOBANK DO BRASIL');
INSERT INTO public.bancos (cod, banco) VALUES (748,'748 - SICREDI');
INSERT INTO public.bancos (cod, banco) VALUES (749,'749 - BRMSANTIL SA');
INSERT INTO public.bancos (cod, banco) VALUES (750,'750 - BANCO REPUBLIC NATIONAL OF NEW YORK (BRA');
INSERT INTO public.bancos (cod, banco) VALUES (751,'751 - DRESDNER BANK LATEINAMERIKA-BRASIL S/A');
INSERT INTO public.bancos (cod, banco) VALUES (752,'752 - BANCO BANQUE NATIONALE DE PARIS BRASIL S');
INSERT INTO public.bancos (cod, banco) VALUES (753,'753 - BANCO COMERCIAL URUGUAI S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (755,'755 - BANCO MERRILL LYNCH S.A');
INSERT INTO public.bancos (cod, banco) VALUES (756,'756 - BANCO COOPERATIVO DO BRASIL S.A.');
INSERT INTO public.bancos (cod, banco) VALUES (757,'757 - BANCO KEB DO BRASIL S.A.');


-- public.titular definition

-- Drop table

DROP TABLE if exists public.titular;
CREATE TABLE if not exists public.titular (
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


-- public.boleto definition

-- Drop table

-- DROP TABLE public.boleto;

CREATE TABLE public.boleto (
	"int" serial NOT NULL,
	banco int NULL,
	localpagamento text NULL,
	cedente text NULL,
	datadocumento date NULL,
	numerodocumento text NULL,
	especie text NULL,
	aceite text NULL,
	dataprocessamento date NULL,
	usobanco text NULL,
	carteira text NULL,
	especiemoeda text NULL,
	quantidade numeric(10, 4) NULL,
	valor numeric(10, 4) NULL,
	vencimento date NULL,
	agencia text NULL,
	conta text NULL,
	digito text NULL,
	seunumero text NULL,
	valordocumento numeric(10, 4) NULL,
	instrucao text NULL,
	mensagem1 text NULL,
	mensagem2 text NULL,
	mensagem3 text NULL,
	sacado text NULL,
	CONSTRAINT boleto_pkey PRIMARY KEY ("int")
);


-- public.boleto foreign keys

ALTER TABLE public.boleto ADD CONSTRAINT boleto_banco_fkey FOREIGN KEY (banco) REFERENCES public.bancos(cod);