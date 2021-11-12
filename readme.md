# Desafio Tecnico inBolso

## Proposta

#### 1 - Criar um módulo de gerenciamento de contas bancárias

##### Informações do titular
 - Nome
 - cpf
 - data de nascimento
 - nome do pai
 - nome da mãe
 - cidade
 - estado...

##### Informações da conta
 - Código do banco
 - Agência
 - Conta
 - Dígito

##### Informações da API
- API para criar conta bancária
- API para editar dados
- API para excluir conta

#### 2 - Criar um módulo de transferência bancária


##### Informações da API
- API para consultar saldo da conta
- API para fazer transferências de uma conta para outra

Obs: Seria legal adicionar aqui também uma opção de adicionar saldo a conta, para testes.

#### 3 - Criar um módulo de geração de boletos

##### Informações da API
- API para criar um boleto bancário (sem necessidade de criar o pdf, só os dados mesmo).
- API para verificar dados de um boleto já gerado

PS1:
Para a geração dos boletos, seria legal gerar uma linha digitável próxima do real, usando as mesmas regras. Vou enviar 2 links que podem ajudar a entender e testar como essa parte funciona:
https://gerencianet.com.br/blog/campos-dos-boletos-linha-digitavel/
http://www.sicadi.com.br/mhouse/boleto/geraboleto.php

## Requisitos
Será bom ter os seguintes itens instalado em sua máquina:

- Golang 1.17+
- Docker 19+
- Ambiente linux

## Instruções
Depois de realizar o download do projeto, dentro de cada pasta há um arquivo chamado `deploy`, o qual realiza o build do módulo referente e cria uma imagem docker
para realizar um deploy separado de cada aplicação.
Garanta que cada arquivo `deploy` tenha a permissão de executar, um `chmod +x deploy` já resolve.
Caso queira excluir todos as imagens e containers existe um arquivo chamado `undeploy` na raiz de, que exclui todas estas peculiaridades.

Exportei via `Postman`, todos os endpoints desenvolvidos para facilitar o teste da APIs, os mesmos contém uma documentação integrada à requisição para poder sanar sobre o que a mesma raliza no sistema.


### Melhorias pós análise
Itens para implementar para melhoria do projeto

- Validador de campos de struct (Utilizando o Validator do go-playground)
- Utilizar GORM para uma melhoria de tratamento de banco de dados
