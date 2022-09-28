package main

import (
	"fmt"
	"os"

	client "github.com/Tlantic/mrs-service-cab-connector/cmd/connector/catalog/app"
)

//go:generate protoc --go_out=plugins=grpc:../../../proto -I ../../../proto  ../../../proto/connector.proto ../../../proto/stock_types.proto

func main() {
	app := client.NewCliApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
