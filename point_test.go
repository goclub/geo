package xgeo

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
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

func TestDistanceInMeters(t *testing.T) {
	p1 := Point{Type: "Point", Coordinates: [2]float64{116.3975, 39.9088}} // 北京市中心
	p2 := Point{Type: "Point", Coordinates: [2]float64{121.4737, 31.2304}} // 上海市中心
	expectedDistance := 1068150.25

	distance := p1.DistanceInMeters(p2)
	log.Print(math.Abs(distance - expectedDistance))
	if math.Abs(distance-expectedDistance) > 2 {
		t.Errorf("DistanceInMeters() = %v, want %v", distance, expectedDistance)
	}
}
