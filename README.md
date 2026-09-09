# Feel-The-Night-Swagger

## Sumário

- [Descrição Geral](#descrição-geral)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Estrutura do Repositório](#estrutura-do-repositório)
- [Pré-requisitos](#pré-requisitos)
- [Instalação de Dependências](#instalação-de-dependências)
- [Configuração da Env](#configuração-da-env)
- [Banco de Dados](#banco-de-dados)
- [Migrations](#migrations)
- [Observações](#observações)
- [Sobre](#sobre)
- [Licença](#licença)

## Descrição Geral

Este projeto é o backend da API REST do projeto [Feel-The-Night](https://github.com/kisalto/Feel-The-Night), responsável por expor endpoints para gerenciamento de usuários, personagens, guias, eventos e demais conteúdos do universo do jogo.

A aplicação foi desenvolvida em Go, utilizando Gin para a API HTTP, GORM para acesso ao PostgreSQL e Swaggo para geração automática da documentação Swagger. O projeto também já está preparado para carregar variáveis de ambiente com `godotenv` e para rodar em modo de desenvolvimento com `air`.

## Tecnologias Utilizadas

- Go
- Gin
- GORM
- PostgreSQL
- Swaggo / Swag
- godotenv
- Air
- Docker / Docker Compose (opcional)

## Estrutura do Repositório

```text
Feel-The-Night-Swagger/
├── cmd/
│   └── api/
│       └── main.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── config/
│   ├── database/
│   │   └── database.go
│   ├── dto/
│   │   └── user_dto.go
│   ├── handler/
│   │   ├── routes.go
│   │   └── user_handler.go
│   ├── models/
│   │   └── models.go
│   ├── repository/
│   └── services/
│       └── user_service.go
├── .air.toml
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── tmp/
```

## Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- Go (versão compatível com o projeto; o arquivo `go.mod` está configurado com `go 1.27.1`)
- PostgreSQL em execução
- Git
- Opcional: Docker e Docker Compose para ambiente local com containers

## Instalação de Dependências

### 1. Instalar Go

Se ainda não tiver o Go instalado, siga a instalação oficial para o seu sistema operacional:

- https://go.dev/dl/

### 2. Instalar dependências do projeto

Na raiz do projeto, execute:

```bash
go mod download
```

Se quiser limpar e baixar novamente as dependências:

```bash
go mod tidy
```

### 3. Instalar PostgreSQL

O projeto usa PostgreSQL como banco principal. Caso ainda não tenha o banco configurado, instale e inicie o PostgreSQL localmente.

Exemplo de criação do banco local:

```bash
createdb feel_the_night
```

### 4. Instalar o Air

O `air` é útil para recarregar automaticamente a API durante o desenvolvimento.

```bash
go install github.com/air-verse/air@latest
```

Depois, para rodar o projeto em modo watch:

```bash
air
```

### 5. Instalar o Swag

O `swag` é usado para gerar a documentação Swagger a partir dos comentários da API.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Para regenerar os arquivos da documentação:

```bash
swag init -d ./cmd/api,./internal/handler,./internal/models,./internal/dto -g main.go
```

## Configuração da Env

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

Você também pode usar o arquivo `.env.example` como referência.

## Banco de Dados

O projeto utiliza PostgreSQL com GORM. A conexão é feita pela aplicação ao iniciar o servidor, conforme a configuração do `.env`.

Quando a API inicia, o projeto tenta executar o `AutoMigrate` automaticamente através da conexão do banco.

## Migrations

A aplicação já faz o mapeamento automático dos modelos com GORM. Em outras palavras, ao iniciar o servidor, as tabelas definidas nos models podem ser criadas ou atualizadas conforme necessário.

### Rodar a aplicação

```bash
go run ./cmd/api
```

### Build da aplicação

```bash
go build -o ./bin/api ./cmd/api
```

### Rodar com Air

```bash
air
```

### Acessar a documentação Swagger

Após iniciar a API, a documentação fica disponível em:

```text
http://localhost:8080/swagger/index.html
```

Se a porta ou rota forem alteradas no `.env`, ajuste de acordo com a configuração atual.

## Observações

- O projeto já possui os principais componentes da API estruturalmente organizados em `cmd`, `internal`, `docs` e `tmp`.
- O banco de dados é gerenciado com GORM e o schema dos modelos fica em `internal/models/models.go`.
- A documentação Swagger é gerada automaticamente a partir dos comentários de rotas e modelos.
- O `air` é recomendado para desenvolvimento local, pois recarrega a aplicação ao salvar alterações.
- O projeto pode ser expandido com novos módulos e endpoints conforme evolui o backend e o frontend do jogo.