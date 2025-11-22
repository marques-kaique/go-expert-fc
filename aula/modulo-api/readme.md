A estrutura de respositorio vai seguir o layout de https://github.com/golang-standards/project-layout

/api -> Responavel por documentação, especificação da API
/internal -> responsavel por rodar a aplicação, não é uma pasta disponivel pois a regra de negocio se encontra nela
/pkg -> são libarias que voce permite que seja publica, por exemplo, uma lib auth, pode ser reaproveitada em outros projeto pelo go mod
/cmd -> aonde fica o projeto, no qual é gerado o executavo main.go - o local que será feito build, o correr seria ser composto por camada, cmd/server/main.go mas muitos deixam apenas em cmd
/configs -> local que fica as configurações do projeto, como variaveis de ambientes, configurações para subir o projeto
/test -> quando se possui arquivo adicionais para os testes, como arquivos e2e, documentação de teste, script (não necessariamente serão arquivos .go)



Para acessar o banco de dados
-> sqlite3 cmd/server/test.db