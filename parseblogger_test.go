// parseblogger project parseblogger_test.go
package parseblogger

import (
	"testing"
	"time"
)

func TestNewBloggerFeed(t *testing.T) {
	url := "http://wilnregina.blogspot.com"
	newBloggerFeed := NewBloggerFeed(url)
	if newBloggerFeed.Url != url {
		t.Errorf("NewBloggerFeed not getting initialized correctly.")
	}
}

func TestGetFeed(t *testing.T) {
	url := "http://wilnregina.blogspot.com"
	newBloggerFeed := NewBloggerFeed(url)

	newBloggerFeed.Limit = 2

	err := newBloggerFeed.GetFeed()
	if err != nil {
		t.Errorf("%v", err)
	} else if newBloggerFeed.Entries[0].Author.Avitar.Url() == "" {
		t.Errorf("Looks like the cache failed, at least on the Avitar URL")
	}

	newBloggerFeed = NewBloggerFeed(url)

	newBloggerFeed.Limit = 2
	newBloggerFeed.MaxDate = time.Now()

	err = newBloggerFeed.GetFeed()
	if err != nil {
		t.Errorf("%v", err)
	} else if newBloggerFeed.Entries[0].Author.Avitar.Url() == "" {
		t.Errorf("Looks like the cache failed for max date")
	}
}
