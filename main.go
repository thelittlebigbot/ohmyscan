package main

import (
	"log"
	"os"

	"github.com/afigueir/ohmyscan/internal/actions"
	"github.com/afigueir/ohmyscan/internal/utils"
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
			Usage: utils.MetadataDownload,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: utils.MetadataFlagsName,
				}, &cli.StringFlag{
					Name:  "number",
					Usage: utils.MetadataFlagsNumber,
				}, &cli.StringFlag{
					Name:  "merge",
					Usage: utils.MetadataFlagsMerge,
				},
			},
			Action: func(c *cli.Context) error {
				actions.Download(c)
				return nil
			},
		},
		{
			Name:  "merge",
			Usage: utils.MetadataMerge,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: utils.MetadataFlagsName,
				}, &cli.StringFlag{
					Name:  "number",
					Usage: utils.MetadataFlagsNumber,
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
