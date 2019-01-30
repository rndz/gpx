package gpx

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	xml := `
<?xml version="1.0" encoding="ISO-8859-1" standalone="yes"?>
<gpx
 version="1.1"
 creator="EasyGPS 1.3.7 - http://www.topografix.com"
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns="http://www.topografix.com/GPX/1/1"
 xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
<metadata>
<time>2004-05-17T06:23:35Z</time>
<bounds minlat="37.407500" minlon="-122.478800" maxlat="47.621220" maxlon="-121.701276"/>
</metadata>
<wpt lat="37.820200000" lon="-122.478800000">
 <name>GGBSFB</name>
 <cmt>Golden Gate Bridge</cmt>
 <desc>Golden Gate Bridge - San Francisco Bay</desc>
 <sym>Waypoint</sym>
 <type>Waypoint</type>
</wpt>
<extensions>
</extensions>
</gpx>
	`
	r := strings.NewReader(xml)
	_, err := Parse(r)
	if err != nil {
		t.Errorf("Error while parsing gpx stream: %s", err)
	}
}

func TestParseOptionalEle(t *testing.T) {
	xml := `
<?xml version="1.0" encoding="ISO-8859-1" standalone="yes"?>
<gpx
 version="1.1"
 creator="EasyGPS 1.3.7 - http://www.topografix.com"
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns="http://www.topografix.com/GPX/1/1"
 xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
  <trk>
   <trkseg>
	  <trkpt lat="50.178127" lon="6.200096">
		<time>2016-05-03T07:19:11Z</time>
	  </trkpt>
	  <trkpt lat="50.178127" lon="6.200096">
		<ele>644.642</ele>
		<time>2016-05-03T07:19:17Z</time>
	  </trkpt>
	</trkseg>
  </trk>
</gpx>
	`
	r := strings.NewReader(xml)
	tracks, err := Parse(r)
	if err != nil {
		t.Errorf("Error while parsing gpx stream: %s", err)
	}

	if tracks.Trk[0].Trkseg[0].Trkpt[0].Ele.Valid {
		t.Errorf("expected unset 'ele' attribute, but was set!")
	}
	wpt1 := tracks.Trk[0].Trkseg[0].Trkpt[1]
	if !wpt1.Ele.Valid {
		t.Errorf("expected set 'ele' attribute, but was unset!")
	}
	if wpt1.Ele.Val != 644.642 {
		t.Errorf("invalid value for 'ele' attribute, expect %f, but was %f", 644.642, tracks.Trk[0].Trkseg[0].Trkpt[1].Ele.Val)
	}
}
