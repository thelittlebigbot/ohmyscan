package actions

import (
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Merge ...
func Merge(c *cli.Context) {
	name := c.String("name")
	number := c.String("number")

	if name == "" || number == "" {
		utils.Message("You need to specify the name and the number.", "error")
	}

	utils.ConvertToPDF(name, number)
}
