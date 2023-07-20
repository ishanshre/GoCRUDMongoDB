package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/connections"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/handlers"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/routers"
)

func main() {
	dsn := flag.String("dsn", "mongodb://localhost:27017", "URL for mongo db")
	port := flag.Int("port", 8000, "Port that server listen to")
	addr := fmt.Sprintf(":%d", *port)

	// connect to mongo db
	client, err := connections.ConnectToNoSql(*dsn)
	if err != nil {
		log.Fatalln("error in connecting to MongoDB")
	}

	// connect to handler interface
	h := handlers.NewHandler(client)

	// connect to router
	router := routers.Router(h)

	log.Printf("Starting server at port :%d", *port)

	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
