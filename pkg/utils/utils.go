package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

// TODO: Docstrings for functions and returns

// GetContent ... Get content from the html page
// 	- url(string): Enter the html page url to handle content
func GetContent(url string) string {
	res, err := soup.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

// GetLinks ... Get scans links from the content
// 	- content(string): Enter the content of an html page to handle scan links
// 	- platform(string): Enter the platform name
func GetLinks(content string, platform string) []string {
	var res []string

	doc := soup.HTMLParse(content)

	if platform == "scan-op" {
		args := doc.Find("div", "id", "all")
		if args.Error != nil {
			log.Fatal("Unable to parse links")
		}

		img := args.FindAll("img")

		for _, elem := range img {
			res = append(res, elem.Attrs()["data-src"])
		}
	} else if platform == "japsan" {
		fmt.Println("Get links from japscan") // TODO: Create parser for japscan
	} else {
		log.Fatal("Invalid platform in GetLinks")
	}

	return res
}

// GetFile ... Get file from specifioed links
// 	- url(string): Specify the url
// 	- path(string): Specify the path
func GetFile(url string, path string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}

// CreateFolder ... Create a folder
// 	- name(string): Folder name
func CreateFolder(url string) string {
	split := strings.Split(url, "/")
	res := split[4] + "/" + split[5]
	os.MkdirAll(("downloads/" + res), os.ModePerm)

	return res
}
