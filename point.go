package xgeo

import (
	"database/sql/driver"
	"fmt"
	xerr "github.com/goclub/error"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
)

// Point GeoJSON  支持 mongo bson mysql
// GeoJson规范（RFC 7946）全文翻译: https://zhuanlan.zhihu.com/p/141554586
// xgeo.NewPoint(xgeo.WGS84{121.48294,31.2328}) // WGS84{经度,纬度}
type Point struct {
	Type string `json:"type" bson:"type"`
	// []float64{longitude, latitude} []float64{经度, 纬度}
	// 可能所有人都至少一次踩过这个坑：地理坐标点用字符串形式表示时是纬经（ "latitude,longitude" ），
	// 而数组形式表示时是经度在前，纬度在后（ [longitude,latitude] ）—顺序刚好相反。
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

func NewPointFormRaw(lngAndLat [2]float64) Point {
	return Point{
		"Point",
		lngAndLat,
	}
}
func NewPoint(data WGS84) Point {
	return Point{
		"Point",
		[2]float64{data.Longitude, data.Latitude},
	}
}
func (p Point) WGS84() WGS84 {
	return WGS84{
		Longitude: p.Coordinates[0],
		Latitude:  p.Coordinates[1],
	}
}

// String returns the WKT (Well Known Text) representation of the point.
// POINT(1 2)
func (p Point) String() string {
	return fmt.Sprintf("POINT(%f %f)", p.Coordinates[0], p.Coordinates[1])
}

func (p Point) Value() (value driver.Value, err error) {
	proxy := orb.Point{}
	proxy = p.Coordinates
	value, err = ewkb.ValuePrefixSRID(proxy, 4326).Value() // indivisible begin
	if err != nil {                                        // indivisible end
		return nil, xerr.WithStack(err)
	}
	return
}

// Scan implements the SQL driver.Scanner interface and will scan the
func (p *Point) Scan(data interface{}) (err error) {
	proxy := orb.Point{}
	err = ewkb.ScannerPrefixSRID(&proxy).Scan(data) // indivisible begin
	if err != nil {                                 // indivisible end
		err = xerr.WithStack(err)
		return
	}
	p.Type = "Point"
	p.Coordinates = proxy
	return
}
func (p Point) Validator() (err error) {
	data := p.WGS84()
	valid := -90 <= data.Latitude && data.Latitude <= 90 && -180 <= data.Longitude && data.Longitude <= 180
	if valid == false {
		return xerr.New(fmt.Sprintf("xgeo.Point{} invalid format"))
	}
	return
}
