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

// Message ... Function that allows you to display many messages, such as errors, etc.
// 	- str(string): Content of the message.
// 	- level(string): Alert level of the message.
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

// GetContent ... Function that allows you to retrieve all the content of an HTML page.
// 	- url(string): URL of the HTML page.
func GetContent(url string) string {
	res, err := soup.Get(url)
	if err != nil {
		Message(err.Error(), "error")
	}

	return res
}

// GetLinks ... Function that allows you to list all the image links present in an HTML document.
// 	- content(string): Content of the html page to retrieve image links.
func GetLinks(content string) []string {
	var res []string

	doc := soup.HTMLParse(content)

	args := doc.Find("div", "id", "all")
	if args.Error != nil {
		Message(args.Error.Error(), "error")
	}

	img := args.FindAll("img")

	for _, elem := range img {
		res = append(res, elem.Attrs()["data-src"])
	}

	return res
}

// GetFile ... Function that allows you to recover files from links.
// 	- url(string): URL of the HTML page.
// 	- dir(string): Folder where the images will be downloaded.
func GetFile(url string, dir string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}

// CreateFolder ... Function that allows you to create a folder.
// 	- name(string): Name of the folder to be created.
func CreateFolder(url string) string {
	split := strings.Split(url, "/")

	res := split[4] + "/" + split[5]
	os.MkdirAll((Folder + res), os.ModePerm)
	Message(("The directory " + Folder + res + " was created."), "success")

	return res
}

// DownloadFile ... Function that allows you to download documents online.
// 	- url(string): URL of the HTML page.
// 	- name(string): Name of the manga to download.
// 	- number(string): Volume number to download.
func DownloadFile(url string, name string, number string) {
	// TODO: Convent number in uint8

	content := GetContent(url)
	imgs := GetLinks(content)
	if imgs == nil {
		Message("No images provided.", "error")
	}

	CreateFolder(url)

	for _, elem := range imgs {
		split := strings.Split(elem, "/")
		path := Folder + name + "/" + number + "/" + split[(len(split)-1)]
		path = strings.TrimSpace(path)
		err := GetFile(strings.TrimSpace(elem), path)
		if err != nil {
			Message(err.Error(), "error")
		}

		Message(("The file" + elem + "was downloaded."), "debug")
	}

	Message(("The volume " + number + " of " + name + " was downloaded."), "success")
}

// ListFiles ... Function that allows you to list all the files in a folder.
// 	- dir(string): Name of the folder to list.
func ListFiles(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		Message(err.Error(), "error")
	}

	return files
}

// RemoveFile ... Function that allows you to delete a file.
//	- dir(string): Name of the folder where the files to be deleted are located.
//	- name(string): Name of the file to be deleted.
func RemoveFile(dir string, name string) {
	err := os.Remove((dir + name))
	Message(("The file " + dir + name + " was deleted."), "debug")
	if err != nil {
		Message(err.Error(), "error")
	}
}

// ConvertToPDF ... Function that allows you to convert multiple images into a single PDF.
// 	- name(string): Name of the manga to convert.
// 	- number(string): Volume number to convert.
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
			210, 0,
			false,
			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
			0,
			"",
		)

		RemoveFile(dir, f.Name())
	}

	err := pdf.OutputFileAndClose((dir + fName))
	if err != nil {
		Message(err.Error(), "error")
	}

	Message(("The volumne " + number + " of " + name + " was converted."), "success")
}
