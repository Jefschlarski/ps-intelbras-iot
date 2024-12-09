package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Url  string
	Port uint
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	Database string
	Drive    string
}

func init() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.drive", "postgres")
	viper.SetDefault("api.url", "http://localhost")
	viper.SetDefault("api.port", "5555")
}

// Load reads the configuration file named "config.toml" located in the root directory.
// It populates the global variable "cfg" with the values read from the file.
//
// It returns an error if there was a problem reading the configuration file.
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.API = APIConfig{
		Url:  viper.GetString("api.url"),
		Port: viper.GetUint("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
		Drive:    viper.GetString("database.drive"),
	}

	return nil
}

// GetApiConfig returns the API configuration.
//
// No parameters.
// Returns an APIConfig struct.
func GetApiConfig() APIConfig {
	return cfg.API
}

// GetDbConfig returns the database configuration.
//
// No parameters.
// Returns a DBConfig struct.
func GetDbConfig() DBConfig {
	return cfg.DB
}
