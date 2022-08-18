package xgeo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolygonDB(t *testing.T) {

	ring1 := NewLineString([]Point{
		NewPoint(WGS84{0, 0}),
		NewPoint(WGS84{10, 10}),
		NewPoint(WGS84{20, 20}),
		NewPoint(WGS84{0, 0}),
	})
	v := NewPolygon([]LineString{ring1})
	result, err := db.Exec(`INSERT INTO geo_polygon(v) VALUES (?)`, v) // indivisible begin
	assert.NoError(t, err)                                             // indivisible end

	id, err := result.LastInsertId() // indivisible begin
	assert.NoError(t, err)           // indivisible end
	assert.Greater(t, id, int64(0))

	var queryValue Polygon
	err = db.QueryRow(`SELECT v FROM geo_polygon WHERE id = ?`, id).Scan(&queryValue) // indivisible begin
	assert.NoError(t, err)                                                            // indivisible end
	assert.Equal(t, queryValue.Type, "Polygon")
	assert.Equal(t, queryValue.Coordinates, [][][2]float64([][][2]float64{[][2]float64{[2]float64{0, 0}, [2]float64{10, 10}, [2]float64{20, 20}, [2]float64{0, 0}}}))
}
