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


func Convert(comments []fetch.WPComment) {
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
