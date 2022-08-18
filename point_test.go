package xgeo

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
)

func TestPoint(t *testing.T) {
	func() struct{} {
		// -------------
		var err error
		_ = err
		ctx := context.Background()
		_ = ctx

		p := NewPoint(WGS84{Latitude: 12.34, Longitude: 56.78})
		assert.Equal(t, p.String(), "POINT(56.780000 12.340000)")
		data, err := json.Marshal(p) // indivisible begin
		assert.NoError(t, err)       // indivisible end
		assert.Equal(t, string(data), `{"type":"Point","coordinates":[56.78,12.34]}`)

		// -------------
		return struct{}{}
	}()
}
func TestPointDB(t *testing.T) {
	result, err := db.Exec(`INSERT INTO geo_point(v) VALUES (?)`, NewPoint(WGS84{Latitude: 12.34, Longitude: 56.78})) // indivisible begin
	assert.NoError(t, err)                                                                                            // indivisible end

	id, err := result.LastInsertId() // indivisible begin
	assert.NoError(t, err)           // indivisible end
	assert.Greater(t, id, int64(0))

	var other Point
	err = db.QueryRow(`SELECT v FROM geo_point WHERE id = ?`, id).Scan(&other) // indivisible begin
	assert.NoError(t, err)                                                     // indivisible end

	require.EqualValues(t, other.WGS84().Latitude, 12.34)
	require.EqualValues(t, other.WGS84().Longitude, 56.78)
	assert.NoError(t, other.WGS84().Validator())
}
