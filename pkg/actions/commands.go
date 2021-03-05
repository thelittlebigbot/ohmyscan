package actions

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// TODO: Better usage's description

// Commands ...
func Commands() {
	app := &cli.App{
		Name:  "Oh My Scan",
		Usage: "Download locally your favorite french scans",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "download",
			Usage: "Choose platform",
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
				Download(c)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
