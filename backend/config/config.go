package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	URI string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Aviso: Não foi possível carregar o arquivo .env: %v", err)
	}

	viper.SetDefault("MONGODB_URI", "mongodb://localhost:27017")
	viper.SetDefault("API_PORT", "8080")

	viper.AutomaticEnv()
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".") 

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Aviso: Nenhum arquivo de configuração encontrado. Usando variáveis de ambiente ou padrões: %v", err)
	}

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("API_PORT"),
		},
		DB: DBConfig{
			URI: viper.GetString("MONGODB_URI"),
		},
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() APIConfig {
	return cfg.API
}
