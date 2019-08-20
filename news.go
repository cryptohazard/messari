package messari

import(

	"github.com/jedib0t/go-pretty/table"

)

//News are curated articles of news and analysis for cryptoassets.
type News struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	References []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"references"`
	ReferenceTitle string `json:"reference_title"`
	PublishedAt    string `json:"published_at"`
	Author         struct {
		Name string `json:"name"`
	} `json:"author"`
	Tags []interface{} `json:"tags"`
	URL  string        `json:"url"`
}

//AllNews is an array of News
type AllNews []News

func (a AllNews) String() string {
	tablewriter.AppendHeader(table.Row{"TITLE", "LINK", "TAGS"})
	for _, v := range a {
		tablewriter.AppendRow([]interface{}{v.Title, v.URL, v.Tags})
	}
	return tablewriter.Render()
}
