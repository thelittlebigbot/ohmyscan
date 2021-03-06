package actions

import (
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Download ... Function that allows us to download the desired scans according to the given parameters.
// 	- c(*cli.Content): Pass the application context.
func Download(c *cli.Context) {
	name := c.String("name")
	number := c.String("number")

	if name == "" || number == "" {
		utils.Message(utils.ErrorArgumentsEmpty, "error")
	}

	url := utils.PlatformURL + name + "/" + number
	utils.DownloadFile(url, name, number)
}
