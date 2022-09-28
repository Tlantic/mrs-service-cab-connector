package main

import (
	"fmt"
	"os"

	client "github.com/Tlantic/mrs-service-cab-connector/cmd/connector/auth/app"
)

//go:generate protoc --go_out=plugins=grpc:../../proto -I ../../proto  ../../proto/auth_types.proto ../../proto/auth_external.proto

func main() {
	app := client.NewCliApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
