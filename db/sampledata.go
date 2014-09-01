// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Database handling, sample data
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"           // Postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// sampledata for testing
func sampleData(db *sql.DB) error {
	var stmt *sql.Stmt
	var err error

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
	defer stmt.Close()

	return nil
}
