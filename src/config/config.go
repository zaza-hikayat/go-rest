package config

var AppConfig Config

type AppConf struct {
	Environment string `mapstructure:"environment"`
	Name        string `mapstructure:"name"`
}

type HttpConf struct {
	Port    string `mapstructure:"port"`
	Timeout int    `mapstructure:"timeout"`
}

type LogConf struct {
	Name string `mapstructure:"name"`
}

type DatabaseConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

type JWTConf struct {
	TokenSecret   string `mapstructure:"tokenSecret"`
	RefreshSecret string `mapstructure:"refreshSecret"`
}

type BasicAuthConf struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Config ...
type Config struct {
	App       AppConf       `mapstructure:"app"`
	Http      HttpConf      `mapstructure:"http"`
	Log       LogConf       `mapstructure:"log"`
	DB        DatabaseConf  `mapstructure:"database"`
	JWT       JWTConf       `mapstructure:"jwt"`
	BasicAuth BasicAuthConf `mapstructure:"basic_auth"`
}

func Make() Config {
	config := Config{
		App: AppConf{
			Environment: GetEnv("APP_ENVIRONMENT", "DEVELOPMENT"),
			Name:        GetEnv("APP_NAME", "REST API"),
		},
		Http: HttpConf{
			Port:    GetEnv("HTTP_PORT", "8080"),
			Timeout: GetEnvInt("HTTP_TIMEOUT", 360),
		},
		JWT: JWTConf{
			TokenSecret:   GetEnv("JWT_TOKEN_SECRET", ""),
			RefreshSecret: GetEnv("JWT_REFRESH_SECRET", ""),
		},
		Log: LogConf{
			Name: GetEnv("LOG_NAME", "REST API"),
		},
		DB: DatabaseConf{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnvInt("DB_PORT", 5432),
			Username: GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "postgres"),
			DBName:   GetEnv("DB_NAME", "go-rest"),
			SSLMode:  GetEnv("SSL_MODE", "disable"),
			Timezone: GetEnv("TIMEZONE", "Asia/Jakarta"),
		},
		BasicAuth: BasicAuthConf{
			Username: GetEnv("BASIC_AUTH_USER", ""),
			Password: GetEnv("BASIC_AUTH_PASSWORD", ""),
		},
	}

	AppConfig = config

	return config
}
