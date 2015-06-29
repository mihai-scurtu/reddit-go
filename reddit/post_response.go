package reddit

import (
	"encoding/json"
	"log"
	_ "time"
)

type postResponse struct {
	Data Post
}

func NewPost(jsonBody []byte) *Post {
	var data map[string]interface{}

	err := json.Unmarshal(jsonBody, &data)
	if err != nil {
		log.Fatal(err)
	}

	data = data["data"].(map[string]interface{})

	return &Post{
	// Thing: Thing{
	// 	Id:   data["id"].(string),
	// 	Name: data["name"].(string),
	// },

	// Author:       data["author"].(string),
	// Created:      time.Unix(int64(data["created"].(float64)), 0),
	// Domain:       data["domain"].(string),
	// IsSelf:       data["is_self"].(bool),
	// IsSticky:     data["stickied"].(bool),
	// Permalink:    data["permalink"].(string),
	// Score:        int(data["score"].(float64)),
	// SelfText:     data["selftext"].(string),
	// SelfTextHtml: data["selftext_html"].(string),
	// Subreddit:    data["subreddit"].(string),
	// SubredditId:  data["subreddit_id"].(string),
	// Thumbnail:    data["thumbnail"].(string),
	// Title:        data["title"].(string),
	// Url:          data["url"].(string),
	// Downvotes:    int(data["downs"].(float64)),
	// Upvotes:      int(data["ups"].(float64)),
	}
}
