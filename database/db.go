package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to the database
func connectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/prayertime")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Handler function to handle HTTP requests and return the
// output of the MySQL query as the response
func handler(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db := connectToDatabase()
	defer db.Close()

	// Select the rows you want to convert and decode as JSON
	rows, err := db.Query("SELECT id, fajr, dhuhr, asr, maghrib, isha FROM timings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to hold the rows
	var values []map[string]interface{}

	// Iterate over the selected rows
	for rows.Next() {
		// Declare variables to hold the selected values
		var id int
		var fajr, dhuhr, asr, maghrib, isha string

		// Scan the selected values into the variables
		err = rows.Scan(&id, &fajr, &dhuhr, &asr, &maghrib, &isha)
		if err != nil {
			log.Fatal(err)
		}

		// Convert the values to the desired type (e.g. int64)
		// and add them to the slice
		row := map[string]interface{}{
			"id":      id,
			"fajr":    fajr,
			"dhuhr":   dhuhr,
			"asr":     asr,
			"maghrib": maghrib,
			"isha":    isha,
		}
		values = append(values, row)
	}

	// Encode the slice as JSON
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	// Set the content type and write the JSON data to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	// Create a new HTTP server and set the handler function
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
	// Start the server
	server.ListenAndServe()
}
