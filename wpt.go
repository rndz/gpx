package gpx

import (
	"encoding/xml"
	"time"
)

// Wpt represents a waypoint, point of interest, or named feature on a map.
type Wpt struct {
	Lat          float64   `xml:"lat,attr"`
	Lon          float64   `xml:"lon,attr"`
	Ele          Elevation `xml:"ele"`
	Time         time.Time `xml:"time"`
	Magvar       float64   `xml:"magvar"`
	GeoIDHeight  float64   `xml:"geoidheight"`
	Name         string    `xml:"name"`
	Cmt          string    `xml:"cmt"`
	Desc         string    `xml:"desc"`
	Src          string    `xml:"src"`
	Link         []*Link   `xml:"link"`
	Sym          string    `xml:"sym"`
	Type         string    `xml:"type"`
	Fix          string    `xml:"fix"` // http://www.topografix.com/gpx/1/1/#type_fixType
	Sat          uint      `xml:"sat"`
	Hdop         float64   `xml:"hdop"`
	Vdop         float64   `xml:"vdop"`
	Pdop         float64   `xml:"pdop"`
	AgeOfGPSData float64   `xml:"ageofgpsdata"`
	DGPSID       uint      `xml:"dgpsid"`
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
