package main

import (
	//"cmd/api"
	"database/sql"
	"log"

	"github.com/jankoskiantonio/simplebank/api"
	db "github.com/jankoskiantonio/simplebank/db/sqlc"
	"github.com/jankoskiantonio/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot laod config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}