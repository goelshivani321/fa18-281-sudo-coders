package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Represents database server and credentials
type Config struct {
	MongoURI         string
	RabbitmqServer   string
	RabbitmqPort     string
	RabbitmqUser     string
	RabbitmqPassword string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

//func GetHardConfig() Config {
//	config := Config{}
//	config.Server = "mongodb"
//	config.Database = "ridesDB"
//	return config
//}
//
//
//func GetOSConfig() Config {
//	config := Config{}
//	config.Server = os.Getenv("MONGO_SERVER")
//	config.Database = os.Getenv("MONGO_DB")
//	return config
//}
