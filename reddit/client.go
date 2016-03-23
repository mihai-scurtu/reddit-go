package reddit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

	json.Unmarshal(resp, &l)

	return l
}

func (c *Client) GetNewPosts() *PostListing {
	var l *PostListing

	resp := c.Get("/new")

	json.Unmarshal(resp, &l)

	return l
}

func RedditUrl(uri string) string {
	if len(uri) == 0 || uri[0] != '/' {
		uri = "/" + uri
	}

	return URL + uri
}
