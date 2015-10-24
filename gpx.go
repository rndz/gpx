// gpx is a package to parse .gpx file as defined in http://www.topografix.com/GPX/1/1
package gpx

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

// Gpx is a root node in .gpx document.
type Gpx struct {
	XMLName  xml.Name  `xml:"gpx"`
	Version  string    `xml:"version,attr"`
	Creator  string    `xml:"creator,attr"`
	Metadata *Metadata `xml:"metadata"`
	Wpt      []*Wpt    `xml:"wpt"`
	Rte      []*Rte    `xml:"rte"`
	Trk      []*Trk    `xml:"trk"`
}

// Parse returns a Gpx struct from gpx stream
func Parse(r io.Reader) (*Gpx, error) {
	return parse(r)
}

// ParseFile returns a Gpx struct from gpx file
func ParseFile(path string) (*Gpx, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	return parse(file)
}

func parse(r io.Reader) (*Gpx, error) {
	var gpx Gpx
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&gpx)
	if err != nil {
		return nil, fmt.Errorf("could not decode gpx to struct: %s", err)
	}
	return &gpx, nil
}
