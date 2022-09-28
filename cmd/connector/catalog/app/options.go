package app

import (
	"gopkg.in/urfave/cli.v1"
)

const (
	Service = "SERVICE"
)

// CatalogCmdOptions ...
type CatalogCmdOptions struct {
	Service string
}

// NewCatalogCmdOptions ...
func NewCatalogCmdOptions() *CatalogCmdOptions {
	return &CatalogCmdOptions{}
}

//AddFlags ...

func (opts *CatalogCmdOptions) AddFlags(app *cli.App) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "svc",
			Usage:       "service",
			EnvVar:      Service,
			Value:       "",
			Destination: &opts.Service,
		},
	}

	app.Flags = append(app.Flags, flags...)
}
