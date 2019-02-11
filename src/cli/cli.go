package cli

import (
	"data/datatype"
	"flag"
)

func Cli() (parsed *datatype.Parse) {
	var (
		url    string
		data   string
		file   string
		spider bool
	)

	flag.StringVar(&url, "u", "http://example.com", "Target URL")
	flag.StringVar(&data, "d", "", "POST data")
	flag.StringVar(&file, "f", "", "fuzz keyword file")
	flag.BoolVar(&spider, "s", false, "use spider")
	flag.Parse()

	Parsed := &datatype.Parse{}
	Parsed.Url = url
	Parsed.Data = data
	Parsed.File = file
	Parsed.Spider = spider

	return Parsed
}
