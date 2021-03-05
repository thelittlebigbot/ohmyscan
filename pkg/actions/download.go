package actions

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

// Download ...
func Download(c *cli.Context) {
	platform := c.String("platform")
	manga := c.String("manga")
	number := c.String("number")

	if platform == "" || manga == "" || number == "" {
		log.Fatal("You need to specify the platform, the manga and the number")
	}

	fmt.Println(platform, manga, number)
}
