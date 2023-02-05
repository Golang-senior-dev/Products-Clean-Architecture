package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henriquehendel/productCleanArchitecture/adapter/postgres"
	"github.com/henriquehendel/productCleanArchitecture/di"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	connection := postgres.GetConnection(ctx)
	defer connection.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(connection)

	router := mux.NewRouter()
	router.Handle("/products", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/products", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")

	log.Printf("LISTEN ON PORT: %v\n", port)

	http.ListenAndServe(fmt.Sprintf(":%v\n", port), router)
}
