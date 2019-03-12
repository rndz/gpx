package gpx

// Rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination.
type Rte struct {
	Name   string  `xml:"name,omitempty"`
	Cmt    string  `xml:"cmt,omitempty"`
	Desc   string  `xml:"desc,omitempty"`
	Src    string  `xml:"src,omitempty"`
	Link   []*Link `xml:"link,omitempty"`
	Number uint    `xml:"number,omitempty"`
	Type   string  `xml:"type,omitempty"`
	Rtept  []*Wpt  `xml:"rtept,omitempty"`
}
