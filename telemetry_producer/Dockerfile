# Imagem base para compilar o código Go
FROM golang:latest AS build

# Diretório de trabalho
WORKDIR /app

# Copia arquivos do projeto para o contêiner
COPY . .

# Compila o código Go
RUN go build -o telemetry_producer cmd/main.go 

# Imagem base para executar o binário Go
FROM debian:latest

# Configura diretório de trabalho
WORKDIR /app

# Copia o binário compilado do estágio anterior
COPY --from=build /app/telemetry_producer .
COPY --from=build /app/config.toml .

# Comando padrão para executar o binário Go
CMD ["./telemetry_producer"]