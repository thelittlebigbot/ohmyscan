package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

// // IsValidLink ... Check if is a valid link
// // 	- link(string)
// // 	- platform(string)
// func IsValidLink(link string, platform string) bool {
// 	res, err := url.Parse(link)

// 	if platform == "scan-op" {
// 		return err == nil && res.Scheme != "" && res.Host != "" && (res.Hostname() == "www.scan-op.cc" || res.Hostname() == "scan-op.cc")
// 	} else if platform == "japcan" {
// 		return err == nil && res.Scheme != "" && res.Host != "" && (res.Hostname() == "www.japscan.se" || res.Hostname() == "japscan.se")
// 	} else {
// 		log.Fatal("Invalid platform in IsInvalidScan")
// 		return false
// 	}
// }

// // RemoveInvalidLinks ... Remove invalid links
// // 	- links([]string)
// // 	- platform(string)
// func RemoveInvalidLinks(links []string, platform string) []string {
// 	for i, elem := range links {
// 		if IsValidLink(elem, platform) == false {
// 			links = append(links[:i], links[i+1:]...)
// 			fmt.Printf("%s is invalid.", elem)
// 		}
// 	}

// 	return links
// }

// CreateFolder ... Create a folder
// 	- name(string): Folder name
func CreateFolder(name string) {
	os.MkdirAll(name, os.ModePerm)
}
