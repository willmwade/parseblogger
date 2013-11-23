// parseblogger project parseblogger_test.go
package parseblogger

import (
	"net/http"
	"testing"
	"time"
)

func TestNewFeed(t *testing.T) {
	url := "http://wilnregina.blogspot.com"
	newFeed := NewFeed(url)
	if newFeed.Url != url {
		t.Errorf("NewFeed not getting initialized correctly.")
	}
}

func TestGetFeed(t *testing.T) {
	url := "http://wilnregina.blogspot.com"
	newFeed := NewFeed(url)

	newFeed.Limit = 2

	var client http.Client

	err := newFeed.GetFeed(&client)
	if err != nil {
		t.Errorf("%v", err)
	} else if newFeed.Entries[0].Author.Avitar.Url() == "" {
		t.Errorf("Looks like the cache failed, at least on the Avitar URL")
	}

	newFeed = NewFeed(url)

	newFeed.Limit = 2
	newFeed.MaxDate = time.Now()

	err = newFeed.GetFeed(&client)
	if err != nil {
		t.Errorf("%v", err)
	} else if newFeed.Entries[0].Author.Avitar.Url() == "" {
		t.Errorf("Looks like the cache failed for max date")
	}
}
