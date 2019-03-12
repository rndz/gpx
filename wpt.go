package gpx

import (
	"encoding/xml"
	"time"
)

// Wpt represents a waypoint, point of interest, or named feature on a map.
type Wpt struct {
	Lat          float64   `xml:"lat,attr"`
	Lon          float64   `xml:"lon,attr"`
	Ele          Elevation `xml:"ele,omitempty"`
	Time         time.Time `xml:"time,omitempty"`
	Magvar       float64   `xml:"magvar,omitempty"`
	GeoIDHeight  float64   `xml:"geoidheight,omitempty"`
	Name         string    `xml:"name,omitempty"`
	Cmt          string    `xml:"cmt,omitempty"`
	Desc         string    `xml:"desc,omitempty"`
	Src          string    `xml:"src,omitempty"`
	Link         []*Link   `xml:"link,omitempty"`
	Sym          string    `xml:"sym,omitempty"`
	Type         string    `xml:"type,omitempty"`
	Fix          string    `xml:"fix,omitempty"` // http://www.topografix.com/gpx/1/1/#type_fixType
	Sat          uint      `xml:"sat,omitempty"`
	Hdop         float64   `xml:"hdop,omitempty"`
	Vdop         float64   `xml:"vdop,omitempty"`
	Pdop         float64   `xml:"pdop,omitempty"`
	AgeOfGPSData float64   `xml:"ageofgpsdata,omitempty"`
	DGPSID       uint      `xml:"dgpsid,omitempty"`
}

type Elevation struct {
	Val   float64
	Valid bool
}

func (ele Elevation) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	if ele.Valid {
		return e.Encode(ele.Val)
	}
	return nil
}

func (ele *Elevation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ele.Valid = true
	return d.DecodeElement(&ele.Val, &start)
}
