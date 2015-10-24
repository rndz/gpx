package gpx

// Rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination.
type Rte struct {
	Name   string  `xml:"name"`
	Cmt    string  `xml:"cmt"`
	Desc   string  `xml:"desc"`
	Src    string  `xml:"src"`
	Link   []*Link `xml:"link"`
	Number uint    `xml:"number"`
	Type   string  `xml:"type"`
	Rtept  []*Wpt  `xml:"rtept"`
}
