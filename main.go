package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/command"
)

const (
	DEFAULT_ENDPOINT = "https://localhost/adp/rest/api/task/executeAdpTask"
	DEFAULT_USER     = "adpuser"

	// Task: Start Application
	DEFAULT_APPLICATIONURL = "appURL"
)

func main() {

	app := &cli.App{
		Name:    "adp-cli",
		Version: "0.1-alpha",
		Usage:   "CLI for ADP rest API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "endpoint",
				Aliases: []string{"e"},
				Usage:   "REST service endpoint",
				Value:   DEFAULT_ENDPOINT,
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "ADP User",
				Value:   DEFAULT_USER,
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "ADP User Password",
				EnvVars: []string{"ADP_USER_PASSWORD"},
			},
			&cli.StringFlag{
				Name:    "taskAccessKey",
				Aliases: []string{"k"},
				Usage:   "ADP Task Access Key (for Query Postgresql DB Task)",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Debug Mode: Log to stderr and AdpTask log enabled",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "pretty",
				Aliases: []string{"b"},
				Usage:   "Pretty Print",
				Value:   false,
			},
		},
		Commands: command.Commands,
		Before: func(c *cli.Context) error {
			client.ADP.RC = client.NewRestClient(
				client.WithEndPoint(c.String("endpoint")),
				client.WithUser(c.String("user")),
				client.WithPassword(c.String("password")),
				client.WithTaskAccessKey(c.String("taskAccessKey")),
			)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
