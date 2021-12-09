package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable to hold db connection, check other better ways for connecting
var DB *sql.DB

func connectDB() {
	db, err := sql.Open("mysql", "root:password@tcp(mysqlhost:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	//Create users table with id and name fields
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id int unsigned NOT NULL AUTO_INCREMENT, name varchar(100) NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB;")

	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Table Successfully Created\n")
	DB = db
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	_, err := DB.Query("INSERT INTO users(name) VALUES('Username ')")
	if err != nil {
		fmt.Fprintf(w, "Error occured while inserting the record. ", err.Error())
		return
	}
	fmt.Fprintf(w, "New record inserted successfully")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	results, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		fmt.Fprintf(w, "Sone error occured while reading from the db")
		return
	}

	var id int
	var name string

	for results.Next() {
		err = results.Scan(&id, &name)
		if err != nil {
			fmt.Fprintf(w, "Sone error occured while looping over the result set")
			return
		}
		outputStr := strconv.Itoa(id) + " - " + name + " <br> \n"
		fmt.Fprintf(w, outputStr)
	}

}

func main() {
	connectDB()

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/list", listHandler)

	fmt.Printf("Starting server at port 80 in Docker\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
