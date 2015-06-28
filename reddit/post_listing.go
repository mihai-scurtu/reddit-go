package reddit

type PostListing struct {
  Data struct {
    Children []PostResponse `json:children`
    After string `json:after`
    Before string `json:before`
  }
}

func (pl *PostListing) GetChildren() []Post {
  var list []Post

  for _, p := range pl.Data.Children {
    list = append(list, p.Data)
  }

  return list
}
