# GraphQL com gqlgen em Go

## 📋 Índice

- [Sobre GraphQL](#sobre-graphql)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Configuração](#instalação-e-configuração)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Conceitos Importantes do Schema](#conceitos-importantes-do-schema)
- [Configuração do Banco de Dados](#configuração-do-banco-de-dados)
- [Trabalhando com Models Customizados](#trabalhando-com-models-customizados)
- [Implementando Relacionamentos](#implementando-relacionamentos)
- [Queries e Mutations](#queries-e-mutations)
- [Executando o Servidor](#executando-o-servidor)

## Sobre GraphQL

GraphQL é uma linguagem de consulta e manipulação de dados para APIs, assim como REST e gRPC. Embora não seja tão amplamente utilizado quanto REST, GraphQL pode ser muito útil, especialmente em arquiteturas BFF (Backend for Frontend), onde oferece:

- **Flexibilidade**: O cliente solicita exatamente os dados que precisa
- **Eficiência**: Reduz over-fetching e under-fetching de dados
- **Tipagem Forte**: Schema fortemente tipado com validação automática
- **Versionamento**: Evolução da API sem necessidade de versionamento
- **Documentação Automática**: O schema serve como documentação

## Pré-requisitos

- Go 1.16 ou superior
- SQLite3
- Conhecimento básico de Go e GraphQL

## Instalação e Configuração

### 1. Inicializar o Projeto

Crie o arquivo `tools.go` na raiz do projeto para gerenciar as dependências de ferramentas:

```go
//go:build tools
// +build tools

package tools

import (
    _ "github.com/99designs/gqlgen"
)
```

### 2. Inicializar o gqlgen

Execute o comando para criar toda a estrutura inicial do GraphQL:

```bash
go run github.com/99designs/gqlgen init
```

**Nota:** O comando correto é `go run github.com/99designs/gqlgen init`, não `go tool gqlgen init`.

Este comando cria:

- `gqlgen.yml` - arquivo de configuração
- `graph/schema.graphqls` - definição do schema GraphQL
- `graph/generated.go` - código gerado automaticamente
- `graph/resolver.go` - resolvers raiz
- `graph/schema.resolvers.go` - implementação dos resolvers
- `cmd/server/server.go` - servidor HTTP

### 3. Documentação Oficial

Para mais detalhes, consulte o tutorial completo em [gqlgen.com](https://gqlgen.com).

## Estrutura do Projeto

```
.
├── cmd/
│   └── server/
│       └── server.go           # Ponto de entrada da aplicação
├── graph/
│   ├── generated.go            # Código gerado (não editar)
│   ├── resolver.go             # Injeção de dependências
│   ├── schema.graphqls         # Definição do schema GraphQL
│   ├── schema.resolvers.go     # Implementação dos resolvers
│   └── model/
│       ├── category.go         # Model customizado de Category
│       ├── course.go           # Model customizado de Course
│       └── models_gen.go       # Models gerados automaticamente
├── internal/
│   └── database/
│       ├── category.go         # Repository de Category
│       └── course.go           # Repository de Course
├── gqlgen.yml                  # Configuração do gqlgen
├── go.mod                      # Dependências do Go
└── tools.go                    # Ferramentas de desenvolvimento
```

## Conceitos Importantes do Schema

No arquivo `schema.graphqls`, alguns conceitos são fundamentais:

### Input

Define dados de entrada para mutations:

```graphql
input NewCategory {
  name: String!
  description: String!
}
```

### Mutation

Operações que criam ou alteram dados, utilizam inputs:

```graphql
type Mutation {
  createCategory(input: NewCategory!): Category!
  createCourse(input: NewCourse!): Course!
}
```

### Type

Define a estrutura de dados retornados:

```graphql
type Category {
  id: ID! # Campo obrigatório
  name: String!
  description: String!
  courses: [Course!]! # Relacionamento (lista de cursos)
}

type Course {
  id: ID!
  name: String!
  description: String!
  category: Category! # Relacionamento (uma categoria)
}
```

**Nota:** O símbolo `!` indica que o campo é obrigatório (não pode ser nulo).

### Query

Define quais consultas são permitidas:

```graphql
type Query {
  categories: [Category!]!
  courses: [Course!]!
}
```

### Regenerando após mudanças no Schema

Sempre que modificar o `schema.graphqls`, execute:

```bash
go run github.com/99designs/gqlgen generate
```

Isso atualiza:

- `graph/model/models_gen.go`
- `graph/schema.resolvers.go`
- `graph/generated.go`

## Configuração do Banco de Dados

### 1. Criar o Banco de Dados SQLite

```bash
sqlite3 data.db
```

### 2. Criar as Tabelas

```sql
CREATE TABLE categories (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
);

CREATE TABLE courses (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  category_id TEXT NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

### 3. Injeção de Dependências

Em `graph/resolver.go`, as dependências são injetadas:

```go
type Resolver struct {
    CategoryDB *database.Category
    CourseDB   *database.Course
}
```

Em `cmd/server/server.go`, a conexão com o banco é estabelecida e as dependências são injetadas:

```go
db, err := sql.Open("sqlite3", "./data.db")
// ...
srv := handler.NewDefaultServer(
    graph.NewExecutableSchema(
        graph.Config{
            Resolvers: &graph.Resolver{
                CategoryDB: database.NewCategory(db),
                CourseDB:   database.NewCourse(db),
            },
        },
    ),
)
```

## Trabalhando com Models Customizados

### Por que customizar models?

Por padrão, o gqlgen gera todos os models em `models_gen.go`. No entanto, você pode querer:

- Adicionar métodos personalizados
- Implementar interfaces
- Controlar melhor os relacionamentos
- Separar a lógica de negócio

### Configurando models customizados

1. **Criar models separados** em `graph/model/`:

   - `category.go`
   - `course.go`

2. **Configurar o `gqlgen.yml`** para usar os models customizados:

```yaml
models:
  Category:
    model: github.com/seu-usuario/seu-repo/graph/model.Category
  Course:
    model: github.com/seu-usuario/seu-repo/graph/model.Course
```

3. **Remover campos de relacionamento** dos models customizados quando você quer que o gqlgen gere resolvers para eles.

## Implementando Relacionamentos

### Relacionamento Category → Courses (1:N)

1. **Remover o campo `Courses`** de `graph/model/category.go`:

```go
type Category struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    // Courses []*Course `json:"courses"` // REMOVER esta linha
}
```

2. **Regenerar o código**:

```bash
go run github.com/99designs/gqlgen generate
```

3. **Implementar o resolver** em `schema.resolvers.go`:

O gqlgen cria automaticamente o método `Courses()` no `CategoryResolver`:

```go
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
    // Buscar cursos relacionados à categoria
    return r.CourseDB.FindByCategoryID(obj.ID)
}
```

### Relacionamento Course → Category (N:1)

1. **Remover o campo `Category`** de `graph/model/course.go`

2. **Regenerar**:

```bash
go run github.com/99designs/gqlgen generate
```

3. **Implementar o resolver**:

```go
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
    // Buscar a categoria relacionada ao curso
    return r.CategoryDB.FindByID(ctx, obj.CategoryID)
}
```

### Vantagens desta abordagem

- **Lazy Loading**: Os relacionamentos são carregados apenas quando solicitados
- **N+1 Prevention**: Pode-se usar DataLoader para otimizar queries
- **Flexibilidade**: Controle total sobre como os dados são buscados
- **Testabilidade**: Fácil de mockar e testar

## Queries e Mutations

### Mutations

#### Criar Categoria

```graphql
mutation createCategory {
  createCategory(
    input: {
      # input envia os dados para criação
      name: "Tecnologia"
      description: "Cursos de tecnologia"
    }
  ) {
    # o que deseja que retorne da category
    id
    name
    description
  }
}
```

#### Criar Curso

```graphql
mutation createCourse {
  createCourse(
    input: {
      name: "Go Expert"
      description: "Curso completo de Go"
      categoryId: "4285232f-47cd-4d0e-8d76-b9fa0e5cc333"
    }
  ) {
    id
    name
    description
  }
}
```

### Queries

#### Listar Categorias

```graphql
query queryCategories {
  categories {
    id
    name
    description
  }
}
```

#### Listar Categorias com Cursos

```graphql
query queryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
      description
    }
  }
}
```

**Resultado:** Esta query retorna categorias com seus cursos relacionados (veja `image.png` para referência visual).

#### Listar Cursos

```graphql
query queryCourses {
  courses {
    id
    name
    description
  }
}
```

#### Listar Cursos com Categoria

```graphql
query queryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
```

**Resultado:** Esta query retorna cursos com suas categorias relacionadas (veja `image-1.png` para referência visual).

### Consultas Complexas

```graphql
query complexQuery {
  categories {
    id
    name
    courses {
      id
      name
      category {
        name
      }
    }
  }
}
```

## Executando o Servidor

### 1. Instalar dependências

```bash
go mod tidy
```

### 2. Iniciar o servidor

```bash
go run cmd/server/server.go
```

### 3. Acessar o Playground

Abra o navegador em: `http://localhost:8080/`

O GraphQL Playground permite:

- Executar queries e mutations
- Ver a documentação automática do schema
- Testar a API interativamente
- Ver histórico de requisições

### 4. Testar a API

Use o Playground para testar as queries e mutations documentadas acima.

## 🛠️ Comandos Úteis

```bash
# Gerar código após mudanças no schema
go run github.com/99designs/gqlgen generate

# Verificar o schema
go run github.com/99designs/gqlgen version

# Executar o servidor
go run cmd/server/server.go

# Acessar o banco de dados
sqlite3 data.db

# Limpar e reconstruir
go clean && go build ./...
```

## 📸 Referências Visuais

O projeto inclui imagens de referência mostrando os resultados das queries:

- `image.png` - Resultado da query `queryCategoriesWithCourses`
- `image-1.png` - Resultado da query `queryCoursesWithCategory`

## 📚 Recursos Adicionais

- [Documentação oficial do gqlgen](https://gqlgen.com/)
- [GraphQL Spec](https://spec.graphql.org/)
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)
- [DataLoader pattern](https://gqlgen.com/reference/dataloaders/)

## 🔍 Troubleshooting

### Erro ao gerar código

- Verifique se o `schema.graphqls` está sintaticamente correto
- Confirme que o `gqlgen.yml` aponta para os caminhos corretos

### Erro de importação circular

- Separe models em pacotes diferentes
- Use interfaces quando apropriado

### Performance em relacionamentos

- Implemente DataLoader para evitar N+1 queries
- Use batching quando possível

---

**Desenvolvido com ❤️ usando Go e GraphQL**
