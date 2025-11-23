# 📘 Guia: Estrutura de Projetos Go, Dependências e Workspaces

## 📂 Estrutura de diretórios

Em projetos Go organizados, é comum manter:

```
/cmd
    /app1
        main.go
    /app2
        main.go
/pkg ou /internal
    (bibliotecas, módulos reutilizáveis)
```

A pasta `cmd/` normalmente contém os binários da aplicação — ou seja, os arquivos que possuem a função `main()`.

Cada subpasta dentro de `cmd/` costuma representar um executável diferente.

---

## 🔗 Dependências entre pacotes

Quando um pacote importa outro dentro do mesmo projeto, chamamos isso de:

### ✔️ Dependência direta

```go
import "meuprojeto/math"
```

Se esse pacote importar outro, temos:

### ✔️ Dependência indireta

```
cmd/app1 → pkg/math → pkg/utils
```

O Go resolve dependências diretas e indiretas automaticamente via `go mod tidy`.

---

### Ignorar o import de dependencia quando não conseguir

```bash
go mod tidy -e
```

Funciona como o `go mod tidy`, mas **ignora erros de módulos faltando**.

É útil quando:

- Você está reorganizando diretórios.
- Criando módulos em etapas.
- Usando workspaces e alguns pacotes ainda não existem.

**Atenção:**  
Não é recomendado como solução permanente, pois oculta problemas reais.  
Para verificação correta, sempre finalize com:

```bash
go mod tidy
```
## ❌ Por que evitar `go mod replace` para dependências locais?

Antes dos Workspaces, era comum usar:

```go
replace github.com/user/projeto/math => ../math
```

Mas isso traz problemas:

- Obriga a editar o `go.mod` antes de commitar
- Pode quebrar builds em outros ambientes
- Depende de caminhos locais
- É considerado **má prática**

Use somente em casos muito específicos.

---

## 🚀 Go Workspaces: a forma moderna e recomendada

Desde o Go 1.18, o recomendado é trabalhar com **Workspaces** para múltiplos módulos locais:

### Criando um workspace

```bash
go work init ./cmd ./math
```

Isso cria o arquivo `go.work`, que informa ao Go que esses módulos fazem parte do mesmo workspace.

### Vantagens do Workspace

- Evita o uso de `replace`
- Nada precisa ser modificado no `go.mod`
- Módulos locais funcionam automaticamente
- Ideal para projetos grandes ou multi-módulo

---

## 💡 Exemplo de `go.work`

```go
go 1.22

use (
    ./cmd
    ./math
)
```

Agora, basta importar normalmente:

```go
import "github.com/user/projeto/math"
```

O Go automaticamente usará o diretório local incluído no workspace.

---

## ✔️ Resumo Geral

- `cmd/` → executáveis (`main.go`)
- `pkg/` ou `internal/` → bibliotecas internas
- Dependências indiretas funcionam automaticamente
- `go mod replace` funciona, mas **não é recomendado**
- **Go Workspaces** é a forma moderna:
  - `go work init`
  - facilita projetos multi-módulo
  - não polui o `go.mod`

---

Se quiser, posso gerar outros arquivos também! 🎉
