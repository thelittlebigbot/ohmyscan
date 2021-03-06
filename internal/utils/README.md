# Utils

## Table of contents

- [Utils](#utils)
  - [Table of contents](#table-of-contents)
  - [Constants](#constants)
  - [Functions](#functions)

## Constants

Location of the download directory.

```go
const Folder string
```

Lists all the colors used in the application.

```go
const (
	ColorReset  string
	ColorRed    string
	ColorGreen  string
	ColorYellow string
	ColorBlue   string
	ColorPurple string
	ColorCyan   string
	ColorGray   string
	ColorWhite  string
)
```

List of metadatas used in the application.

```go
const (
	MetadataName  string
	MetadataUsage string
)
```

List of the names and links of the platforms used in this application.

```go
const (
	PlatformScanOPName  string
	PlatformScanOPURL   string
	PlatformJapScanName string
	PlatformJapScanURL  string
)
```

List of recurring errors used in this application.

```go
const (
	ErrorPlatformInvalid string
	ErrorArgumentsEmpty  string
)
```

## Functions

Function that allows you to create a folder.

```go
func CreateFolder(url string, platform string) string
```

Function that allows you to convert multiple images into a single PDF.

```go
func ConvertToPdf(name string, number string)
```

Function that allows you to download documents online.

```go
func DownloadFile(url string, platform string, name string, number string)
```

Function that allows you to retrieve all the content of an HTML page.

```go
func GetContent(url string) string
```

Function that allows you to recover files from links.

```go
func GetFile(url string, dir string) error
```

Function that allows you to list all the image links present in an HTML document.

```go
func GetLinks(content string, platform string) []string
```

Function that allows you to list all the files in a folder.

```go
func ListFiles(dir string) []fs.Fileinfo
```

Function that allows you to display many messages, such as errors, etc.

```go
func Message(str string, level string)
```

Function that allows you to delete a file.

```go
func RemoveFile(dir string, name string)
```
