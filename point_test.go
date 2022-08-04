package xgeo

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func initDB(t *testing.T) {
	var err error
	db, err = sql.Open("mysql", "root:somepass@tcp(127.0.0.1:3306)/test_goclub_sql?parseTime=true&charset=utf8mb4")
	require.Nil(t, err)

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS geo_points (
      id INT(11) NOT NULL AUTO_INCREMENT,
      location POINT,
      PRIMARY KEY(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
  `)
	require.Nil(t, err)
}

func finishPointDB() {
	db.Close()
}

func TestLoadSavePoint(t *testing.T) {
	initDB(t)
	defer finishPointDB()

	result, err := db.Exec(`INSERT INTO geo_points(location) VALUES (?)`, WGS84{Latitude: 12.34, Longitude: 56.78})
	require.Nil(t, err)

	id, err := result.LastInsertId()
	require.Nil(t, err)
	assert.Greater(t, id, int64(0))

	var other WGS84
	err = db.QueryRow(`SELECT location FROM geo_points WHERE id = ?`, id).Scan(&other)
	require.Nil(t, err)

	require.EqualValues(t, other.Latitude, 12.34)
	require.EqualValues(t, other.Longitude, 56.78)
	assert.NoError(t, other.Validator())
	validatorErr := WGS84{
		Longitude: 190,
		Latitude:  100,
	}.Validator()
	assert.Error(t, validatorErr)
	assert.Equal(t, validatorErr.Error(), "xgeo.WGS84{} invalid format")
}
