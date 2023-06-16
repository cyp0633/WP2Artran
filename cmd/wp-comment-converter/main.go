package main

import (
	"os"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
)

func main() {
	// read the first cli parameter as confPath
	confPath := os.Args[1]
	conf.ParseConf(confPath)
}
