package main

import (
	"fmt"
	"jmaxml/jmaxmldomain"
	"jmaxml/usecase"
)

func main() {
	req()
}

func req() {
	//url := "https://www.data.jma.go.jp/developer/xml/feed/extra.xml"
	urlEq := "https://www.data.jma.go.jp/developer/xml/feed/eqvol.xml"

	feed, err := usecase.GetFeed(urlEq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", feed.Title)
	for _, entry := range feed.Entry {
		fmt.Printf("  Title: %s\n", entry.Title)
		fmt.Printf("    Link: %s\n", entry.Link.Href)
		fmt.Printf("    Content: %s\n", entry.Content.Content)
	}
}

func req2() {
	url := "https://www.data.jma.go.jp/developer/xml/data/20240209095253_0_VXSE53_010000.xml"
	report, err := usecase.GetReport[jmaxmldomain.ShindoShingen](url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Title: %s\n", report.Control.Title)
	fmt.Printf("ReportDateTime: %s\n", report.Control.DateTime)
	fmt.Printf("Status: %s\n", report.Control.Status)
	fmt.Printf("EditorialOffice: %s\n", report.Control.EditorialOffice)
	fmt.Printf("PublishingOffice: %s\n", report.Control.PublishingOffice)
	fmt.Printf("MaxInt: %s\n", report.Body.Intensity.MaxInt)
}
