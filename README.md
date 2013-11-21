parseblogger
============

Simple Go Lang library to Parse Blogger Post Feeds without the need for an API key.

Notes
-----------
If you have a Google Application API Key it might be much easier to use the official (alpha) Go lang library http://code.google.com/p/google-api-go-client/

Example
----------
```
bloggerFeed := parseblogger.NewBloggerFeed("http://test.blogspot.com")
bloggerFeed.Limit = 6
bloggerFeed.GetFeed()

for _, e := range bloggerFeed.Entries {
  // Do something with those entries.
}
```
