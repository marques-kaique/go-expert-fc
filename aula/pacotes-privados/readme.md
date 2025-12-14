# Trabalhando com Pacotes Privados e Dependências no Go

Em ambientes corporativos, é muito comum trabalhar com **pacotes internos** que **não devem ser públicos**. O Go possui mecanismos nativos para lidar corretamente com esse cenário, tanto no acesso a repositórios privados quanto no gerenciamento de dependências.

---

## Pacotes Privados e a Variável `GOPRIVATE`

Para informar ao Go que determinados repositórios são **privados**, é necessário configurar a variável de ambiente `GOPRIVATE`.

Essa variável indica ao Go quais domínios ou caminhos **não devem passar pelo Go Proxy público**, evitando erros de autenticação e problemas de acesso.

### Configurando o `GOPRIVATE`

```bash
export GOPRIVATE=github.com/seu-usuario-ou-org/*
```

> 💡 Você pode informar um domínio inteiro ou usar curingas (`*`) para abranger vários repositórios.

Após isso, o Go já consegue identificar quais módulos são privados.

---

## Autenticação para Repositórios Privados

Caso ocorram problemas de **acesso ou autenticação**, será necessário fornecer credenciais ao Git. Isso pode ser feito via **SSH** ou via **login + token**.

### Usando `.netrc` (login + token)

O arquivo `.netrc` é utilizado para armazenar credenciais de acesso.

Edite ou crie o arquivo:

```bash
vim ~/.netrc
```

Conteúdo esperado:

```text
machine github.com
login SEU_LOGIN
password SEU_TOKEN_DO_GITHUB
```

> 🔐 Recomenda-se utilizar **GitHub Personal Access Token (PAT)** em vez de senha.

---

## Configurando o Git para Usar SSH

Outra abordagem (e geralmente a mais recomendada) é configurar o Git para **substituir automaticamente URLs HTTPS por SSH**.

Isso pode ser feito no repositório específico:

```bash
vim .git/config
```

Ou de forma global:

```bash
vim ~/.gitconfig
```

Adicione a configuração abaixo:

```ini
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

### O que é o `insteadOf`?

Essa configuração indica que, **sempre que o Git encontrar uma URL HTTPS**, ele deverá **utilizar SSH automaticamente**, evitando problemas de autenticação.

---

## Garantindo Dependências no Projeto (`go mod vendor`)

Ao trabalhar com **pacotes externos**, existe o risco de o dono do repositório **privar ou remover o acesso** ao código no futuro.

O Go utiliza o **Go Proxy** para baixar e cachear dependências, porém isso **não é 100% confiável** em todos os cenários.

### Utilizando `go mod vendor`

Para garantir que todas as dependências estejam disponíveis dentro do projeto, utilize:

```bash
go mod vendor
```

Esse comando:

* Copia **todas as dependências utilizadas** para o diretório `vendor/`
* Evita problemas futuros com pacotes removidos ou privados
* Garante builds mais previsíveis

> 📦 Muito utilizado em ambientes corporativos e pipelines de CI/CD.

---

## Resumo

* Use `GOPRIVATE` para informar ao Go quais repositórios são privados
* Configure autenticação via `.netrc` ou SSH
* Utilize `insteadOf` para forçar o uso de SSH no Git
* Use `go mod vendor` para garantir dependências locais e estabilidade do projeto

---

✅ Seguindo essas práticas, você reduz problemas de autenticação, aumenta a segurança e garante maior previsibilidade no gerenciamento de dependências em projetos Go.
