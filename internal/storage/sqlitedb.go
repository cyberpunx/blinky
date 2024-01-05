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

	db, err := sql.Open("sqlite3", "hrdata.db")
	util.Panic(err)

	createConfigTableSQL := `CREATE TABLE IF NOT EXISTS Config (
        "username" TEXT,
        "password" TEXT,
        "remember" BOOLEAN,
        "baseUrl" TEXT,
        "unicodeOutput" BOOLEAN,
        "gSheetTokenFile" TEXT,
        "gSheetCredFile" TEXT,
        "gSheetId" TEXT
    );`
	_, err = db.Exec(createConfigTableSQL)
	util.Panic(err)

	selectConfig := `SELECT * FROM Config;`
	rows, err := db.Query(selectConfig)
	util.Panic(err)
	defer rows.Close()

	var configs []config.Config
	for rows.Next() {
		var config config.Config
		if err := rows.Scan(
			&config.Username,
			&config.Password,
			&config.Remember,
			&config.BaseUrl,
			&config.UnicodeOutput,
			&config.GSheetTokenFile,
			&config.GSheetCredFile,
			&config.GSheetId); err != nil {
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
		insertDefault := `INSERT INTO Config (
                    username, 
                    password, 
                    remember, 
                    baseUrl, 
                    unicodeOutput,
                    gSheetTokenFile,
                    gSheetCredFile,
                    gSheetId)
                          VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
		_, err := db.Exec(insertDefault, "", "", false, "https://www.hogwartsrol.com/", true, "token.json", "client_secret.json", "13CCYZ4veljB6ItPNHdvxvClBZJaC1w-QMkq-H5btR74")
		util.Panic(err)
	}

	createPotionSubTableSQL := `CREATE TABLE IF NOT EXISTS PotionSubforumConfig (
        "url" TEXT PRIMARY KEY,
        "timeLimit" INTEGER NOT NULL,
        "turnLimit" INTEGER NOT NULL
    );`
	_, err = db.Exec(createPotionSubTableSQL)
	util.Panic(err)

	// Select all rows from PotionSubforumConfig
	selectPotionSub := `SELECT * FROM PotionSubforumConfig;`
	rows, err = db.Query(selectPotionSub)
	util.Panic(err)
	defer rows.Close()

	var potionSubs []config.PotionSubforumConfig
	for rows.Next() {
		var potionSub config.PotionSubforumConfig
		if err := rows.Scan(
			&potionSub.Url,
			&potionSub.TimeLimit,
			&potionSub.TurnLimit); err != nil {
			util.Panic(err)
		}
		potionSubs = append(potionSubs, potionSub)
	}
	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		util.Panic(err)
	}

	if len(potionSubs) == 0 {
		println("No PotionSubforumConfig found. Inserting default config...")
		insertDefault := `INSERT INTO PotionSubforumConfig (
					url, 
					timeLimit, 
					turnLimit)
						  VALUES (?, ?, ?);`
		_, err := db.Exec(insertDefault, "f98-club-de-pociones", 72, 8)
		util.Panic(err)
	}

	createPotionThrTableSQL := `CREATE TABLE IF NOT EXISTS PotionThreadConfig (
        "url" TEXT PRIMARY KEY,
        "timeLimit" INTEGER NOT NULL,
        "turnLimit" INTEGER NOT NULL
    );`
	_, err = db.Exec(createPotionThrTableSQL)
	util.Panic(err)

	return db
}
func GetConfig(db *sql.DB) *config.Config {
	query := `SELECT * FROM Config LIMIT 1;`

	var config config.Config

	// Execute the query
	row := db.QueryRow(query)
	err := row.Scan(
		&config.Username,
		&config.Password,
		&config.Remember,
		&config.BaseUrl,
		&config.UnicodeOutput,
		&config.GSheetTokenFile,
		&config.GSheetCredFile,
		&config.GSheetId)
	util.Panic(err)

	return &config
}

func UpdateConfig(db *sql.DB, config *config.Config) {
	// First, ensure that there is a row to update
	ensureRowSQL := `INSERT INTO Config (
                    username, 
                    password, 
                    remember, 
                    baseUrl, 
                    unicodeOutput, 
                    gSheetTokenFile, 
                    gSheetCredFile, 
                    gSheetId)
                    SELECT '', '',false, '', false, '', '', '' WHERE NOT EXISTS (SELECT 1 FROM Config);`
	_, err := db.Exec(ensureRowSQL)
	util.Panic(err)

	// Now, update the existing row with new values
	updateSQL := `UPDATE Config SET 
                  username = ?, 
                  password = ?, 
                  remember = ?,
                  baseUrl = ?, 
                  unicodeOutput = ?, 
                  gSheetTokenFile = ?, 
                  gSheetCredFile = ?, 
                  gSheetId = ?;`
	_, err = db.Exec(updateSQL,
		config.Username,
		config.Password,
		config.Remember,
		config.BaseUrl,
		config.UnicodeOutput,
		config.GSheetTokenFile,
		config.GSheetCredFile,
		config.GSheetId)
	util.Panic(err)
}

func GetPotionSubforum(db *sql.DB) *[]config.PotionSubforumConfig {
	query := `SELECT * FROM PotionSubforumConfig;`

	var potionSubConfig []config.PotionSubforumConfig

	rows, err := db.Query(query)
	util.Panic(err)

	//loop all rows
	for rows.Next() {
		var potionSubforum config.PotionSubforumConfig
		err := rows.Scan(&potionSubforum.Url, &potionSubforum.TimeLimit, &potionSubforum.TurnLimit)
		util.Panic(err)
		potionSubConfig = append(potionSubConfig, potionSubforum)
	}
	return &potionSubConfig
}

func UpdatePotionSubforum(db *sql.DB, potionSubConfig *[]config.PotionSubforumConfig) {
	// Truncate the table and insert the new values
	truncateSQL := `DELETE FROM PotionSubforumConfig;`
	_, err := db.Exec(truncateSQL)
	util.Panic(err)

	//insert one by one the potionSubConfig
	for _, potionSubforum := range *potionSubConfig {
		insertSQL := `INSERT INTO PotionSubforumConfig (url, timeLimit, turnLimit)
					SELECT ?, ?, ?;`
		_, err := db.Exec(insertSQL, potionSubforum.Url, potionSubforum.TimeLimit, potionSubforum.TurnLimit)
		util.Panic(err)
	}
}

func GetPotionThread(db *sql.DB) *[]config.PotionThreadConfig {
	query := `SELECT * FROM PotionThreadConfig;`

	var potionThrConfig []config.PotionThreadConfig

	rows, err := db.Query(query)
	util.Panic(err)

	//loop all rows
	for rows.Next() {
		var potionThread config.PotionThreadConfig
		err := rows.Scan(&potionThread.Url, &potionThread.TimeLimit, &potionThread.TurnLimit)
		util.Panic(err)
		potionThrConfig = append(potionThrConfig, potionThread)
	}
	return &potionThrConfig
}

func UpdatePotionThread(db *sql.DB, potionThrConfig *[]config.PotionThreadConfig) {
	// Truncate the table and insert the new values
	truncateSQL := `DELETE FROM PotionThreadConfig;`
	_, err := db.Exec(truncateSQL)
	util.Panic(err)

	//insert one by one the potionThrConfig
	for _, potionThread := range *potionThrConfig {
		insertSQL := `INSERT INTO PotionThreadConfig (url, timeLimit, turnLimit)
					SELECT ?, ?, ?;`
		_, err := db.Exec(insertSQL, potionThread.Url, potionThread.TimeLimit, potionThread.TurnLimit)
		util.Panic(err)
	}
}
