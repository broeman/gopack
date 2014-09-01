// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Database handling, Initialization of Database
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"           // Postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"os"
	"os/user"
)

// placeholder, should be put as a setting
const PATH = "GoPack"

// Helper function to find the correct path to $HOME/.config/PATH
func getLibraryDir(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path = usr.HomeDir + "/.config/" + path

	_, err = os.Stat(path)
	if err != nil {
		fmt.Println("gopack config not found")
	}
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	return path + "/packages.db"
}

// Initialize the Database
func InitDB() {
	var db *sql.DB
	var err error

	_, err = os.Stat(getLibraryDir(PATH))
	if err == nil {
		fmt.Println("packages.db already exist")
		return
	}

	db, err = sql.Open("sqlite3", getLibraryDir(PATH))
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

	err = sampleData(db)
	if err != nil {
		fmt.Printf("sample data error: %v\n", err)
		return
	}

}

// Initializing the database
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

	return nil
}
