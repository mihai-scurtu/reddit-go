package main

import (
	"fmt"

	"github.com/mihai-scurtu/reddit-go/reddit"
)

func main() {
	reddit := reddit.NewClient("Testing Go Wrapper")

	fmt.Println(reddit.UserAgent)

	// fmt.Println(reddit.Get("/?limit=1"))

	postList := reddit.GetFrontPage()

	for _, p := range postList.GetChildren() {
		fmt.Println(p.Title)
	}
}
