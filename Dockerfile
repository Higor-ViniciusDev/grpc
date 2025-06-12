# Usa a imagem golang:1.24.1-alpine como base
FROM golang:1.24.1-alpine

# Define o diretório de trabalho
WORKDIR /app

# Instala dependências necessárias: curl, unzip, protobuf e sudo
RUN apk add --no-cache curl unzip protobuf sudo

# Instala os plugins Go para Protocol Buffers
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Cria o usuário 'higor' com permissões sudo
RUN adduser -D higor && \
    echo "higor ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/higor && \
    chmod 0440 /etc/sudoers.d/higor

# Define o usuário padrão como 'higor'
USER higor

# Configura o PATH para incluir os binários Go instalados
ENV PATH=$PATH:/go/bin:/usr/local/bin
ENV GOPATH=/go

# Mantém o container ativo com um shell interativo
CMD ["/bin/sh", "-c", "while true; do sleep 3600; done"]