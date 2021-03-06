package actions

import (
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Download ...
func Download(c *cli.Context) {
	platform := c.String("platform")
	name := c.String("name")
	number := c.String("number")

	if platform == "" || name == "" || number == "" {
		utils.Message(utils.ErrorArgumentsEmpty, "error")
	}

	if platform == utils.PlatformScanOPName {
		url := utils.PlatformScanOPURL + name + "/" + number
		utils.DownloadFile(url, utils.PlatformScanOPName, name, number)
	}

	if platform == utils.PlatformJapScanName {
		utils.Message("Comming soon from Japscan.", "debug")
	}
}
