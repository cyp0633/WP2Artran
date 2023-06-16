package fetch

import (
	"io"
	"net/http"
	"strconv"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
)

type WPComment struct {
}

func FetchComments() (comments []WPComment) {
	for page := 1; ; page++ {
		body := request(page)
		newComments := parse(body)
		if len(newComments) < 20 {
			break
		}
		comments = append(comments, newComments...)
	}
	return
}

func request(page int) (body []byte) {
	url := conf.Conf.Old.Hostname + "/wp-json/wp/v2/comments?per_page=20&page=" + strconv.Itoa(page) + "&context=edit"
	user := conf.Conf.Auth.Username
	pass := conf.Conf.Auth.Password
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return
}

func parse(body []byte) (comments []WPComment) {
	return
}
