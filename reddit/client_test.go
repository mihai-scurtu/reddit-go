package reddit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientConstructor(t *testing.T) {
	assert := assert.New(t)

	userAgent := "test"

	reddit := NewClient(userAgent)

	assert.Equal(userAgent, reddit.UserAgent)
}

func TestUrl(t *testing.T) {
	var uri, expected string
	assert := assert.New(t)

	uri = "/test"
	expected = URL + uri
	assert.Equal(expected, RedditUrl(uri), "accepts uri beginning with slash")

	uri = "test"
	expected = URL + "/" + uri
	assert.Equal(expected, RedditUrl(uri), "accepts uri with no slash")

	uri = ""
	expected = URL + "/"
	assert.Equal(expected, RedditUrl(uri), "accepts empty string")
}

func TestRequestCreation(t *testing.T) {
	assert := assert.New(t)
	userAgent := "test"
	reddit := NewClient(userAgent)

	url := RedditUrl("")

	request := reddit.createRequest("GET", url, nil)

	assert.Equal(userAgent, request.Header.Get("User-Agent"), "it has the correct user agent")
	assert.Equal("GET", request.Method, "it has the correct method")
	assert.Equal("/"+url, request.URL.Path, "it has the correct url")
}
