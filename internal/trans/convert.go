package trans

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
	"github.com/cyp0633/wp-comment-converter/internal/fetch"
)

func Convert(wpComments []fetch.WPComment) {
	artrans := []Artran{}
	for _, wpComment := range wpComments {
		artrans = append(artrans, convert(wpComment))
	}
	writeToFile(artrans)
}

func convert(comment fetch.WPComment) (artran Artran) {
	artran.ID = comment.ID
	artran.RID = comment.Parent
	artran.Content = comment.Content.Rendered
	artran.UA = comment.AuthorUserAgent
	artran.IP = comment.AuthorIP
	artran.CreatedAt = convertTime(comment.DateGMT)
	artran.UpdatedAt = convertTime(comment.DateGMT)
	artran.IsCollapsed = false
	if comment.Status == "approved" {
		artran.IsPending = false
	} else {
		artran.IsPending = true
	}
	artran.VoteUp = 0
	artran.VoteDown = 0
	artran.Nickname = comment.AuthorName
	artran.Email = comment.AuthorEmail
	artran.Link = comment.AuthorURL
	artran.PageKey = getPageKey(comment.Post)
	return
}

func writeToFile(artrans []Artran) {
	bytes, err := json.Marshal(artrans)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(conf.Conf.OutputPath, bytes, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("Written to", conf.Conf.OutputPath)
}

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
		log.Println("Failed to find page key for post", post)
		return string(post)
	}
	pageKey = string(matches[1])
	if pageKey[0] != '/' {
		pageKey = "/" + pageKey
	}
	log.Println("Found page key", pageKey, "for post", post)
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
