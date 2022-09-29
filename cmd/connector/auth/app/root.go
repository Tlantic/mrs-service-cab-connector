package app

import (
	"sort"

	"gopkg.in/urfave/cli.v1"
)

// NewCliApp ...
func NewCliApp() *cli.App {

	app := cli.NewApp()

	app.Name = "mrs-service-cab-connector-auth"
	app.Usage = "mrs-service-connector-auth for CAB"

	opts := NewAuthCmdOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {

		proc := NewAuthApp()
		return proc.Run(opts)
	}

	// sort flags by name
	sort.Sort(cli.FlagsByName(app.Flags))

	return app
}
