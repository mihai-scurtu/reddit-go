package main

import (
	"fmt"
	"github.com/mihai-scurtu/reddit-go/reddit"
	"sort"
	"strings"
	_ "time"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func anagramHash(s string) string {
	return SortString(strings.Replace(s, " ", "", -1))
}

func main() {
	client := reddit.NewClient("Testing Go Wrapper")

	comments := client.GetComments("golang", "romania").GetChildren()

	fmt.Println(len(comments))

	for _, comment := range comments {
		fmt.Println(comment)
	}

	//var anagrams = make(map[string]string)
	//
	//fmt.Println(Reddit.UserAgent)
	//
	//for {
	//	postList := Reddit.GetNewPosts()
	//
	//	for _, p := range postList.GetChildren() {
	//		hash := anagramHash(p.Title)
	//		match := anagrams[hash]
	//
	//		if len(match) > 0 && match != p.Title {
	//			fmt.Println("!!! FOUND !!! '" + p.Title + "' = '" + match + "'")
	//		}
	//
	//		match = p.Title
	//	}
	//
	//	fmt.Println(".")
	//	time.Sleep(1 * time.Minute)
	//}
}
