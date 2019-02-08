package cli

import (
	"flag"
)


type Parse struct{
    url string
    data string
    file string
    spider bool
}

func Cli()(parsed *Parse){
	var (
		url string
		data string
		file string
        spider bool
	)

	flag.StringVar(&url, "u", "http://example.com", "Target URL")
	flag.StringVar(&data, "d", "", "POST data")
	flag.StringVar(&file, "f", "", "fuzz keyword file")
    flag.BoolVar(&spider, "s", false, "use spider")
	flag.Parse()

    Parsed := &Parse{}
    Parsed.url = url
    Parsed.data = data
    Parsed.file = file
    Parsed.spider = spider

    return
}
