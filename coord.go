package xgeo

import (
	"fmt"
	xconv "github.com/goclub/conv"
	xerr "github.com/goclub/error"
	"github.com/goclub/geo/internal/coord"
)

// WGS84 实现了 sql 接口 driver.Valuer sql.Scanner
type WGS84 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude  float64 `json:"latitude" note:"纬度"`
}

func (data WGS84) Validator() (err error) {
	valid := -90 <= data.Latitude && data.Latitude <= 90 && -180 <= data.Longitude && data.Longitude <= 180
	if valid == false {
		return xerr.New(fmt.Sprintf("xgeo.WGS84{} invalid format"))
	}
	return
}
func (data WGS84) GCJ02() GCJ02 {
	lng, lat := coord.WGS84toGCJ02(data.Longitude, data.Latitude)
	return GCJ02{
		Longitude: lng,
		Latitude:  lat,
	}
}
func (data WGS84) BD09() BD09 {
	lng, lat := coord.WGS84toBD09(data.Longitude, data.Latitude)
	return BD09{
		Longitude: lng,
		Latitude:  lat,
	}
}

type GCJ02 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude  float64 `json:"latitude" note:"纬度"`
}

func (data GCJ02) Validator() (err error) {
	valid := -90 <= data.Latitude && data.Latitude <= 90 && -180 <= data.Longitude && data.Longitude <= 180
	if valid == false {
		return xerr.New(fmt.Sprintf("xgeo.GCJ02{} invalid format"))
	}
	return
}

// LatCommaLngString
// 返回 "纬度,经度" 格式字符串
// 可能所有人都至少一次踩过这个坑：地理坐标点用字符串形式表示时是纬度在前，经度在后（ "latitude,longitude" ），
// 而数组形式表示时是经度在前，纬度在后（ [longitude,latitude] ）—顺序刚好相反。
func (data GCJ02) LatCommaLngString() (latCommaLng string) {
	return xconv.Float64String(data.Latitude) + "," + xconv.Float64String(data.Longitude)
}
func (data GCJ02) WGS84() WGS84 {
	lng, lat := coord.GCJ02toWGS84(data.Longitude, data.Latitude)
	return WGS84{
		Longitude: lng,
		Latitude:  lat,
	}
}
func (data GCJ02) BD09() BD09 {
	lng, lat := coord.GCJ02toBD09(data.Longitude, data.Latitude)
	return BD09{
		Longitude: lng,
		Latitude:  lat,
	}
}

type BD09 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude  float64 `json:"latitude" note:"纬度"`
}

func (data BD09) Validator() (err error) {
	valid := -90 <= data.Latitude && data.Latitude <= 90 && -180 <= data.Longitude && data.Longitude <= 180
	if valid == false {
		return xerr.New(fmt.Sprintf("xgeo.BD09{} invalid format"))
	}
	return
}
func (data BD09) WGS84() WGS84 {
	lng, lat := coord.BD09toWGS84(data.Longitude, data.Latitude)
	return WGS84{
		Longitude: lng,
		Latitude:  lat,
	}
}
func (data BD09) GCJ02() GCJ02 {
	lng, lat := coord.BD09toGCJ02(data.Longitude, data.Latitude)
	return GCJ02{
		Longitude: lng,
		Latitude:  lat,
	}
}
