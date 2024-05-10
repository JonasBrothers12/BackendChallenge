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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Config struct {
	InternalConfig *InternalConfig
	MySQLConfig    *MySQLConfig
	RedisConfig    *RedisConfig
}

func Get() *Config {
	

	return &Config{
		MySQLConfig: &MySQLConfig{
			Host:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "root",
			Database: "pupper_example",
		},
		
	}
}
