package main

import (
	"log"

	"shadowsoni666/db"
	"shadowsoni666/lib/config"
	logger "shadowsoni666/lib/logger"
	"shadowsoni666/note"
	server "shadowsoni666/server/http"
)

func main() {

	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	l := logger.New(config.LogLevel)

	sqldb, err := db.NewSQL("postgres", config.DBConnString, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	s := note.NewService(sqldb)

	r, err := server.NewChiRouter(s, config.PASETOSecret, config.AccessTokenDuration, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	httpServer, err := server.NewHTTP(r, config.HTTPServerAddress, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Fatal().Err(err).Send()
	}
}
