package config

type InternalConfig struct {
	RunningLocal bool
	ServerPort   int
	ServiceName  string
}

type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Config struct {
	InternalConfig *InternalConfig
	MySQLConfig    *MySQLConfig
}

func Getconfigs() *Config {
	return &Config{
		InternalConfig: &InternalConfig{
			RunningLocal: true,
			ServerPort: 8080,
			ServiceName: "Distopia",
		},
		MySQLConfig: &MySQLConfig{
			Host:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "root",
			Database: "distopia_example",
		},
		
	}
}
