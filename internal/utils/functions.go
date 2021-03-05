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

	if platform == ScanOPName {
		args := doc.Find("div", "id", "all")
		if args.Error != nil {
			log.Fatal(args.Error)
		}

		img := args.FindAll("img")

		for _, elem := range img {
			res = append(res, elem.Attrs()["data-src"])
		}
	} else if platform == JapScanName {
		fmt.Println("Get links from japscan") // TODO: Create parser for japscan
	} else {
		log.Fatal("Invalid platform name in GetLinks function") // DEBUG
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
func CreateFolder(url string, platform string) string {
	var res string

	split := strings.Split(url, "/")

	if platform == ScanOPName {
		res = split[4] + "/" + split[5]
		os.MkdirAll((Folder + res), os.ModePerm)
		fmt.Printf("The directory %s was created.\n", Folder+res) // TODO: Add color
	} else if platform == JapScanName {
		fmt.Println("Create a forlder for japscan") // TODO: Create a folder for japscan
	} else {
		log.Fatal("Invalid platform name in CreateFolder function") // DEBUG
	}

	return res
}

// DownloadFile ... Download file
// 	- url(string)
// 	- platform(string)
// 	- manga(string)
// 	- number(uint8)
func DownloadFile(url string, platform string, manga string, number string) {
	// TODO: Convent number in uint8

	content := GetContent(url)
	imgs := GetLinks(content, platform)
	if imgs == nil {
		log.Fatal("No images provided.") // TODO: Add color
	}

	CreateFolder(url, platform)

	for _, elem := range imgs {
		split := strings.Split(elem, "/")
		path := Folder + manga + "/" + number + "/" + split[(len(split)-1)]
		path = strings.TrimSpace(path)
		err := GetFile(strings.TrimSpace(elem), path)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%s is downloaded.\n", elem) // TODO: Add color
		}
	}
}
