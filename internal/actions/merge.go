package actions

import (
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Merge ...
func Merge(c *cli.Context) {
	manga := c.String("manga")
	number := c.String("number")

	name := "doubt" // DEBUG
	numb := "1"     // DEBUG

	utils.ConvertToPDF(name, numb)
	// TODO: Merge all pdf files in one file

	utils.Message((manga + number), "debug")
}
