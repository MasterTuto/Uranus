package flags

import "flag"

func ParseFilename() string {
	var filenameFlag string

	flag.StringVar(&filenameFlag, "filename", "0", "Name of the json file to parse")
	flag.Parse()

	return filenameFlag
}
