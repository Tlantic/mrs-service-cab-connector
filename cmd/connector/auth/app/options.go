package app

import (
	"gopkg.in/urfave/cli.v1"
)

const (
	Service = "SERVICE"
)

// AuthCmdOptions ...
type AuthCmdOptions struct {
	Service string
}

// NewAuthCmdOptions ...
func NewAuthCmdOptions() *AuthCmdOptions {
	return &AuthCmdOptions{}
}

//AddFlags ...

func (opts *AuthCmdOptions) AddFlags(app *cli.App) {
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
