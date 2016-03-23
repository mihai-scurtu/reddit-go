package reddit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"fmt"
)

const URL = "https://api.reddit.com"

var httpClient = &http.Client{}

type Client struct {
	UserAgent string
}

func NewClient(userAgent string) *Client {

	return &Client{UserAgent: userAgent}
}

func (c *Client) createRequest(method string, uri string, body io.Reader) *http.Request {
	url := RedditUrl(uri)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", c.UserAgent)

	return req
}

func (c *Client) Get(uri string) []byte {
	req := c.createRequest("GET", uri, nil)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func (c *Client) GetFrontPage() *PostListing {
	var l *PostListing

	resp := c.Get("/")

	if err := json.Unmarshal(resp, &l); err !=nil {
		log.Fatal(err)
	}

	return l
}

func (c *Client) GetNewPosts() *PostListing {
	var l *PostListing

	resp := c.Get("/new")

	if err := json.Unmarshal(resp, &l); err !=nil {
		log.Fatal(err)
	}

	return l
}

func (c *Client) GetComments(subreddits ...string) *CommentListing {
	var l *CommentListing

	uri := "/r/" + strings.Join(subreddits, "+") + "/comments"

	resp := c.Get(uri)

	fmt.Println(string(resp))

	if err := json.Unmarshal(resp, &l); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", l)

	return l
}

func RedditUrl(uri string) string {
	if len(uri) == 0 || uri[0] != '/' {
		uri = "/" + uri
	}

	return URL + uri
}
