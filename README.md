# Go library to parse gpx files

So far it only maps gpx file to a struct (as defined in http://www.topografix.com/GPX/1/1) and is able to count track total distance.

Usage:
```
import "github.com/rndz/gpx"

func main() {
    g, err := gpx.ParseFile("file.gpx")
	if err != nil {
		panic(err)
	}
    // Print track distance(km).
	fmt.Println(g.Distance())

    // Use track point information
	for _, track := range g.Trk {
		for _, segment := range track.Trkseg {
			for _, pt := range segment.Trkpt {
                // Do something with pt.Lat, pt.Lon, etc...
			}
		}
	}
}

```
