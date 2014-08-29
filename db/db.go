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
	defer stmt.Close()

	stmt, err = db.Prepare("INSERT INTO packages(id, name, versionregex, installed) VALUES($1, $2, $3, $4)")
	if err != nil {
		fmt.Printf("stmt.Prepare error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(0, "Super", "$", true)
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}

	_, err = stmt.Exec(1, "Test", "$", false)
	if err != nil {
		fmt.Printf("stmt.Exec error: %v\n", err)
		return err
	}

	return nil
}

func QueryPackage(pack string) (rows *sql.Rows, err error) {
	var db *sql.DB

	db, err = sql.Open("sqlite3", "./packages.db")
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

	db, err = sql.Open("sqlite3", "./packages.db")
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

	db, err = sql.Open("sqlite3", "./packages.db")
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
