package gpx

import "math"

const earthRadius = 6371 // km

// Distance returns gpx track distance
func (g *Gpx) Distance() float64 {
	dist := 0.0
	for _, trk := range g.Trk {
		for _, seg := range trk.Trkseg {
			var prev, curr *location
			for _, p := range seg.Trkpt {
				if prev == nil {
					prev = &location{lat: p.Lat, lon: p.Lon}
					continue
				}
				curr = &location{lat: p.Lat, lon: p.Lon}
				dist += distanceBetween(prev, curr)
				prev = curr
			}
		}
	}
	return dist
}

type location struct {
	lat, lon float64
}

// for more info of what is happening here: https://en.wikipedia.org/wiki/Haversine_formula
func distanceBetween(l1, l2 *location) float64 {
	p1la, p1lo := l1.toRad()
	p2la, p2lo := l2.toRad()
	return 2 * earthRadius * math.Asin(math.Sqrt(haversine(p2la-p1la)+math.Cos(p1la)*math.Cos(p2la)*haversine(p2lo-p1lo)))
}

func (l *location) toRad() (float64, float64) {
	return l.lat * math.Pi / 180, l.lon * math.Pi / 180
}

func haversine(theta float64) float64 {
	return .5 * (1 - math.Cos(theta))
}
