package utils

// Folder ... Location of the download directory.
const Folder string = "downloads/"

// Colors ... Lists all the colors used in the application.
const (
	ColorReset  string = "\033[0m"
	ColorRed    string = "\033[31m"
	ColorGreen  string = "\033[32m"
	ColorYellow string = "\033[33m"
	ColorBlue   string = "\033[34m"
	ColorPurple string = "\033[35m"
	ColorCyan   string = "\033[36m"
	ColorGray   string = "\033[37m"
	ColorWhite  string = "\033[97m"
)

// Metadata ... List of metadatas used in the application.
const (
	MetadataName  string = "Oh My Scan"
	MetadataUsage string = "Download locally your favorite french scans."
)

// Platforms ... List of the name and link of the platform used in this application.
const (
	PlatformName string = "scan-op"
	PlatformURL  string = "https://scan-op.cc/manga/"
)

// Errors ... List of recurring errors used in this application.
const (
	ErrorArgumentsEmpty string = "You need to specify the the manga's name and the volume's number."
)
