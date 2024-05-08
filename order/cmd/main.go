package main

import (
	"log"

	"github.com/KrongkaewPon/grpc-go/order/config"
	"github.com/KrongkaewPon/grpc-go/order/internal/adapters/db"
	"github.com/KrongkaewPon/grpc-go/order/internal/adapters/grpc"
	"github.com/KrongkaewPon/grpc-go/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
