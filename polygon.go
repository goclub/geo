package xgeo

import (
	"database/sql/driver"
	"fmt"
	xerr "github.com/goclub/error"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"strings"
)

// Polygon GeoJSON  支持 mongo bson
// GeoJson规范（RFC 7946）全文翻译: https://zhuanlan.zhihu.com/p/141554586
// xgeo.NewPolygon([]xgeo.WGS84{{-1,-1}, {1, -1}, {1, 1}, {-1, -1}) // WGS84{经度,纬度}
type Polygon struct {
	Type        string         `json:"type" bson:"type"`
	Coordinates [][][2]float64 `json:"coordinates" bson:"coordinates"`
}

func NewPolygonFormRaw(lngAndLat [][][2]float64) Polygon {
	return Polygon{
		"Polygon",
		lngAndLat,
	}
}

// NewPolygon new Polygon
func NewPolygon(lineStrings []LineString) Polygon {
	value := Polygon{
		Type: "Polygon",
	}
	for _, line := range lineStrings {
		value.Coordinates = append(value.Coordinates, line.Coordinates)
	}
	return value
}

// String returns the WKT (Well Known Text) representation of the point.
// POLYGON((1 2,1 4,3 4,3 2,1 2))
func (p Polygon) String() string {
	var lineStrings []string
	for _, lineString := range p.Coordinates {
		points := []string{}
		for _, item := range lineString {
			points = append(points, fmt.Sprintf("%f %f", item[0], item[1]))
		}
		lineStrings = append(lineStrings, "("+strings.Join(points, ",")+")")
	}

	return "POLYGON(" + strings.Join(lineStrings, ",") + ")"
}

func (p Polygon) Value() (value driver.Value, err error) {
	proxy := orb.Polygon{}
	for _, lineString := range p.Coordinates {
		ls := orb.Ring{}
		for _, point := range lineString {
			ls = append(ls, point)
		}
		proxy = append(proxy, ls)
	}
	value, err = ewkb.ValuePrefixSRID(proxy, 4326).Value() // indivisible begin
	if err != nil {                                        // indivisible end
		return nil, xerr.WithStack(err)
	}
	return
}

// Scan implements the SQL driver.Scanner interface and will scan the
func (p *Polygon) Scan(data interface{}) (err error) {
	proxy := orb.Polygon{}
	err = ewkb.ScannerPrefixSRID(&proxy).Scan(data) // indivisible begin
	if err != nil {                                 // indivisible end
		err = xerr.WithStack(err)
		return
	}
	p.Type = "Polygon"
	for _, ls := range proxy {
		newLs := [][2]float64{}
		for _, point := range ls {
			newLs = append(newLs, point)
		}
		p.Coordinates = append(p.Coordinates, newLs)
	}
	return
}

func (p Polygon) Validator() (err error) {
	for _, lineString := range p.Coordinates {
		for _, point := range lineString {
			longitude := point[0]
			latitude := point[1]
			valid := -90 <= latitude && latitude <= 90 && -180 <= longitude && longitude <= 180
			if valid == false {
				return xerr.New(fmt.Sprintf("xgeo.Polygon{} invalid format"))
			}
		}
	}
	return
}
