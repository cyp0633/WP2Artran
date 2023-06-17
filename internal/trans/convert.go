package trans

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
	"github.com/cyp0633/wp-comment-converter/internal/fetch"
)


func getPageKey(post int) (pageKey string) {
	// e.g. https://new.example.com/archives/123
	url := "https://" + conf.Conf.New.Hostname + conf.Conf.Old.PermalinkPrefix + strconv.Itoa(post)
	// GET web content in url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bodyContent, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// find <meta http-equiv="refresh" content="0; url=https://new.exmaple.com/new-pagekey">
	regex := regexp.MustCompile(`<meta http-equiv="refresh" content="0; url=https://` + conf.Conf.New.Hostname + `/([^"]+)">`)
	matches := regex.FindSubmatch(bodyContent)
	if len(matches) != 2 {
		panic("no meta refresh found")
	}
	pageKey = string(matches[1])
	if pageKey[0] != '/' {
		pageKey = "/" + pageKey
	}
	return
}

func convertTime(from string) string {
	inputFormat := "2006-01-02T15:04:05"
	outputFormat := "2006-01-02 15:04:05 -0700 -0700"

	t, err := time.Parse(inputFormat, from)
	if err != nil {
		return ""
	}

	to := t.Format(outputFormat)

	return to
}
