# 📘 Guia Rápido de Configuração e Compilação em Go

Este README fornece instruções e dicas importantes para trabalhar com projetos Go, incluindo gerenciamento de dependências, compilação, variáveis de ambiente e mais.

---

## 📦 Comandos Essenciais

### 🔁 Gerenciamento de Dependências

- `go mod init <repo_git>`  
  Cria o arquivo `go.mod`.  
  `<repo_git>` é o nome do repositório no Git ou URL para identificação do módulo (importação futura).

- `go mod tidy`  
  Atualiza e limpa as dependências do projeto com base no uso real do código.

- `go get <pacote>`  
  Instala o pacote especificado.

- `go get -u <pacote>`  
  Atualiza o pacote para a última versão disponível.

- `go get -u ./...`  
  Atualiza todos os pacotes do projeto.

- `go get -u <pacote>@<versão>`  
  Atualiza o pacote para uma versão específica.

- `go get -u <pacote>@latest`  
  Atualiza o pacote para a versão mais recente.

- `go list -m all`  
  Lista todos os módulos utilizados pelo projeto.

- `go mod graph`  
  Exibe o grafo de dependências dos módulos.

- `go mod verify`  
  Verifica se os módulos no cache estão íntegros.

### 📁 Diretórios e Cache

- `GOPATH`  
  Local onde os pacotes são instalados.  
  ⚠️ **Evite configurar em pastas que exijam permissão de administrador.**

- `GOMODCACHE`  
  Local onde o cache de módulos Go é armazenado.

- `GONOPROXY`  
  Define domínios que serão acessados diretamente, ignorando o proxy.

- `GOPROXY=https://proxy.golang.org,direct`  
  Define o proxy para baixar módulos Go.

- `GOBIN`  
  Define onde os binários compilados serão salvos ao usar `go install`.

- `GOENV`  
  Caminho para o arquivo de configurações do ambiente Go.

---

## ⚙️ Compilação

### Comandos Básicos

- `go run .`  
  Compila e executa todos os arquivos Go do pacote atual.

- `go build .`  
  Compila todos os arquivos do pacote atual e gera um executável.

- `go build main.go`  
  Compila um arquivo Go específico.

- `go install`  
  Compila e instala o pacote como um binário (salvo em `$GOBIN` ou `$GOPATH/bin`).

### Compilação Cruzada

Para compilar para diferentes sistemas operacionais e arquiteturas:

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe main.go
```

- `GOOS` – Define o sistema operacional de destino.  
- `GOARCH` – Define a arquitetura de destino.

🔍 Verifique as opções disponíveis com:

```bash
go tool dist list
```

📌 Verifique as configurações atuais com:

```bash
go env GOOS GOARCH
```

---

## 🧪 Testes

- `go test ./...`  
  Executa todos os testes dos pacotes do projeto.

- `go test -v`  
  Mostra saída detalhada dos testes.

- `go test -cover`  
  Exibe cobertura de testes.

---

## 🔧 Variáveis de Ambiente (GOENV)

Visualize todas as variáveis com:

```bash
go env
```

Variáveis úteis:

- `GOOS`, `GOARCH` – Sistema operacional e arquitetura alvo para compilação.
- `GOPATH` – Diretório de trabalho Go (pacotes, binários).
- `GOBIN` – Caminho onde binários são instalados.
- `GOROOT` – Diretório onde o Go está instalado.
- `GOMODCACHE` – Cache dos módulos.
- `GO111MODULE=on|off|auto` – Controle explícito sobre o uso de módulos Go.

---

## 🛠️ Outras Dicas Importantes

- Sempre use `go mod tidy` para manter o `go.mod` e `go.sum` organizados.
- Inclua `go.mod` e `go.sum` no controle de versão (Git).
- Utilize `go fmt ./...` para formatar seu código automaticamente.
- Use `go vet ./...` para identificar problemas potenciais no código.

---

## 📚 Referências

- [Go Modules](https://golang.org/ref/mod)
- [DigitalOcean: Cross Compilation](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)
- [Go Command Documentation](https://pkg.go.dev/cmd/go)

---