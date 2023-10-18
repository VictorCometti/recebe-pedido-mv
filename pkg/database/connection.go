package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getStringConnection() (connectionString string, erro error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("pkg/config")

	if erro = viper.ReadInConfig(); erro != nil {
		erro = fmt.Errorf("erro ao tentar ler o árquivo de conexão: %v", erro)
		return
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	hostname := viper.GetString("hostname")
	port := viper.GetString("port")
	database := viper.GetString("database_name")

	connectionString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"

	return

}

func GetConnection() (connection *sql.DB, erro error) {
	connectionString, erro := getStringConnection()

	if erro != nil {
		erro = fmt.Errorf("erro ao tentar pegar a string de conexão: %v", erro)
		return
	}

	connection, erro = sql.Open(
		"postgres",
		connectionString,
	)

	if erro != nil {
		erro = fmt.Errorf("erro ao tentar abrir a conexão com banco de dados: %v", erro)
		return
	}

	if erro = connection.Ping(); erro != nil {
		erro = fmt.Errorf("erro ao tentar pegar o ping de conexão: %v", erro)
		return
	}

	return
}
