package utils

// TODO: Add docstrings

// Folder ...
const Folder = "downloads/"

// Colors ...
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

// Metadata ...
const (
	MetadataName  string = "Oh My Scan"
	MetadataUsage string = "Download locally your favorite french scans."
)

// Platforms ...
const (
	PlatformScanOPName  string = "scan-op"
	PlatformScanOPURL   string = "https://scan-op.cc/manga/"
	PlatformJapScanName string = "japscan"
	PlatformJapScanURL  string = "https://japscan.se/lecture-en-ligne/"
)

// Errors ...
const (
	ErrorPlatformInvalid string = "Invalid platform name."                                      // TODO: Add color
	ErrorArgumentsEmpty  string = "You need to specify the platform, the manga and the number." // TODO: Add color
)
