package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TODO: Add docstring and create the content of the function

// Merge ...
func Merge(c *cli.Context) {
	manga := c.String("manga")
	number := c.String("number")

	fmt.Println("Merge", manga, number)
}
