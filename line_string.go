package xgeo

import (
	"database/sql/driver"
	"fmt"
	xerr "github.com/goclub/error"
	orb "github.com/paulmach/orb"
	ewkb "github.com/paulmach/orb/encoding/ewkb"
	"strings"
)

// LineString GeoJSON  支持 mongo bson mysql
// GeoJson规范（RFC 7946）全文翻译: https://zhuanlan.zhihu.com/p/141554586
// xgeo.NewLineString()
type LineString struct {
	Type        string       `json:"type" bson:"type"`
	Coordinates [][2]float64 `json:"coordinates" bson:"coordinates"`
}

// NewLineString
// xgeo.NewLineString([]xgeo.Point{
// 	xgeo.NewPoint(xgeo.WGS84{Longitude: 0, Latitude: 0}),
// 	xgeo.NewPoint(xgeo.WGS84{Longitude: 10, Latitude: 10}),
// 	xgeo.NewPoint(xgeo.WGS84{Longitude: 20, Latitude: 25}),
// })

func NewLineStringFormRaw(lngAndLat [][2]float64) LineString {
	return LineString{
		"LineString",
		lngAndLat,
	}
}
func NewLineString(points []Point) LineString {
	value := LineString{
		Type: "LineString",
	}
	for _, point := range points {
		value.Coordinates = append(value.Coordinates, point.Coordinates)
	}
	return value
}

// String returns the WKT (Well Known Text) representation of the point.
// LINESTRING(1.5 2.45,3.21 4)
func (p LineString) String() string {
	points := []string{}
	for _, item := range p.Coordinates {
		points = append(points, fmt.Sprintf("%f %f", item[0], item[1]))
	}
	return "LINESTRING(" + strings.Join(points, ",") + ")"
}
func (p LineString) Value() (value driver.Value, err error) {
	proxy := orb.LineString{}
	for _, point := range p.Coordinates {
		proxy = append(proxy, point)
	}
	value, err = ewkb.ValuePrefixSRID(proxy, 4326).Value() // indivisible begin
	if err != nil {                                        // indivisible end
		return nil, xerr.WithStack(err)
	}
	return
}

// Scan implements the SQL driver.Scanner interface and will scan the
func (p *LineString) Scan(data interface{}) (err error) {
	proxy := orb.LineString{}
	err = ewkb.ScannerPrefixSRID(&proxy).Scan(data) // indivisible begin
	if err != nil {                                 // indivisible end
		err = xerr.WithStack(err)
		return
	}
	p.Type = "LineString"
	for _, point := range proxy {
		p.Coordinates = append(p.Coordinates, point)
	}
	return
}
func (p LineString) Validator(custom ...error) (err error) {
	outError := xerr.New(fmt.Sprintf("xgeo.LineString{} invalid format"))
	if len(custom) != 0 {
		outError = custom[0]
	}
	for _, point := range p.Coordinates {
		longitude := point[0]
		latitude := point[1]
		valid := -90 <= latitude && latitude <= 90 && -180 <= longitude && longitude <= 180
		if valid == false {
			return outError
		}
	}
	return
}
