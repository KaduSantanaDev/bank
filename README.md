# Bank API

Este projeto foi desenvolvido durante o curso **"Backend Master Class [Golang + Postgres + Kubernetes + gRPC]"** e visa criar uma API robusta para gerenciamento de banco de dados. A API é construída com a linguagem **Golang** e integra-se ao **PostgreSQL** para armazenamento de dados. A autenticação é realizada utilizando **JWT** e **Paseto**.

## Tecnologias Utilizadas

- **[Golang](https://golang.org/)**: Linguagem principal para o desenvolvimento da API.
- **[Gin](https://github.com/gin-gonic/gin)**: Framework para construção da API em Go.
- **[PostgreSQL](https://www.postgresql.org/)**: Banco de dados relacional.
- **[JWT](https://github.com/golang-jwt/jwt)** e **[Paseto](https://github.com/o1egl/paseto)**: Autenticação segura.
- **[Docker](https://www.docker.com/)**: Containerização da aplicação.
- **[AWS EC2](https://aws.amazon.com/ec2/)** e **[RDS](https://aws.amazon.com/rds/)**: Infraestrutura para deploy.
- **[GitHub Actions](https://github.com/features/actions)**: para automação do fluxo de CI/CD (build e deploy).

## Como Rodar o Projeto

### 1. Clonar o Repositório

```bash
git clone https://github.com/KaduSantanaDev/bank.git
cd bank
```

### 2. Configuração do Banco de Dados
crie um arquivo .env com as configurações do banco de dados:
```bash
AWS_DB_USER=your_db_user
AWS_DB_PASSWORD=your_db_password
AWS_DB_HOST=your_db_host
AWS_DB_PORT=5432
AWS_DB_NAME=your_db_name
LOCAL_DB_SOURCE=your_local_db_source
```

### 3. Rodando a Aplicação com Docker
```bash
docker-compose up --build
```

### 4. Executando Testes
Para rodar os testes, execute:
```bash
make test
```
ou
```bash
go test -v -cover ./...
```
### 5. Rodar as Migrations
```bash
make localmigrateup
```
ou
```bash
make migrate
```

## Deploy

O deploy é realizado com GitHub Actions, onde a imagem Docker é construída e enviada para o Amazon ECR. A infraestrutura de produção utiliza AWS EC2 e RDS para garantir escalabilidade e alta disponibilidade.

## Contribuição

Sinta-se à vontade para contribuir com melhorias ou correções. Faça um fork, crie uma branch, envie suas alterações e envie um pull request!
