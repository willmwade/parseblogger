// parseblogger project parseblogger.go
package parseblogger

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type BloggerFeed struct {
	Url        string
	Limit      int64
	MinDate    time.Time
	MaxDate    time.Time
	XMLName    xml.Name   `xml:"feed"`
	Updated    string     `xml:"updated"`
	Categories []Category `xml:"category"`
	Id         string     `xml:"id"`
	Title      string     `xml:"title"`
	Subtitle   string     `xml:"subtitle"`
	Entries    []Entry    `xml:"entry"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

type Entry struct {
	XMLName    xml.Name   `xml:"entry"`
	Id         string     `xml:"id"`
	Title      string     `xml:"title"`
	Content    string     `xml:"content"`
	Author     Author     `xml:"author"`
	Published  string     `xml:"published"`
	Updated    string     `xml:"updated"`
	Categories []Category `xml:"category"`
}

type Author struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
	GPlus   string   `xml:"uri"`
	Email   string   `xml:"email"`
	Avitar  Avitar   `xml:"image"`
}

type Avitar struct {
	Src string `xml:"src,attr"`
}

func NewBloggerFeed(url string) *BloggerFeed {
	var bf BloggerFeed
	bf.Url = url
	return &bf
}

func (b *BloggerFeed) FetchUrl() string {
	url := b.Url + "/feeds/posts/default?"
	if b.Limit != 0 {
		url += "max-results=" + strconv.FormatInt(b.Limit, 10)
	} else {
		url += "max-results=100000"
	}

	if !b.MaxDate.IsZero() {
		url += "&" + b.MaxDate.Format("2006-01-02T15:04:05-07:00")
	}

	if !b.MinDate.IsZero() {
		url += "&" + b.MinDate.Format("2006-01-02T15:04:05-07:00")
	}
	return url
}

func (b *BloggerFeed) GetFeed(client *http.Client) error {
	xmlrsp, err := client.Get(b.FetchUrl())
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(xmlrsp.Body)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(body, b); err != nil {
		return err
	}
	return nil
}

func (a Avitar) Url() string {
	return "http:" + a.Src
}
