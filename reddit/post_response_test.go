package reddit

import (
	"encoding/json"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

const JSON = "{\"kind\": \"t3\", \"data\": {\"domain\": \"i.imgur.com\", \"banned_by\": null, \"media_embed\": {}, \"subreddit\": \"funny\", \"selftext_html\": null, \"selftext\": \"\", \"likes\": null, \"suggested_sort\": null, \"user_reports\": [], \"secure_media\": null, \"link_flair_text\": null, \"id\": \"38snhz\", \"from_kind\": null, \"gilded\": 0, \"archived\": false, \"clicked\": false, \"report_reasons\": null, \"author\": \"preh1storic\", \"media\": null, \"score\": 3347, \"approved_by\": null, \"over_18\": false, \"hidden\": false, \"num_comments\": 221, \"thumbnail\": \"http://b.thumbs.redditmedia.com/Bd52w-oxg4cC1-oBo0P1PcQUY1x-O8ud2kGenGcPHOw.jpg\", \"subreddit_id\": \"t5_2qh33\", \"edited\": false, \"link_flair_css_class\": null, \"author_flair_css_class\": null, \"downs\": 0, \"secure_media_embed\": {}, \"saved\": false, \"removal_reason\": null, \"stickied\": false, \"from\": null, \"is_self\": false, \"from_id\": null, \"permalink\": \"/r/funny/comments/38snhz/jon_stewart_responds_in_the_only_logical_way/\", \"name\": \"t3_38snhz\", \"created\": 1433627489.0, \"url\": \"http://i.imgur.com/IzR95Du.png\", \"author_flair_text\": null, \"title\": \"Jon Stewart responds in the only logical way\", \"created_utc\": 1433598689.0, \"distinguished\": null, \"mod_reports\": [], \"visited\": false, \"num_reports\": null, \"ups\": 3347}}"

// const JSON = `"{
//   \"kind\": \"t3\",
//   \"data\": {
//     \"domain\": \"i.imgur.com\",
//     \"banned_by\": null,
//     \"media_embed\": {

//     },
//     \"subreddit\": \"funny\",
//     \"selftext_html\": null,
//     \"selftext\": \"\",
//     \"likes\": null,
//     \"suggested_sort\": null,
//     \"user_reports\": [

//     ],
//     \"secure_media\": null,
//     \"link_flair_text\": null,
//     \"id\": \"38snhz\",
//     \"from_kind\": null,
//     \"gilded\": 0,
//     \"archived\": false,
//     \"clicked\": false,
//     \"report_reasons\": null,
//     \"author\": \"preh1storic\",
//     \"media\": null,
//     \"score\": 3347,
//     \"approved_by\": null,
//     \"over_18\": false,
//     \"hidden\": false,
//     \"num_comments\": 221,
//     \"thumbnail\": \"http://b.thumbs.redditmedia.com/Bd52w-oxg4cC1-oBo0P1PcQUY1x-O8ud2kGenGcPHOw.jpg\",
//     \"subreddit_id\": \"t5_2qh33\",
//     \"edited\": false,
//     \"link_flair_css_class\": null,
//     \"author_flair_css_class\": null,
//     \"downs\": 0,
//     \"secure_media_embed\": {

//     },
//     \"saved\": false,
//     \"removal_reason\": null,
//     \"stickied\": false,
//     \"from\": null,
//     \"is_self\": false,
//     \"from_id\": null,
//     \"permalink\": \"/r/funny/comments/38snhz/jon_stewart_responds_in_the_only_logical_way/\",
//     \"name\": \"t3_38snhz\",
//     \"created\": 1433627489.0,
//     \"url\": \"http://i.imgur.com/IzR95Du.png\",
//     \"author_flair_text\": null,
//     \"title\": \"Jon Stewart responds in the only logical way\",
//     \"created_utc\": 1433598689.0,
//     \"distinguished\": null,
//     \"mod_reports\": [

//     ],
//     \"visited\": false,
//     \"num_reports\": null,
//     \"ups\": 3347
//   }
// }"`

func TestPostConstructor(t *testing.T) {
	assert := assert.New(t)

	var r PostResponse

	json.Unmarshal([]byte(JSON), &r)

	assert.Equal(r.Data.Id, "38snhz", "it correctly decodes id")

	fmt.Println(r)
}
