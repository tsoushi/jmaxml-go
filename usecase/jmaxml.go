package usecase

import (
	"encoding/xml"
	"io"
	"jmaxml/jmaxmldomain"
	"net/http"
)

func GetFeed(url string) (*jmaxmldomain.Feed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feed jmaxmldomain.Feed
	xml.Unmarshal(body, &feed)
	return &feed, nil
}

func GetReport[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var report T
	xml.Unmarshal(body, &report)
	return &report, nil
}
