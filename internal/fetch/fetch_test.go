package fetch

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cyp0633/wp-comment-converter/internal/conf"
)

func TestParse(*testing.T) {
	comment := parse(exampleComment)
	b, _ := json.Marshal(comment)
	fmt.Println(string(b))
}

func TestRequest(*testing.T) {
	conf.Conf.Old.Hostname = "example.com"
	conf.Conf.Auth.Username = "test"
	conf.Conf.Auth.Password = "ABCD abcd 1234 5678"
	body := request(1)
	fmt.Println(string(body))
}

var exampleComment = []byte(`
[
    {
        "id": 123,
        "post": 123,
        "parent": 121,
        "author": 0,
        "author_name": "test",
        "author_url": "",
        "date": "2022-12-22T16:30:29",
        "date_gmt": "2022-12-28T08:30:29",
        "content": {
            "rendered": "<p>测试<\/p>\n"
        },
        "link": "https:\/\/example.com\/archives\/123#comment-123",
        "status": "approved",
        "type": "comment",
        "author_avatar_urls": {
            "24": "https:\/\/dn-qiniu-avatar.qbox.me\/avatar\/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?s=24&r=pg&d=404",
            "48": "https:\/\/dn-qiniu-avatar.qbox.me\/avatar\/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?s=48&r=pg&d=404",
            "96": "https:\/\/dn-qiniu-avatar.qbox.me\/avatar\/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?s=96&r=pg&d=404"
        },
        "meta": [],
        "_links": {
            "self": [
                {
                    "href": "https:\/\/example.com\/wp-json\/wp\/v2\/comments\/123"
                }
            ],
            "collection": [
                {
                    "href": "https:\/\/example.com\/wp-json\/wp\/v2\/comments"
                }
            ],
            "up": [
                {
                    "embeddable": true,
                    "post_type": "post",
                    "href": "https:\/\/example.com\/wp-json\/wp\/v2\/posts\/123"
                }
            ],
            "in-reply-to": [
                {
                    "embeddable": true,
                    "href": "https:\/\/example.com\/wp-json\/wp\/v2\/comments\/123"
                }
            ]
        }
    }
]`)
