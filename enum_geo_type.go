package xgeo

// GeoType Generate by https://goclub.run/?k=enum
// ---------------------- DO NOT EDIT (Begin) ----------------------
import (
	"fmt"
	xerr "github.com/goclub/error"
)

// GeoType
// Source enums:
// {"name":"GeoType","type":"string","items":[{"field":"Point","value":"Point","tailed":", ","label":"点"},{"field":"LineString","value":"LineString","label":"线","tailed":", "},{"field":"Polygon","value":"Polygon","label":"面","tailed":", "},{"field":"MultiPoint","value":"MultiPoint","label":"多点","tailed":", "},{"field":"MultiLineString","value":"MultiLineString","label":"多线","tailed":", "},{"field":"MultiPolygon","value":"MultiPolygon","label":"多面"}]}
type GeoType string

// NewGeoType Create GeoType by string
func NewGeoType(v string) (geoType GeoType, err error) {
	geoType = GeoType(v)
	err = geoType.Validator()
	return
}

// String return GeoType basic types
func (v GeoType) String() string { return string(v) }

// EnumGeoType Example: if geoType == EnumGeoType().xxx {...}
func EnumGeoType() (e struct {
	Point           GeoType
	LineString      GeoType
	Polygon         GeoType
	MultiPoint      GeoType
	MultiLineString GeoType
	MultiPolygon    GeoType
}) {
	e.Point = "Point"
	e.LineString = "LineString"
	e.Polygon = "Polygon"
	e.MultiPoint = "MultiPoint"
	e.MultiLineString = "MultiLineString"
	e.MultiPolygon = "MultiPolygon"
	return
}
func EnumGeoTypeExampleSwitch() {
	_ = `
switch point, lineString, polygon, multiPoint, multiLineString, multiPolygon := m.EnumGeoTypeSwitch(); v {
case point.Point:
    // @TODO write some code
case lineString.LineString:
    // @TODO write some code
case polygon.Polygon:
    // @TODO write some code
case multiPoint.MultiPoint:
    // @TODO write some code
case multiLineString.MultiLineString:
    // @TODO write some code
case multiPolygon.MultiPolygon:
    // @TODO write some code
default:
    err = xerr.New(fmt.Sprintf("GeoType can not be %v", v))
    return
}
`
}

// EnumGeoTypeSwitch safe switch of all values
// example: m.EnumGeoTypeExampleSwitch()
func EnumGeoTypeSwitch() (
	point struct{ Point GeoType },
	lineString struct{ LineString GeoType },
	polygon struct{ Polygon GeoType },
	multiPoint struct{ MultiPoint GeoType },
	multiLineString struct{ MultiLineString GeoType },
	multiPolygon struct{ MultiPolygon GeoType },
) {
	e := EnumGeoType()
	point.Point = e.Point
	lineString.LineString = e.LineString
	polygon.Polygon = e.Polygon
	multiPoint.MultiPoint = e.MultiPoint
	multiLineString.MultiLineString = e.MultiLineString
	multiPolygon.MultiPolygon = e.MultiPolygon
	return
}

// Validator Verify data
func (v GeoType) Validator(custom ...error) error {
	outError := xerr.New(fmt.Sprintf("GeoType can not be %v", v))
	if len(custom) != 0 {
		outError = custom[0]
	}
	switch point, lineString, polygon, multiPoint, multiLineString, multiPolygon := EnumGeoTypeSwitch(); v {
	case point.Point:
	case lineString.LineString:
	case polygon.Polygon:
	case multiPoint.MultiPoint:
	case multiLineString.MultiLineString:
	case multiPolygon.MultiPolygon:
	default:
		return outError
	}
	return nil
}

// IsZero
func (v GeoType) IsZero() bool {
	return v == ""
}

// JavaScript code for https://github.com/2type/admin#_enum
/*
TA.enum.geoType = [
    {
        key: "Point",
        value: "Point",
        label: "点",
    },
    {
        key: "LineString",
        value: "LineString",
        label: "线",
    },
    {
        key: "Polygon",
        value: "Polygon",
        label: "面",
    },
    {
        key: "MultiPoint",
        value: "MultiPoint",
        label: "多点",
    },
    {
        key: "MultiLineString",
        value: "MultiLineString",
        label: "多线",
    },
    {
        key: "MultiPolygon",
        value: "MultiPolygon",
        label: "多面",
    },
]
*/
// ---------------------- DO NOT EDIT (End) ----------------------
