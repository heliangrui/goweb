package config

import (
	"os"

	"github.com/spf13/viper"
)

// Env EnvironmentVariable 环境变量
var Env = environmentVariable{}

type Server struct {
	Name string
	Port string
}

type Postgres struct {
	hostname string
	port     string
	username string
	password string
	database string
}

type MqttConfig struct {
	Broker   string `地址`
	ClientId string `客户端id`
	Username string `用户名`
	Password string `密码`
}

type environmentVariable struct {
	EcityosRoot string //项目根路径
	EcityosRes  string //项目资源根目录
}

type MqttUser struct {
	Url      string
	User     string
	Password string
}

func init() {
	viper.SetConfigFile("./././config/application.yaml")
	viper.ReadInConfig()
	Env.EcityosRoot = os.Getenv("ECITYOS_ROOT")
	Env.EcityosRes = os.Getenv("ECITYOS_RES")
}

func GetMqttConfig() MqttConfig {
	broker := viper.GetString("mqtt.broker")
	clientId := viper.GetString("mqtt.clientId")
	username := viper.GetString("mqtt.username")
	password := viper.GetString("mqtt.password")
	mqtt := MqttConfig{broker, clientId, username, password}
	return mqtt
}

func GetServer() Server {
	servername := viper.GetString("server.name")
	port := viper.GetString("server.port")
	server := Server{servername, port}
	return server
}
func GetPostgres() Postgres {
	hostname := viper.GetString("postgres.hostname")
	port := viper.GetString("postgres.port")
	username := viper.GetString("postgres.username")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.database")
	postgres := Postgres{hostname, port, username, password, database}
	return postgres
}

func GetMqttUser() MqttUser {
	url := viper.GetString("mqtt-user.url")
	user := viper.GetString("mqtt-user.user")
	password := viper.GetString("mqtt-user.password")
	mqttUser := MqttUser{url, user, password}
	return mqttUser
}
