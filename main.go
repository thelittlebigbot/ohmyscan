package main

import (
	"log"
	"os"

	"github.com/thelittlebigbot/ohmyscan/internal/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Oh My Scan",
		Usage: "Download locally your favorite french scans",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "download",
			Usage: "Download", // TODO: Add description
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "platform",
					Usage: "Select your platform",
				}, &cli.StringFlag{
					Name:  "manga",
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
					Name:  "manga",
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
