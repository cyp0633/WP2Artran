package main

import (
	"os"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
	"github.com/cyp0633/wp-comment-converter/internal/fetch"
	"github.com/cyp0633/wp-comment-converter/internal/trans"
)

func main() {
	// read the first cli parameter as confPath
	confPath := os.Args[1]
	conf.ParseConf(confPath)
	comments := fetch.FetchComments()
	trans.Convert(comments)
	return
}
