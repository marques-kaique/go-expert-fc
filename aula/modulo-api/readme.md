# Projeto Go

Este projeto segue o padrão recomendado pelo
**golang-standards/project-layout**, garantindo organização,
escalabilidade e facilidade de manutenção.

------------------------------------------------------------------------

## 📁 Estrutura de Pastas

### **/api**

Responsável por documentação e especificações da API.\
Aqui podem estar arquivos OpenAPI/Swagger, exemplos de request/response
e guias de uso.

### **/internal**

Contém toda **regra de negócio** e código que **não deve ser importado
por outros projetos**.\
Tudo que roda a aplicação vive aqui: services, repositories, handlers,
use-cases etc.

### **/pkg**

Bibliotecas públicas que podem ser reutilizadas em outros projetos via
`go mod`.\
Exemplos: - `pkg/auth` -- módulo de autenticação reutilizável -
`pkg/logger` -- abstração de logs - `pkg/validator` -- validações comuns

### **/cmd**

Onde ficam os **executáveis da aplicação**.\
O mais comum é ter:

    /cmd/server/main.go

Aqui é feito o bootstrap da aplicação.\
Alguns projetos simplificam e deixam apenas `/cmd/main.go`.

### **/configs**

Armazena arquivos de configuração do projeto:\
- `.env`\
- `config.yaml`\
- templates e exemplos de configuração

### **/test**

Contém arquivos adicionais para testes:\
- testes E2E\
- documentação de cenários\
- scripts auxiliares\
Nem tudo aqui é necessariamente `.go`.

------------------------------------------------------------------------

## 🗃️ Acesso ao Banco de Dados

### **SQLite**

    sqlite3 cmd/server/test.db

### **Ambiente (ENV)**

    DB_DRIVER=mysql
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=root
    DB_NAME=goexpert
    WEB_SERVER_PORT=8000
    JWT_SECRET=your-secret-key-here
    JWT_EXPIRESIN=300   # 300s = 5 minutos

------------------------------------------------------------------------

## 🔐 Token JWT --- Estrutura e Funcionamento

O JWT é dividido em **3 partes**, separadas por ponto:

### 1️⃣ **Header**

Antes do primeiro ponto.\
Define: - algoritmo de assinatura (ex: HS256, RS256) - tipo do token
(`JWT`)

### 2️⃣ **Payload**

Dados transmitidos após o primeiro ponto.\
Contém *claims*, como: - `sub` → normalmente armazena o `user_id` -
`name` - `exp` - `iat` - claims customizados

### 3️⃣ **Signature**

Após o segundo ponto.\
Responsável por validar: - autenticidade do token - integridade (garante
que não foi alterado)

Tokens podem ser assinados com chave secreta (HMAC) ou chave
pública/privada (RSA).

### 🔄 Token Expirado

Se o token estiver **válido mas expirado**, utiliza-se um **refresh
token** para gerar um novo access token.

------------------------------------------------------------------------

## 📦 Gerenciamento de Dependências

### **`go get`**

Baixa dependências e atualiza o `go.mod`.

### **`go install`**

Baixa e instala binários para uso local.\
Os binários vão para:

    $GOPATH/bin

-----------------------------------------------------------------------
