package cmd

import (
	"github.com/urfave/cli/v3"
)

var RootCommand = &cli.Command{
	Name:  "Welding",
	Usage: "Welding is a management system",
	Commands: []*cli.Command{
		serveCommand,
		migrateCommand,
		seedCommand,
	},
}
