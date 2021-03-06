package main

import (
	"log"
	"os"

	"github.com/thelittlebigbot/ohmyscan/internal/actions"
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  utils.MetadataName,
		Usage: utils.MetadataUsage,
	}

	app.Commands = []*cli.Command{
		{
			Name:  "download",
			Usage: "Download", // TODO: Add description
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "Select your manga",
				}, &cli.StringFlag{
					Name:  "number",
					Usage: "Select your manga",
				},
			},
			Action: func(c *cli.Context) error {
				actions.Download(c)
				return nil
			},
		},
		{
			Name:  "merge",
			Usage: "Merge", // TODO: Add description
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "Select your manga",
				}, &cli.StringFlag{
					Name:  "number",
					Usage: "Select your manga",
				},
			},
			Action: func(c *cli.Context) error {
				actions.Merge(c)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
