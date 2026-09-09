# Feel-The-Night-Swagger

## Sumário

- [Descrição Geral](#descrição-geral)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Estrutura do Repositório](#estrutura-do-repositório)
- [Funcionamento](#funcionamento)
- [Configuração](#configuração)
- [Execução](#execução)
- [Swagger / Documentação da API](#swagger--documentação-da-api)
- [Observações](#observações)

## Descrição Geral

Este projeto será o backend e a camada de banco de dados do projeto [Feel-The-Night](https://github.com/kisalto/Feel-The-Night), com foco em fornecer uma API REST para gestão de usuários, personagens, guias, eventos e outros conteúdos relacionados ao universo do jogo.

A API será desenvolvida em Go, com documentação Swagger gerada automaticamente, banco de dados PostgreSQL e acesso ao banco via GORM, mantendo uma estrutura organizada para evoluir com o frontend do projeto.

## Tecnologias Utilizadas

- Go: linguagem principal da API
- Gin: framework HTTP para construção da REST API
- Swaggo / Swag: geração automática de documentação Swagger
- PostgreSQL: banco de dados relacional
- GORM: ORM para acesso e manipulação de dados no PostgreSQL
- godotenv: carregamento de variáveis de ambiente
- Docker / Docker Compose: opcional para facilitar o ambiente local

## Estrutura do Repositório

A estrutura abaixo serve como referência para a organização do backend:

```text
Feel-The-Night-Swagger/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── database.go
│   │   └── migrations/
│   │       └── ...
│   ├── handlers/
│   │   └── ...
│   ├── models/
│   │   └── ...
│   ├── routes/
│   │   └── ...
│   ├── services/
│   │   └── ...
│   └── swagger/
│       └── docs.go
├── docs/
│   ├── swagger.json
│   └── swagger.yaml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── docker-compose.yml
├── LICENSE
├── README.md
└── Makefile
```

## Funcionamento

1. O servidor Go inicia em `cmd/api/main.go`.
2. A aplicação carrega as configurações de ambiente e se conecta ao PostgreSQL.
3. O GORM realiza o mapeamento dos modelos para o banco de dados.
4. As rotas HTTP são registradas e expõem endpoints REST para o frontend consumi-los.
5. A documentação Swagger é gerada a partir dos comentários das rotas e modelos.
6. As respostas são retornadas em JSON para o cliente.

## Configuração

### 1. Pré-requisitos

Antes de iniciar o projeto, verifique se os itens abaixo já estão instalados:

```bash
Go 1.22+
PostgreSQL em execução
Git
Docker (opcional, para ambiente local com containers)
```

### 2. Instalação das dependências

```bash
go mod download
```

Se ainda não houver o módulo do projeto, inicialize com:

```bash
go mod init github.com/seu-usuario/Feel-The-Night-Swagger
```

### 3. Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com base no exemplo abaixo:

```env
APP_PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=feel_the_night
DB_SSLMODE=disable
```

Também pode usar um arquivo `.env.example` para documentar as variáveis esperadas.

### 4. Banco de Dados

O projeto utiliza PostgreSQL como banco principal. Para configurar localmente:

```bash
createdb feel_the_night
```

Ou, se preferir usar Docker Compose:

```bash
docker-compose up -d postgres
```

### 5. Migrações

As migrações podem ser gerenciadas com GORM ou scripts SQL, dependendo da abordagem adotada no projeto.

Exemplo de uso com GORM:

```bash
go run cmd/api/main.go
```

A aplicação pode automaticamente criar as tabelas com `AutoMigrate`, ou você pode manter migrações manuais em `internal/database/migrations`.

## Execução

### Ambiente de desenvolvimento

```bash
go run ./cmd/api
```

### Build para produção

```bash
go build -o ./bin/api ./cmd/api
./bin/api
```

## Swagger / Documentação da API

A documentação da API será gerada com Swaggo, com base nos comentários dos endpoints e modelos.

### Instalar Swaggo

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Gerar documentação

```bash
swag init -g cmd/api/main.go -o docs
```

### Acessar a documentação

Após iniciar a API, a documentação fica disponível em:

```text
http://localhost:8080/swagger/index.html
```

Se estiver usando outra rota ou porta configurada, ajuste de acordo com o `.env`.

## Observações

- Este README foi estruturado com base no README original do projeto `Feel-The-Night`, adaptando a descrição para uma API REST em Go.
- O objetivo principal é manter a proposta do projeto original, mas trazendo uma arquitetura moderna com Swagger, PostgreSQL e GORM.
- O projeto pode ser expandido com módulos para usuários, personagens, eventos, guias e outras entidades do jogo.

## Sobre

Um backend para o site de conteúdo de Under Night In-Birth, com foco em API REST, banco de dados PostgreSQL e documentação Swagger.

## Licença

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais informações.
