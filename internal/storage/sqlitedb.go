package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/util"
)

// InitDB initializes the SQLite database.
func InitDB() *sql.DB {

	db, err := sql.Open("sqlite3", "C:\\Users\\Franco\\go\\src\\localdev\\HrHelper\\hrdata.db")
	util.Panic(err)

	createConfigTableSQL := `CREATE TABLE IF NOT EXISTS Config (
        "username" TEXT,
        "password" TEXT,
        "remember" BOOLEAN,
        "baseUrl" TEXT,
        "unicodeOutput" BOOLEAN
    );`
	_, err = db.Exec(createConfigTableSQL)
	util.Panic(err)

	selectConfig := `SELECT username, password,remember, baseUrl, unicodeOutput FROM Config;`
	rows, err := db.Query(selectConfig)
	util.Panic(err)
	defer rows.Close()

	var configs []config.Config
	for rows.Next() {
		var config config.Config
		if err := rows.Scan(&config.Username, &config.Password, &config.Remember, &config.BaseUrl, &config.UnicodeOutput); err != nil {
			util.Panic(err)
		}
		configs = append(configs, config)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		util.Panic(err)
	}

	// Check if only one row was returned
	if len(configs) > 1 {
		errorMsg := fmt.Sprintf("Expected 1 row in Config table, got %d", len(configs))
		panic(errorMsg)
	}

	if len(configs) == 0 {
		println("No config found. Inserting default config...")
		insertDefault := `INSERT INTO Config (username, password, remember, baseUrl, unicodeOutput)
                          VALUES (?, ?, ?, ?, ?);`
		_, err := db.Exec(insertDefault, "", "", false, "https://www.hogwartsrol.com/", true)
		util.Panic(err)
	}

	return db
}
func GetConfig(db *sql.DB) *config.Config {
	query := `SELECT username, password, remember, baseUrl, unicodeOutput FROM Config LIMIT 1;`

	var config config.Config

	// Execute the query
	row := db.QueryRow(query)
	err := row.Scan(&config.Username, &config.Password, &config.Remember, &config.BaseUrl, &config.UnicodeOutput)
	util.Panic(err)

	return &config
}

func UpdateConfig(db *sql.DB, config *config.Config) {
	// First, ensure that there is a row to update
	ensureRowSQL := `INSERT INTO Config (username, password, remember, baseUrl, unicodeOutput)
                    SELECT '', '',false, '', false WHERE NOT EXISTS (SELECT 1 FROM Config);`
	_, err := db.Exec(ensureRowSQL)
	util.Panic(err)

	// Now, update the existing row with new values
	updateSQL := `UPDATE Config SET username = ?, password = ?, remember = ?,baseUrl = ?, unicodeOutput = ?;`
	_, err = db.Exec(updateSQL, config.Username, config.Password, config.Remember, config.BaseUrl, config.UnicodeOutput)
	util.Panic(err)
}
