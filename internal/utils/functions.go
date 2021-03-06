package utils

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/TwinProduction/go-color"
	"github.com/anaskhan96/soup"
	"github.com/jung-kurt/gofpdf"
)

// Message ... Basic message
// 	- str(string)
// 	- level(string)
func Message(str string, level string) {
	// TODO: Create a level interface/type struct

	if level == "error" {
		log.Fatal(color.Colorize(ColorRed, str))
	}
	if level == "warning" {
		fmt.Println(color.Colorize(ColorYellow, str))
	}
	if level == "success" {
		fmt.Println(color.Colorize(ColorGreen, str))
	}
	if level == "debug" {
		fmt.Println(color.Colorize(ColorCyan, str))
	}
}

// GetContent ... Get content from the html page
// 	- url(string): Enter the html page url to handle content
func GetContent(url string) string {
	res, err := soup.Get(url)
	if err != nil {
		Message(err.Error(), "error")
	}

	return res
}

// GetLinks ... Get scans links from the content
// 	- content(string): Enter the content of an html page to handle scan links
// 	- platform(string): Enter the platform name
func GetLinks(content string, platform string) []string {
	var res []string

	doc := soup.HTMLParse(content)

	if platform == PlatformScanOPName {
		args := doc.Find("div", "id", "all")
		if args.Error != nil {
			Message(args.Error.Error(), "error")
		}

		img := args.FindAll("img")

		for _, elem := range img {
			res = append(res, elem.Attrs()["data-src"])
		}
	} else if platform == PlatformJapScanName {
		Message("Get links from japscan", "debug") // TODO: Create parser for japscan
	} else {
		Message(("GetLinks: " + ErrorPlatformInvalid), "error")
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

	if platform == PlatformScanOPName {
		res = split[4] + "/" + split[5]
		os.MkdirAll((Folder + res), os.ModePerm)
		Message(("The directory " + Folder + res + " was created."), "success")
	} else if platform == PlatformJapScanName {
		Message("Create a forlder for japscan", "debug") // TODO: Create folder for japscan
	} else {
		Message(ErrorPlatformInvalid, "error")
	}

	return res
}

// DownloadFile ... Download file
// 	- url(string)
// 	- platform(string)
// 	- manga(string)
// 	- number(uint8)
func DownloadFile(url string, platform string, name string, number string) {
	// TODO: Convent number in uint8

	content := GetContent(url)
	imgs := GetLinks(content, platform)
	if imgs == nil {
		Message("No images provided.", "error")
	}

	CreateFolder(url, platform)

	for _, elem := range imgs {
		split := strings.Split(elem, "/")
		path := Folder + name + "/" + number + "/" + split[(len(split)-1)]
		path = strings.TrimSpace(path)
		err := GetFile(strings.TrimSpace(elem), path)
		if err != nil {
			Message(err.Error(), "error")
		}

		Message(("The file" + elem + "is downloaded."), "debug")
	}

	Message(("The directory " + Folder + name + "/" + number + " was downloaded."), "success") // TODO: Better description
}

// ListFiles ...
// 	- dir(string)
func ListFiles(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		Message(err.Error(), "error")
	}

	return files
}

// RemoveFiles ...
//	- dir(string)
//	- name(string)
func RemoveFiles(dir string, name string) {
	err := os.Remove((dir + name))
	Message(("The file " + dir + name + " was deleted."), "debug")
	if err != nil {
		Message(err.Error(), "error")
	}
}

// ConvertToPDF ... Convert files to pdf and remove jpg
// 	- name(string)
// 	- number(string)
func ConvertToPDF(name string, number string) {
	dir := Folder + name + "/" + number + "/"
	files := ListFiles(dir)
	fName := name + "-" + number + ".pdf"

	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, f := range files {
		pdf.AddPage()

		pdf.ImageOptions(
			(dir + f.Name()),
			0, 0,
			0, 0,
			false,
			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
			0,
			"",
		)

		RemoveFiles(dir, f.Name())
	}

	err := pdf.OutputFileAndClose((dir + fName))
	if err != nil {
		Message(err.Error(), "error")
	}

	Message(("The file " + dir + fName + " was converted."), "success")
}
