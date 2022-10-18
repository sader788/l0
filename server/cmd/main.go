package main

import (
	"WildberriesL0/server/internal/cache"
	"WildberriesL0/server/internal/config"
	"WildberriesL0/server/internal/handlers"
	"WildberriesL0/server/pkg/nats-streaming"
	"WildberriesL0/server/pkg/postgres"
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})

	cfg, err := config.GetCfg()
	if err != nil {
		l.Fatal("config: " + err.Error())
	}
	l.Info("config: configuration init")

	// PG DB CONNECT
	pgm, err := postgres.ConnectDB(&cfg.Postgres)
	if err != nil {
		l.Fatal("postgresDB: " + err.Error())
	}
	l.Info("postgresDB: connected to the database")
	defer pgm.Close(context.Background())

	// CACHE INIT
	cache, err := cache.CacheInit(pgm, l)
	if err != nil {
		l.Fatal("order cache: " + err.Error())
	}
	l.Info("order cache: cache init, cache size: " + strconv.Itoa(cache.CacheLen()))

	// NATS STREAMING REGISTER SERVICE
	nm := natsstreaming.NewNatsManager()
	err = nm.Register(&cfg.Nats, handlers.StanHandler(cache, l))
	if err != nil {
		l.Fatal("nats-streaming: " + err.Error())
	}
	l.Info("nats-streaming: nats-streaming connected, sub chanel: " + cfg.Nats.SubjectName)
	defer nm.Unregister()

	router := httprouter.New()
	handler := handlers.RouterHandler(cache)
	handler.Register(router)
	l.Info("httprouter: router init")

	err = startHttpServer(router, l, &cfg.Server)
	if err != nil {
		l.Fatal("httpserver: " + err.Error())
	}
}

func startHttpServer(r *httprouter.Router, l *logrus.Logger, cfg *config.ConfigServer) error {

	servAddr := cfg.HttpServerHost + ":" + strconv.Itoa(cfg.HttpServerPort)

	listener, err := net.Listen("tcp", servAddr)

	if err != nil {
		return err
	}

	server := http.Server{
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	l.Info("httpserver: http server started " + servAddr)
	l.Fatalln(server.Serve(listener))
	return nil
}
