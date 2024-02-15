package jmaxmldomain

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	XMLName  xml.Name  `xml:"feed"`
	Title    string    `xml:"title"`
	Subtitle string    `xml:"subtitle"`
	Updated  time.Time `xml:"updated"`
	Id       string    `xml:"id"`
	Links    []link    `xml:"link"`
	Rights   rights    `xml:"rights"`
	Entry    []entry   `xml:"entry"`
}

type link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type rights struct {
	Type    string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

type entry struct {
	Title   string       `xml:"title"`
	Id      string       `xml:"id"`
	Updated time.Time    `xml:"updated"`
	Author  author       `xml:"author"`
	Link    entryLink    `xml:"link"`
	Content entryContent `xml:"content"`
}

func (e *entry) GetDetail() {
	return
}

type entryLink struct {
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

type author struct {
	Name string `xml:"name"`
}

type entryContent struct {
	Type    string `xml:"type,attr"`
	Content string `xml:",chardata"`
}
