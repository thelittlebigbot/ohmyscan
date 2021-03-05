package actions

import (
	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Download ...
func Download(c *cli.Context) {
	platform := c.String("platform")
	manga := c.String("manga")
	number := c.String("number")

	if platform == "" || manga == "" || number == "" {
		utils.Message(utils.ErrorArgumentsEmpty, "error")
	} else if platform == utils.PlatformScanOPName {
		platform := utils.PlatformScanOPName
		url := utils.PlatformScanOPURL + manga + "/" + number

		utils.DownloadFile(url, platform, manga, number)

	} else if platform == utils.PlatformJapScanName {
		utils.Message("Comming soon from Japscan.", "debug")
	}
}
