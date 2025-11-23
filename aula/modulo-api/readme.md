A estrutura de respositorio vai seguir o layout de https://github.com/golang-standards/project-layout

/api -> Responsavel por documentação, especificação da API
/internal -> responsavel por rodar a aplicação, não é uma pasta disponivel pois a regra de negocio se encontra nela
/pkg -> são libarias que voce permite que seja publica, por exemplo, uma lib auth, pode ser reaproveitada em outros projeto pelo go mod
/cmd -> aonde fica o projeto, no qual é gerado o executavo main.go - o local que será feito build, o correr seria ser composto por camada, cmd/server/main.go mas muitos deixam apenas em cmd
/configs -> local que fica as configurações do projeto, como variaveis de ambientes, configurações para subir o projeto
/test -> quando se possui arquivo adicionais para os testes, como arquivos e2e, documentação de teste, script (não necessariamente serão arquivos .go)

Para acessar o banco de dados
-> sqlite3 cmd/server/test.db

Env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=goexpert
WEB_SERVER_PORT=8000
JWT_SECRET=your-secret-key-here
JWT_EXPIRESIN=300 ## tempo está em segundos, 5 minutos


Token JWT 

o token jwt é composto por 3 partes
antes do primeito ponto
- algoritmo utilizado para criptografia

após o primeiro ponto, dados que estão sendo transmitido, podendo ser:
  - sub -> costuma carregar o user_id nesse campo
  - name 
  - etc

Após o segundo ponto, temos a assinatura, o qual validade autenticidade do token, garante que ele não foi forçado, então aqui tem uma chave secreta que o sistema consegue validar se o token foi gerado pelo proprio sistema (pode ser utilizada chave RCA)

Quando token estiver valido mas expirado, pode ter um refresh token para gerar um novo token


go get -> para baixar dependencia que ficam localizado no go mod
go install -> baixar o arquivo binario para pode ser utilizado, todo os arquivo ficam na pasta /bin do GOPATH