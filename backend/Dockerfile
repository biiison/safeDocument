# Etapa 1: Construção do aplicativo
FROM golang:1.23.1-alpine AS build

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar dependências do Go
RUN go mod tidy

# Copiar o código fonte para o contêiner
COPY . .

# Compilar o código Go
RUN go build -o app .

# Etapa 2: Imagem final
FROM alpine:latest

# Instalar dependências para rodar o binário (caso precise de algum)
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário da etapa anterior
COPY --from=build /app/app .

# Expor a porta que o aplicativo vai usar
EXPOSE 8080

# Comando para rodar o aplicativo
CMD ["./app"]
