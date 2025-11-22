package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// essa var é do tipo conf
// o proposito dela é ajustar as configurações do projeto no momento da subida/inicialização
// não precisa ser nivel de pacote, pode ser nivel de função LoadConfig
var cfg *conf

// poderia ser serparado as struct por dominio, ex: db, web, jwt
// por segurança, é uma boa prática não expor as configurações
// pois os campos podem ser alterados por engano, o correto é expor apenas o necessário e o restante ficar privado
// remomendavel que se utiliza o encapsulamento, utilizando apenas o metodo get
type conf struct {
	DBDriver      string           `mapstructure:"DB_DRIVER"`
	DBHost        string           `mapstructure:"DB_HOST"`
	DBPort        string           `mapstructure:"DB_PORT"`
	DBUser        string           `mapstructure:"DB_USER"`
	DBPassword    string           `mapstructure:"DB_PASSWORD"`
	DBName        string           `mapstructure:"DB_NAME"`
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string           `mapstructure:"JWT_SECRET"`
	JWTWExpresIn  int              `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth // roteador?
}

// init é uma função especial do go que é executada antes do main
// ela é executada antes do main e serve para inicializar as variáveis
// que serão utilizadas no main
// func init()

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")  // tipo do arquivo de configuração
	viper.AddConfigPath(path)   // caminho do arquivo de configuração
	viper.SetConfigFile(".env") // o nome do arquivo de configuração
	viper.AutomaticEnv()        // caso tenha, substitui os valores de .env por variáveis de ambiente

	err := viper.ReadInConfig() // lê o arquivo de configuração
	if err != nil {
		panic(err) // a aplicação não vai subir caso tenha problema nesse processo
	}

	err = viper.Unmarshal(&cfg) // mapeia o arquivo de configuração para a struct conf
	if err != nil {
		panic(err)
	}

	// HS256 é o algoritmo de criptografia
	// []byte(cfg.JWTSecret) é a chave de criptografia
	// nil é o validador de token
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
