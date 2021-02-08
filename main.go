package main

import (
	"company-api/db"
	"company-api/router"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting Application.......")
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://akpovee:ajirioghene@cluster0.020uy.mongodb.net/people?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Client.Disconnect(ctx)

	router := router.NewRouter()

	router.Logger.Fatal(router.Start(":5000"))
}
