package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFedd struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToRssFeed(url string) (RSSFedd, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := httpClient.Get(url)
	if err != nil {
		return RSSFedd{}, err
	}
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return RSSFedd{}, err
	}
	rssfeed := RSSFedd{}

	err = xml.Unmarshal(data, &rssfeed)
	if err != nil {
		return RSSFedd{}, err
	}
	return rssfeed, nil

}
