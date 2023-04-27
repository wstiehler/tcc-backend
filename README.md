| Branch |                                                                         
:------------------------------------------------------------------------------------------------------------------------------------------------------: | -------------------------------------------------------------------------------------------------------------------------------------------------------: |
| master | 

# alerts platform api



## Desenvolvimento

- docker (para dependências)
- docker-compose (para dependências)
- golang ([site oficial do golang](https://golang.org/doc/install) ou [instalando via asdf](https://github.com/kennyp/asdf-golang))

### Setup

Instale as dependências e faça as configurações nescessarioas para o pre-commit

```sh
make setup
```


### Rodando o serviço

```sh
make dev-start-with-tools
```


#### Adicione as sequintes variaveis de ambiente para rodar a aplicação. 

Crie o arquivo `.env.local` na raiz do projeto.

.env.local

|Nome|Tipo|
|----|----|
|APPLICATION_PORT| string|
|APPLICATION_ADDRESS| string|
|HOST| string|
|USER| number|
|PASSWORD|  string|
|DB_NAME |string|
|MYSQL_ROOT_PASSWORD |string|
|MYSQL_DATABASE|  string|
|CORS_URL| string|


#### Makefile

Criamos um `Makefile` para simplificar o desenvolvimento. Para rodar o serviço em ambiente local basta rodar o comando abaixo:

Inicia a infraestrutura local, e o serviço:

```sh
make dev-start
```

Destroi a infraestrutura local, e o para serviço:

```sh
make dev-destroy
```

Visualizar os logs

```sh
make dev-logs
```

Para visualizar os comandos disponíveis basta rodar o comando `make help`

### Rodando lint

Antes de rodar o comando lint, rode o comando abaixo para instalar o pacote [staticcheck](https://staticcheck.io/docs/install):

```sh
go install honnef.co/go/tools/cmd/staticcheck@latest
```

Para rodar o comando lint:

```sh
make lint
```


## Testes unitários e coverage

Os testes unitários estão configurados para serem executados na pipeline da aplicação.

Execução dos testes unitários localmente.

```sh
make test-unit
```

```sh
make test-e2e-local
```

Exibir o relatório de coverage da aplicação

```sh
make showcover
```

## Testar as operações de rest-api:

- Baixar o insomnia ou postman.
- Baixar o json e importar no insomnia/postman. Link abaixo. 


