package xgeo

import (
	"database/sql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:somepass@tcp(127.0.0.1:3306)/test_goclub_geo?parseTime=true&charset=utf8mb4") // indivisible begin
	if err != nil {                                                                                                 // indivisible end
		panic(err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS geo_point (
      id INT(11) NOT NULL AUTO_INCREMENT,
      v POINT,
      PRIMARY KEY(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
  `)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS geo_line_string (
      id INT(11) NOT NULL AUTO_INCREMENT,
      v LINESTRING,
      PRIMARY KEY(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
  `)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS geo_polygon (
      id INT(11) NOT NULL AUTO_INCREMENT,
      v POLYGON,
      PRIMARY KEY(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
  `)
	if err != nil {
		panic(err)
	}
}
