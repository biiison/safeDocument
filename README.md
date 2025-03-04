# Safe Document

Este é o backend do projeto desenvolvido em Go, utilizando MongoDB como banco de dados e Docker para conteinerização.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação principal.
- **MongoDB**: Banco de dados NoSQL.
- **Docker**: Conteinerização do projeto.
- **Docker Compose**: Facilita a execução de containers de forma orquestrada.
- **Godotenv**: Para carregar variáveis de ambiente a partir de um arquivo `.env`.

## Pré-requisitos

Antes de rodar o projeto, certifique-se de que você possui os seguintes programas instalados:

- [Docker](https://docs.docker.com/get-docker/) (inclui Docker Compose)
- [Go 1.23.1 ou superior](https://golang.org/dl/)
- [MongoDB](https://www.mongodb.com/try/download/community) (se não usar o MongoDB no Docker)

## Configuração do Ambiente

1. **Clone o repositório**:

   Se ainda não tiver feito isso, clone o repositório para o seu diretório local:

   ```bash
   git clone <url-do-repositorio>
   cd <nome-do-repositorio>

2. **Crie o arquivo .env**:

Dentro da pasta do seu projeto, crie um arquivo .env com as variáveis de ambiente necessárias:

   ```bash
   MONGO_URI=mongodb://mongo:27017
   MONGO_DBNAME=safeDocument
```

3. **Inicie o Docker**:

No terminal, dentro da pasta do seu projeto, execute:

```bash
docker-compose up --build
```

Rodar o servidor localmente (sem Docker):

Se preferir rodar localmente (sem usar o Docker), siga os passos abaixo:

1. Instale as dependências (se ainda não fez isso):
```bash
go mod tidy
```
2. Execute o servidor:

```bash
go run main.go
```
## Frontend do Projeto

O frontend do projeto é desenvolvido em **Vue.js** e utiliza o **npm** para gerenciamento de dependências e execução do servidor de desenvolvimento.

### Pré-requisitos para o Frontend

Antes de rodar o frontend, certifique-se de que você tem os seguintes programas instalados:

- [Node.js e npm](https://nodejs.org/)
   - O npm será instalado automaticamente junto com o Node.js.

### Como rodar o Frontend
Antes de tudo entre na pasta frontend pelo seu terminal
```bash
cd ..
cd frontend
```

Instale as dependências:

No diretório do frontend do seu projeto, execute o seguinte comando para instalar todas as dependências necessárias:

```bash
npm install
```

Inicie o servidor de desenvolvimento:

Após as dependências serem instaladas, inicie o servidor de desenvolvimento com o comando:

```bash
npm run dev
```
