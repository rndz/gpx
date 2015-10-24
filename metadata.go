package gpx

import "time"

// Metadata holds information about .gpx file.
type Metadata struct {
	Name      string     `xml:"name"`
	Desc      string     `xml:"desc"`
	Author    *Person    `xml:"author"`
	Copyright *Copyright `xml:"copyright"`
	Link      []*Link    `xml:"link"`
	Time      time.Time  `xml:"time"`
	Keywords  string     `xml:"keywords"`
	Bounds    *Bounds    `xml:"bounds"`
}

// Person represents person or organization
type Person struct {
	Name  string `xml:"name"`
	Email *Email `xml:"email"`
	Link  *Link  `xml:"link"`
}

// Email represents email address
type Email struct {
	ID     string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

// Link is a link to external resource with additional information
type Link struct {
	Href string `xml:"href,attr"`
	Text string `xml:"text"`
	Type string `xml:"type"`
}

// Copyright represents information about copyright holder and license.
type Copyright struct {
	Author  string `xml:"author,attr"`
	Year    string `xml:"year"`
	License string `xml:"license"`
}

// Bounds is two lat/lon pairs defining the extent of an element.
type Bounds struct {
	MinLat float64 `xml:"minlat,attr"`
	MinLon float64 `xml:"minlon,attr"`
	MaxLat float64 `xml:"maxlat,attr"`
	MaxLon float64 `xml:"maxlon,attr"`
}
