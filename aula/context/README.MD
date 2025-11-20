# 🧠 Context em Go --- Cancelamento, Timeout e Propagação

O pacote `context` do Go fornece uma forma padronizada de controlar
**deadlines**, **cancelamentos** e **metadados** durante a execução de
operações --- especialmente em aplicações HTTP, serviços distribuídos e
chamadas externas.

Ele é amplamente usado para evitar operações presas, liberar recursos
corretamente e interromper tarefas quando o cliente desconecta.

------------------------------------------------------------------------

## 🔍 O que é um Context?

Um `context.Context` é um objeto imutável que carrega:

  -------------------------------------------------------------------------
  Recurso            Descrição
  ------------------ ------------------------------------------------------
  **Deadline**       Quando a operação deve expirar

  **Cancelamento**   Permite abortar operações encadeadas

  **Valores**        Dados que viajam junto com a requisição (usado com
                     cuidado)

  **Propagação**     Quando um contexto é cancelado, todos os filhos também
  -------------------------------------------------------------------------

👉 *O Context é sempre passado como primeiro parâmetro em funções longas
ou que fazem I/O.*

``` go
func Processar(ctx context.Context) error { ... }
```

------------------------------------------------------------------------

## ✨ Tipos de Context

Go fornece funções para criar Contexts com diferentes comportamentos:

  ----------------------------------------------------------------------------
  Função                     Para que serve
  -------------------------- -------------------------------------------------
  `context.Background()`     Contexto vazio/raiz (preferido em servidores)

  `context.TODO()`           Quando você ainda não sabe qual contexto usar

  `context.WithCancel()`     Permite cancelamento manual

  `context.WithTimeout()`    Cancela automaticamente após uma duração

  `context.WithDeadline()`   Cancela em um horário específico

  `context.WithValue()`      Adiciona valores (use com moderação)
  ----------------------------------------------------------------------------

------------------------------------------------------------------------

## 📘 Como Context funciona internamente

Sempre que um contexto é cancelado:

-   Todos os contextos-filho são cancelados automaticamente
-   Todas as goroutines ouvindo `<-ctx.Done()` são finalizadas
-   Erros são propagados pela função `ctx.Err()`

Isso evita goroutines zumbis e deadlocks.

------------------------------------------------------------------------

# 🌐 Exemplos HTTP na prática

## 🚀 1. Cancelando uma operação quando o cliente desconecta

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(5 * time.Second):
        fmt.Fprintln(w, "Processo finalizado com sucesso!")
    case <-ctx.Done():
        fmt.Println("❌ Cliente desconectou ou cancelou a requisição:", ctx.Err())
        return
    }
}
```

------------------------------------------------------------------------

## ⏳ 2. Timeout na requisição HTTP (server-side)

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()

    select {
    case <-time.After(3 * time.Second):
        fmt.Fprintln(w, "Finalizado!")
    case <-ctx.Done():
        http.Error(w, "Timeout atingido", http.StatusGatewayTimeout)
    }
}
```

------------------------------------------------------------------------

## 🌍 3. Requisição HTTP externa + Context

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Fprintf(w, "Erro: %v", err)
        return
    }

    fmt.Fprintln(w, "Status:", resp.StatusCode)
}
```

------------------------------------------------------------------------

## 💾 4. Passando valores com Context (mas com cautela!)
Não é muito recomendavel devido não saber o que possui dentro do contexto, a melhor forma é passar como parametro
É igual um map, usando key e value, nesse caso a key será "userID" e value "42"

``` go
type key string

func handler(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), key("userID"), 42)
    Processar(ctx)
}

func Processar(ctx context.Context) {
    fmt.Println("User ID:", ctx.Value(key("userID")))
}
```

------------------------------------------------------------------------

# 🔐 Boas Práticas com Context

✔ Use `context.Background()` na inicialização da aplicação\
✔ Sempre passe o context como **primeiro parâmetro**\
✔ Sempre chame `cancel()` para contextos com timeout/cancel\
✔ Nunca armazene contextos em variáveis globais\
✔ Evite usar `context.WithValue` para dados complexos

------------------------------------------------------------------------

# 📎 Recursos úteis

- [Documentação oficial](https://pkg.go.dev/context)
- [Effective Go](https://go.dev/doc/effective_go#contexts)
- [Go blog](https://go.dev/blog/context)
