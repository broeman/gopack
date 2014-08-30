package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"           // Postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func InitDB() {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./packages.db")
	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return
	}
	defer db.Close()

	err = doInitialize(db)
	if err != nil {
		fmt.Printf("initialize error: %v\n", err)
		return
	}
}

func doInitialize(db *sql.DB) error {
	var stmt *sql.Stmt
	var err error

	// create the packages table
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS packages(id SERIAL PRIMARY KEY, name VARCHAR(255), versionregex VARCHAR(100), installed BOOL);")
	if err != nil {
		fmt.Printf("db.Prepare initializing error: %v\n", err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}

	// create the versions table
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS versions(id SERIAL PRIMARY KEY, package_id INT, version VARCHAR(50), description VARCHAR(255), state VARCHAR(40));")
	if err != nil {
		fmt.Printf("db.Prepare initializing error: %v\n", err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}

	// create the depencies table
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS dependencies(id SERIAL PRIMARY KEY, version_id INT, package VARCHAR(255), version VARCHAR(255));")
	if err != nil {
		fmt.Printf("db.Prepare initializing error: %v\n", err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}
	defer stmt.Close()

	// inserting sample data
	stmt, err = db.Prepare("INSERT INTO packages(id, name, versionregex, installed) VALUES($1, $2, $3, $4)")
	if err != nil {
		fmt.Printf("stmt.Prepare error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(0, "Super", "", true)
	if err != nil {
		fmt.Printf("Make super package error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(1, "Test", "", false)
	if err != nil {
		fmt.Printf("Make test package error: %v\n", err)
		return err
	}

	stmt, err = db.Prepare("INSERT INTO versions(id, package_id, version, description, state) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Printf("stmt.Prepare error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(0, 0, "0.2-beta3", "Super Package 0.2-beta", "testing")
	if err != nil {
		fmt.Printf("Make Super Version 0.2b error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(1, 0, "0.1", "Super Package 0.1", "stable")
	if err != nil {
		fmt.Printf("Make Super Version 0.1 error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(2, 1, "0.1", "Testing 0.1", "deprecated")
	if err != nil {
		fmt.Printf("Make Test Version 0.1 error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(3, 1, "0.5", "Testing 0.5", "stable")
	if err != nil {
		fmt.Printf("Make Test Version 0.5 error: %v\n", err)
		return err
	}

	stmt, err = db.Prepare("INSERT INTO dependencies(id, version_id, package, version) VALUES($1, $2, $3, $4)")
	if err != nil {
		fmt.Printf("stmt.Prepare error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(0, 0, "Test", "0.5")
	if err != nil {
		fmt.Printf("Make Super Depencency error: %v\n", err)
		return err
	}

	return nil
}
