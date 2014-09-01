// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Database handling, package related
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"           // Postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func QueryPackage(pack string) (rows *sql.Rows, err error) {
	var db *sql.DB

	db, err = sql.Open("sqlite3", getLibraryDir(PATH))
	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return nil, err
	}
	defer db.Close()

	rows, err = db.Query("SELECT id, name, versionregex, installed FROM packages WHERE name=?", pack)
	if err != nil {
		fmt.Printf("sql.Query error: %v\n", err)
	}

	return rows, nil
}

func UpdatePackageInstalled(pack string, update bool) (err error) {
	var db *sql.DB

	db, err = sql.Open("sqlite3", getLibraryDir(PATH))
	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE packages SET installed=$1 WHERE name=$2")
	if err != nil {
		fmt.Printf("stmt.Prepare error: %v\n", err)
		return err
	}
	_, err = stmt.Exec(update, pack)
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}
	return nil
}

// placeholder for showing all packages, not really production ready ;)
func QueryAllPackages() (rows *sql.Rows, err error) {
	var db *sql.DB

	db, err = sql.Open("sqlite3", getLibraryDir(PATH))
	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return nil, err
	}
	defer db.Close()

	rows, err = db.Query("SELECT id, name, versionregex, installed FROM packages")
	if err != nil {
		fmt.Printf("sql.Query error: %v\n", err)
	}

	return rows, nil
}
