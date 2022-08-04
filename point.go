package xgeo

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	xconv "github.com/goclub/conv"
	xerr "github.com/goclub/error"
	"github.com/goclub/geo/internal/coord"
)

// PointJSON GeoJSON  支持 mongo bson
// sql 请使用 xgeo.WGS84{}
// GeoJson规范（RFC 7946）全文翻译: https://zhuanlan.zhihu.com/p/141554586
// xgeo.NewPoint(xgeo.WGS84{121.48294,31.2328}) // WGS84{经度,纬度}
type PointJSON struct {
	Type pointType `json:"type" bson:"type"`
	// []float64{longitude, latitude} []float64{经度, 纬度}
	// 可能所有人都至少一次踩过这个坑：地理坐标点用字符串形式表示时是纬度在前，经度在后（ "latitude,longitude" ），
	// 而数组形式表示时是经度在前，纬度在后（ [longitude,latitude] ）—顺序刚好相反。
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

// 用内部类型来强制调用者使用 NewPoint() 来创造 PointJSON
type pointType *string

func NewPointJSON(data WGS84) PointJSON {
	typevalue := "Point"
	return PointJSON{
		&typevalue,
		[]float64{data.Longitude, data.Latitude},
	}
}

func (p PointJSON) WGS84() WGS84 {
	return WGS84{
		Longitude: p.Coordinates[0],
		Latitude:  p.Coordinates[1],
	}
}

// WGS84 实现了 sql 接口 driver.Valuer sql.Scanner
type WGS84 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude  float64 `json:"latitude" note:"纬度"`
}

// String returns the WKT (Well Known Text) representation of the point.
func (data WGS84) String() string {
	return fmt.Sprintf("POINT(%v %v)", data.Longitude, data.Latitude)
}

// Scan implements the SQL driver.Scanner interface and will scan the
// MySQL POINT(x y) into the PointJSON struct.
func (data *WGS84) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return xerr.New(fmt.Sprintf("goclub/geo: cannot scan type into bytes: %T", b))
	}

	// MySQL bug, it returns the internal representation with 4 zero bytes before
	// the value: https://bugs.mysql.com/bug.php?id=69798
	b = b[4:]

	r := bytes.NewReader(b)

	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return xerr.WithStack(err)
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return xerr.New(fmt.Sprintf("goclub/geo: invalid byte order %v", wkbByteOrder))
	}

	var wkbGeometryType uint32
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return xerr.WithStack(err)
	}

	if wkbGeometryType != 1 {
		return xerr.New(fmt.Sprintf("goclub/geo: unexpected geometry type: wanted 1 (point), got %d", wkbGeometryType))
	}

	if err := binary.Read(r, byteOrder, data); err != nil {
		return xerr.WithStack(err)
	}

	return nil
}

// Value implements the SQL driver.Valuer interface and will return the string
// representation of the WGS84 struct by calling the String() method
func (data WGS84) Value() (driver.Value, error) {
	w := bytes.NewBuffer(nil)

	// MySQL bug, it returns the internal representation with 4 zero bytes before
	// the value: https://bugs.mysql.com/bug.php?id=69798
	w.Write([]byte{0, 0, 0, 0})

	var wkbByteOrder uint8 = 1
	if err := binary.Write(w, binary.LittleEndian, wkbByteOrder); err != nil {
		return nil, xerr.WithStack(err)
	}

	var wkbGeometryType uint32 = 1
	if err := binary.Write(w, binary.LittleEndian, wkbGeometryType); err != nil {
		return nil, xerr.WithStack(err)
	}

	if err := binary.Write(w, binary.LittleEndian, data); err != nil {
		return nil, xerr.WithStack(err)
	}

	return w.Bytes(), nil
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
