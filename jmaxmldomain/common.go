package jmaxmldomain

import "time"

type ReportControl struct {
	Title            string    `xml:"Title"`
	DateTime         time.Time `xml:"DateTime"`
	Status           string    `xml:"Status"`
	EditorialOffice  string    `xml:"EditorialOffice"`
	PublishingOffice string    `xml:"PublishingOffice"`
}
