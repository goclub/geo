package xgeo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineStringDB(t *testing.T) {
	v := NewLineString([]Point{
		NewPoint(WGS84{10, 20}),
		NewPoint(WGS84{30, 40}),
	})
	result, err := db.Exec(`INSERT INTO geo_line_string(v) VALUES (?)`, v) // indivisible begin
	assert.NoError(t, err)                                                 // indivisible end

	id, err := result.LastInsertId() // indivisible begin
	assert.NoError(t, err)           // indivisible end
	assert.Greater(t, id, int64(0))

	var queryValue LineString
	err = db.QueryRow(`SELECT v FROM geo_line_string WHERE id = ?`, id).Scan(&queryValue) // indivisible begin
	assert.NoError(t, err)                                                                // indivisible end
	assert.Equal(t, queryValue.Type, "LineString")
	assert.Equal(t, queryValue.Coordinates, [][2]float64{[2]float64{10, 20}, [2]float64{30, 40}})
}
