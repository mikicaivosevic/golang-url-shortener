package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

const (
	TINY_URL = 1
	IS_GD = 2
)

type UrlShortener struct {
	ShortUrl string
	OriginalUrl string
}

type Shortener interface {
	short(url string)
}

func getResponseData(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	return string(contents)
}

func tinyUrlShortener(url string) (string, string) {
	tinyUrl := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", url)
	return getResponseData(tinyUrl), url
}

func isGdShortener(url string) (string, string) {
	isGdUrl := fmt.Sprintf("http://is.gd/create.php?url=%s&format=simple", url)
	return getResponseData(isGdUrl), url
}

func (u *UrlShortener) short(url string, shortener int) *UrlShortener {
	switch shortener {
	case TINY_URL:
		shortUrl, originalUrl := tinyUrlShortener(url)
		u.ShortUrl = shortUrl
		u.OriginalUrl = originalUrl
		return u
	case IS_GD:
		shortUrl, originalUrl := isGdShortener(url)
		u.ShortUrl = shortUrl
		u.OriginalUrl = originalUrl
		return u
	}
	return u
}


func main()  {
	url := UrlShortener{}
	url.short("http://www.mondo.rs", IS_GD)
	fmt.Println(url.ShortUrl)
	fmt.Println(url.OriginalUrl)
}