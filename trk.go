package gpx

// Trk represents a track - an ordered list of points describing a path.
type Trk struct {
	Name   string    `xml:"name"`
	Cmt    string    `xml:"cmt"`
	Desc   string    `xml:"desc"`
	Src    string    `xml:"src"`
	Link   []*Link   `xml:"link"`
	Number uint      `xml:"number"`
	Type   string    `xml:"type"`
	Trkseg []*Trkseg `xml:"trkseg"`
}

// Trkseg holds a list of Track Points which are logically connected in order.
type Trkseg struct {
	Trkpt []*Wpt `xml:"trkpt"`
}
