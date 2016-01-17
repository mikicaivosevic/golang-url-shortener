# Url shortener written in Go Language

###Url Shortener struct
```go
type UrlShortener struct {
	ShortUrl string
	OriginalUrl string
}
```

###Tiny Url Shortener
```go
func tinyUrlShortener(url string) (string, string) {
	tinyUrl := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", url)
	return getResponseData(tinyUrl), url
}
```

###Usage
```go
url := UrlShortener{}
url.short("http://www.example.com", TINY_URL)
//url.short("http://www.example.com", IS_GD)
fmt.Println(url.ShortUrl)
fmt.Println(url.OriginalUrl)
```