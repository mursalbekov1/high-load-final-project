package models

type AppConfigs struct {
	Postgres    Postgres
	SwaggerAuth SwaggerAuth
	Logs        LogsConfig
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type LogsConfig struct {
	Path string
}

type SwaggerAuth struct {
	Username string
	Password string
}
