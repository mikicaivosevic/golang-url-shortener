# Url shortener written in Go Language

##Usage
```go
    url := UrlShortener{}
	url.short("http://www.example.com", TINY_URL)
	fmt.Println(url.ShortUrl)
	fmt.Println(url.OriginalUrl)
```