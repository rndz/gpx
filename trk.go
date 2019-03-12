package gpx

// Trk represents a track - an ordered list of points describing a path.
type Trk struct {
	Name   string    `xml:"name,omitempty"`
	Cmt    string    `xml:"cmt,omitempty"`
	Desc   string    `xml:"desc,omitempty"`
	Src    string    `xml:"src,omitempty"`
	Link   []*Link   `xml:"link,omitempty"`
	Number uint      `xml:"number,omitempty"`
	Type   string    `xml:"type,omitempty"`
	Trkseg []*Trkseg `xml:"trkseg,omitempty"`
}

// Trkseg holds a list of Track Points which are logically connected in order.
type Trkseg struct {
	Trkpt []*Wpt `xml:"trkpt"`
}
