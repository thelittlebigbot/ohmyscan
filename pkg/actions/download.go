package actions

import (
	"fmt"
	"log"
	"strings"

	"github.com/thelittlebigbot/ohmyscan/pkg/utils"
	"github.com/urfave/cli/v2"
)

// Download ...
func Download(c *cli.Context) {
	platform := c.String("platform")
	manga := c.String("manga")
	number := c.String("number")

	if platform == "" || manga == "" || number == "" {
		log.Fatal("You need to specify the platform, the manga and the number")
	} else if platform == "scan-op" {
		var url = "https://scan-op.cc/manga/" + manga + "/" + number

		html := utils.GetContent(url)
		imgs := utils.GetLinks(html, "scan-op")
		if imgs == nil {
			log.Fatal("No images provided")
		}

		utils.CreateFolder(url)

		for _, elem := range imgs {
			split := strings.Split(elem, "/")
			name := "downloads/" + manga + "/" + number + "/" + split[(len(split)-1)]
			name = strings.TrimSpace(name)
			fmt.Println(split, name)

			err := utils.GetFile(strings.TrimSpace(elem), name)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%s is downloaded.\n", elem)
			}
		}
	}
}
