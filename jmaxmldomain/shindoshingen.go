package jmaxmldomain

import (
	"encoding/xml"
	"time"
)

type ShindoShingen struct {
	XMLName xml.Name      `xml:"Report"`
	Control ReportControl `xml:"Control"`
	Head    head          `xml:"Head"`
	Body    body          `xml:"Body"`
}

type head struct {
	Title           string    `xml:"Title"`
	ReportDateTime  time.Time `xml:"ReportDateTime"`
	TargetDateTime  time.Time `xml:"TargetDateTime"`
	EventID         string    `xml:"EventID"`
	InfoType        string    `xml:"InfoType"`
	Serial          int       `xml:"Serial"`
	InfoKind        string    `xml:"InfoKind"`
	InfoKindVersion string    `xml:"InfoKindVersion"`
	Headline        string    `xml:"Headline"`
}

type ShindoShingenHeadLine struct {
	Text string `xml:"Text"`
}

type body struct {
	Earthquake earthquake `xml:"Earthquake"`
	Intensity  intensity  `xml:"Intensity"`
}

type earthquake struct {
	OriginTime           time.Time  `xml:"OriginTime"`
	ArrivalTime          time.Time  `xml:"ArrivalTime"`
	HypocenterName       string     `xml:"Hypocenter>Area>Name"`
	HypocenterCode       string     `xml:"Hypocenter>Area>Code"`
	HypocenterCoordinate coordinate `xml:"Hypocenter>Coordinate"`
	Magnitude            magnitude  `xml:"Magnitude"`
}

type coordinate struct {
	Description string `xml:"Description,attr"`
	Content     string `xml:",chardata"`
}

type magnitude struct {
	Type        string  `xml:"type,attr"`
	Description string  `xml:"description,attr"`
	Content     float64 `xml:",chardata"`
}

type intensity struct {
	MaxInt string `xml:"Observation>MaxInt"`
}
