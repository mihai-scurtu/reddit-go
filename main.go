package main

import (
	"fmt"

	"github.com/mihai-scurtu/reddit-go/reddit"
)

func main() {
	Reddit := reddit.NewClient("Testing Go Wrapper")

	fmt.Println(Reddit.UserAgent)

	// fmt.Println(reddit.Get("/?limit=1"))

	postList := Reddit.GetFrontPage()

	for _, p := range postList.GetChildren() {
		fmt.Println(p.Title)
	}
}
