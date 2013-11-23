parseblogger
============

Simple Go Lang library to Parse Blogger Post Feeds without the need for an API key.

Notes
-----------
If you have a Google Application API Key it might be much easier to use the official (alpha) Go lang library http://code.google.com/p/google-api-go-client/

Example
----------
```
feed := parseblogger.NewFeed("http://test.blogspot.com")
feed.Limit = 6
feed.GetFeed()

for _, e := range feed.Entries {
  // Do something with those entries.
}
```
