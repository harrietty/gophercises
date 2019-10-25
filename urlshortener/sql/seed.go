package seed

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func SeedDB() {
	db, err := sql.Open("sqlite3", "./urlmap.db")
	if err != nil {
		fmt.Printf("Error opening sqlite3: %v\n", err)
	}

	statement, err := db.Prepare("DROP TABLE IF EXISTS paths")
	statement.Exec()
	handleErr(err)
	if err == nil {
		fmt.Println("Dropped existing table paths")
	}
	
	statement, err = db.Prepare("CREATE TABLE paths (id integer primary key autoincrement, path varchar(255), url varchar(500))")
	statement.Exec()
	handleErr(err)
	if err == nil {
		fmt.Println("Created table paths")
	}

	statement, err = db.Prepare("insert into paths (path, url) values (?, ?), (?, ?)")
	_, err = statement.Exec("/mice", "https://en.wikipedia.org/wiki/Mouse", "/bats", "https://www.bats.org.uk/about-bats")
	handleErr(err)

	fmt.Println("Seeded the database, closing connection")
	db.Close()
}

func handleErr(err interface{}) {
	if err != nil {
		fmt.Printf("Error seeding DB: %v\n", err)
	}
}