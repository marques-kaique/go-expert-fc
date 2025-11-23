package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/configs"
	_ "github.com/marques-kaique/go-expert-fc/aula/modulo-api/docs"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/database"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/* anotações do swagger
/ para gerar a documentação, execute o comando na pasta raiz do projeto até o main
/ swag init -g /cmd/server/main.go 
*/

// @title Modulo API
// @version 1.0
// @description This is a sample server Modulo API server.
// @termsOfService http://swagger.io/terms/

// @contact.email [email protected]
// @contact.url http://anywhere.com.br
// @contact.name Qualquer nome

// @license.name qualquer nome license
// @license.url http://anywhere.com.br

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Println("Iniciando o servidor...")
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// productDB é uma instância de database.ProductInterface
	productDB := database.NewProduct(db)

	// productHandler é uma instância de handlers.ProductHandler
	// handler é responsável por receber a requisição, chamar o caso de uso e retornar a resposta
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, conf.TokenAuth, conf.JWTWExpresIn)

	//implementação default do go
	/*
		http.HandleFunc("/products", productHandler.CreateProduct)
	*/

	// implementação com chi
	r := chi.NewRouter()
	/* middleware é uma função que intercepta a requisição antes de chegar no handler
	/ seria assim: request -> middleware(usa os dados, realiza uma ação e no fim continua) -> handler -> response
	/ pode ser usado para log, autenticação, etc
	/ o chi já tem um middleware de log
	*/
	r.Use(middleware.Logger)
	// como criamos um proprio middleware, vamos utilizar ele
	r.Use(LogRequest)
	/* mesmo que sua aplicação caia, o recoverer vai garantir que o servidor continue rodando
	/ ele é um middleware de recuperação

	r.Use(middleware.Recoverer)
	*/

	/* agrupar rotas
	/ um middleware pode ser aplicado em um grupo de rotas
	/ por exemplo, um grupo de rotas que precisa de autenticação
	/ o middleware de autenticação pode ser aplicado apenas nesse grupo de rotas
	/ o chi já tem um middleware de autenticação, jwt verify
	/ quando recebe uma autenticação ele utiliza o conf.TokenAuth e guarda no context dele
	/ possibiltando recuperar essa instancia em qualquer lugar da aplicação
	*/
	r.Route("/products", func(r chi.Router) {
		// o verifier é o middleware de autenticação
		// ele procura o token no header da requisição, no cookie, no query param, etc
		r.Use(jwtauth.Verifier(conf.TokenAuth)) // ele pega o token da requisição e guarda no context
		r.Use(jwtauth.Authenticator)            // ele verifica se o token é valido
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	// para documentação do swagger
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)

	/* roteadores permitem atachar roteamentos das requições ate chegar no handler
	/ isso significa que o handler não precisa saber qual é a rota
	/ roteador do GO não permite trabalhar com variaveis
	/ de agrupar rotas
	/ de usar regex
	*/
}

// criando o proprio middleware
// next -> envia a requisição para o próximo middleware ou handler
func LogRequest(next http.Handler) http.Handler {
	// HandlerFunc é um tipo que implementa a interface Handler
	// ele retorna um Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		// envia a requisição para o próximo middleware ou handler
		next.ServeHTTP(w, r)
	})
}
