package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable to hold db connection, check other better ways for connecting
var DB *sql.DB

func connectDB() {
	db, err := sql.Open("mysql", "root:password@tcp(database-service:3306)/testdb")

	if err != nil {
		fmt.Errorf("Error : %v", err.Error())
	}

	// http://www.matthiassommer.it/programming/docker-compose-retry-database-connect-with-docker-and-go/
	// below retrial steps are followed by referring above blog

	retryCount := 10
	for {
		err := db.Ping()
		if err != nil {
			if retryCount == 0 {
				log.Fatalf("Not able to establish connection to database")
				return
			}

			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	log.Println("Connection to MySQL has created successfully")

	//Create users table with id and name fields
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id int unsigned NOT NULL AUTO_INCREMENT, name varchar(100) NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB;")

	if err != nil {
		fmt.Errorf("Error : %v", err.Error())
		return
	}

	log.Println("Table Successfully Created")
	DB = db
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	_, err := DB.Query("INSERT INTO users(name) VALUES('Username ')")

	if err != nil {
		fmt.Fprintf(w, "Error occured while inserting the record. ", err.Error())
		return
	}
	log.Println("New record add to users table")
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
		outputStr := strconv.Itoa(id) + " - " + name + " \n"
		fmt.Fprintf(w, outputStr)
	}
	log.Println("Listing service called")
}

func main() {
	connectDB()
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/list", listHandler)

	log.Println("Starting server at port 80 inside Docker")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
