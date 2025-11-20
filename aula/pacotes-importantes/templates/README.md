# 🧩 Templates em Go — `text/template` vs `html/template`

Go possui dois pacotes poderosos para geração de conteúdo dinâmico por meio de templates: `text/template` e `html/template`. Ambos permitem embutir dados em arquivos estruturados com marcação (como HTML, texto, etc), mas há **diferenças importantes de segurança** e uso.

---

## ✨ Diferença entre `text/template` e `html/template`

| Característica   | `text/template`                  | `html/template`                                      |
| ---------------- | -------------------------------- | ---------------------------------------------------- |
| Foco             | Texto genérico (emails, configs) | HTML/JS seguros para páginas web                     |
| Escapamento HTML | ❌ Não escapa                    | ✅ Escapa automaticamente para prevenir XSS          |
| Segurança Web    | ❌ Vulnerável a XSS              | ✅ Protege contra injeção maliciosa (ex: `<script>`) |
| Usado para       | Arquivos de config, markdowns    | Páginas HTML com dados dinâmicos                     |

👉 **Recomendado usar `html/template` sempre que for renderizar HTML** no navegador.

---

## 🔧 Sintaxe básica do template Go

Blocos de template usam `{{ }}`.

```gotemplate
{{ .Nome }}          // Acessa um campo "Nome"
{{ if .Ativo }} ... {{ end }}   // Condicional
{{ range .Cursos }} ... {{ end }} // Loop
{{ template "nome" . }} // Inclui subtemplate
```

---

## 📚 Funções disponíveis em templates

Você pode usar funções internas para comparar, formatar, etc:

| Função   | Descrição                  | Exemplo                         |
| -------- | -------------------------- | ------------------------------- |
| `eq`     | Igual a                    | `{{ if eq .Preco 29.90 }}`      |
| `ne`     | Diferente de               | `{{ if ne .Status "inativo" }}` |
| `lt`     | Menor que (`less than`)    | `{{ if lt .Preco 100 }}`        |
| `gt`     | Maior que (`greater than`) | `{{ if gt .Preco 200 }}`        |
| `printf` | Formata strings            | `{{ printf "%.2f" .Preco }}`    |

---

## 💡 Exemplo prático com `html/template`

### Go Struct:

```go
type Curso struct {
    Nome  string
    Preco float64
}
```

### Template HTML (`template.html`):

```html
{{range .}} {{ if eq .Preco 29.90 }}
<div class="highlight-yellow">{{ .Nome }} - R$ {{ printf "%.2f" .Preco }}</div>
{{ else if lt .Preco 100 }}
<div class="highlight-green">{{ .Nome }} - R$ {{ printf "%.2f" .Preco }}</div>
{{ else }}
<div class="highlight-blue">{{ .Nome }} - R$ {{ printf "%.2f" .Preco }}</div>
{{ end }} {{end}}
```

### Código Go para renderizar:

```go
tmpl := template.Must(template.ParseFiles("template.html"))
tmpl.Execute(w, cursos)
```

> 🔐 **Dica de segurança:** Use `html/template` para garantir que entradas do usuário sejam automaticamente escapadas. Isso evita ataques XSS (Cross-Site Scripting).

---

## 🔐 Sobre segurança com `html/template`

Quando você usa `html/template`, qualquer conteúdo passado para o template que poderia ser perigoso — como um `<script>` — será automaticamente escapado, protegendo seu site contra ataques.

```go
html/template --> <script>alert("XSS")</script>
renderiza como --> &lt;script&gt;alert("XSS")&lt;/script&gt;
```

✅ Você ainda pode renderizar HTML manualmente, se confiar na fonte, usando `template.HTML(...)`, mas **isso deve ser evitado sempre que possível**.

---

## 📎 Recursos úteis

- [Documentação oficial text/template](https://pkg.go.dev/text/template)
- [Documentação oficial html/template](https://pkg.go.dev/html/template)
- [Documentação oficial pkg-functions](https://pkg.go.dev/html/template#pkg-functions)
- [Guia de Templates Go (Go.dev)](https://go.dev/doc/articles/wiki/)
