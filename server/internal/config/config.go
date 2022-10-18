package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigPostgres struct {
	PostgresHost string `env:"WB_PG_HOST" env-default:"localhost"`
	PostgresPort int    `env:"WB_PG_PORT" env-default:"5432"`
	UserName     string `env:"WB_PG_NAME" env-default:"user"`
	UserPassword string `env:"WB_PG_PASS" env-default:"password"`
	Database     string `env:"WB_PG_DB" env-default:"WB"`
}

type ConfigNats struct {
	NatsHost    string `env:"WB_NATS_HOST" env-default:"localhost"`
	NatsPort    int    `env:"WB_NATS_PORT" env-default:"4222"`
	ClusterID   string `env:"WB_CLUSTER_ID" env-default:"test-cluster"`
	ClientID    string `env:"WB_CLIENT_ID" env-default:"server"`
	SubjectName string `env:"WB_SUBJECT_NAME" env-default:"json-notification"`
}

type ConfigServer struct {
	HttpServerHost string `env:"WB_SERVER_HOST" env-default:"0.0.0.0"`
	HttpServerPort int    `env:"WB_SERVER_HOST" env-default:"80"`
}

type Configs struct {
	Postgres ConfigPostgres
	Nats     ConfigNats
	Server   ConfigServer
}

var cfg *Configs

func GetCfg() (*Configs, error) {
	cfg = &Configs{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
