package actions

import (
	"fmt"
	"log"

	"github.com/thelittlebigbot/ohmyscan/internal/utils"
	"github.com/urfave/cli/v2"
)

// Download ...
func Download(c *cli.Context) {
	platform := c.String("platform")
	manga := c.String("manga")
	number := c.String("number")

	if platform == "" || manga == "" || number == "" {
		log.Fatal("You need to specify the platform, the manga and the number")
	} else if platform == utils.ScanOPName {
		platform := utils.ScanOPName
		url := utils.ScanOPURL + manga + "/" + number

		utils.DownloadFile(url, platform, manga, number)

	} else if platform == utils.JapScanName {
		fmt.Println("Comming soon from japscan") // TODO: Download files from japscan
	}
}
