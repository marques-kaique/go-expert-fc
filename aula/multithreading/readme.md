# Concorrência, Paralelismo e Threads 
A computação moderna depende fortemente de **concorrência**,
**paralelismo**, **threads** e **schedulers**. Este documento explica
como o Go resolve problemas clássicos dessas áreas usando goroutines, um
runtime próprio e mecanismos leves de sincronização.

------------------------------------------------------------------------

# 🧵 Problemas com Threads Tradicionais

Threads compartilham memória, o que causa:

## 🏃‍♂️ Race Condition

Quando duas threads alteram uma variável ao mesmo tempo, gerando
resultados inesperados.

Documentação:\
- Go Memory Model: https://go.dev/ref/mem\
- Race Detector: https://go.dev/doc/articles/race_detector

------------------------------------------------------------------------

# 🔒 Mutex (Mutual Exclusion)

Um **mutex** impede que duas threads acessem simultaneamente uma região
crítica.

Funções: - `Lock()` --- bloqueia - `Unlock()` --- libera

Erro comum: esquecer o `Unlock()`.\
Solução idiomática no Go:

``` go
mu.Lock()
defer mu.Unlock()
```

Documentação: https://pkg.go.dev/sync#Mutex

------------------------------------------------------------------------

# ⚡ Concorrência vs Paralelismo

## Concorrência

Tarefas progredindo simultaneamente, dividindo o processador.

## Paralelismo

Tarefas executadas *ao mesmo tempo* em múltiplos cores.

Rob Pike: \> "Go é sobre concorrência. Paralelismo é um bônus quando há
múltiplos cores."

------------------------------------------------------------------------

# 💾 Custos das Threads Tradicionais

Em linguagens comuns: - Uma thread = **\~1 MB** - Custo alto de
criação - Troca de contexto pesada

Ineficiente para tarefas pequenas.

------------------------------------------------------------------------

# 🔄 Troca de Contexto (Context Switch)

Toda troca de thread envolve: - salvar registradores - salvar stack -
carregar estado da nova thread

Custa tempo e memória.

------------------------------------------------------------------------

# ⏱️ Schedulers

## Preemptivo

Interrompe a thread após um tempo pré-definido.

✔ Evita travamentos\
✘ Aumenta troca de contexto

## Cooperativo

Cada thread decide quando ceder CPU.

✔ Menos troca de contexto\
✘ Uma thread longa trava tudo

------------------------------------------------------------------------

# 🌱 Green Threads no Go

Goroutines **não são threads do SO**.\
Elas são "light threads" gerenciadas pelo próprio runtime.

Características: - custo inicial: **\~2 KB** - podem ser centenas de
milhares - criadas sem custo de SO - multiplexadas sobre poucas threads
reais

Documentação:\
https://go.dev/doc/faq#goroutines

------------------------------------------------------------------------

# ⚙ Scheduler do Go (Modelo M:N)

Componentes: - **G** = goroutine\
- **M** = thread do SO\
- **P** = processador lógico do runtime

Funcionamento: - G é executada por um M - P decide quais goroutines
podem rodar

O Go alterna entre: - modo cooperativo - preempção (desde Go 1.14)

Documentação: https://go.dev/src/runtime/proc.go

------------------------------------------------------------------------

# 📚 Fontes Oficiais

-   Effective Go -- Concurrency:
    https://go.dev/doc/effective_go#concurrency\
-   Go Concurrency Patterns (Rob Pike): https://go.dev/blog/pipelines\
-   Goroutines FAQ: https://go.dev/doc/faq#goroutines\
-   Memory Model: https://go.dev/ref/mem

------------------------------------------------------------------------

# ✔ Conclusão

Go facilita concorrência através de: - goroutines leves - scheduler
próprio - preempção inteligente - channels e mutexes - baixo custo de
memória

Isso torna o Go ideal para sistemas concorrentes e de alta performance.
