package trans

import (
	"testing"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
)

func TestConvertTime(t *testing.T) {
	output := convertTime("2022-12-28T08:30:29")
	t.Log(output)
	if output != "2022-12-28 08:30:29 +0000 +0000" {
		t.Error("convertTime failed")
	}
}

func TestGetPageKey(t *testing.T) {
	conf.Conf.New.Hostname = "new.example.com"
	conf.Conf.Old.PermalinkPrefix = "/archives/"
	output := getPageKey(123)
	t.Log(output)
}
