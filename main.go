package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Timings struct {
	Fajr    string `json:"Fajr"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type Data struct {
	Timings Timings `json:"timings"`
}

type Response struct {
	Data Data `json:"data"`
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/prayertime")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	url := "https://api.aladhan.com/v1/timingsByCity?city=Dhaka&country=Bangladesh"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var r Response
	json.Unmarshal(body, &r)

	// fajr := r.Data.Timings.Fajr
	// dhuhr := r.Data.Timings.Dhuhr
	// asr := r.Data.Timings.Asr
	// maghrib := r.Data.Timings.Maghrib
	// isha := r.Data.Timings.Isha

	// stmt, err := db.Prepare("INSERT INTO timings(fajr, dhuhr, asr, maghrib, isha) VALUES(?, ?, ?, ?, ?)")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(fajr, dhuhr, asr, maghrib, isha)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	
	/* commented codes are used to insert value on mysql table.
	I commeneted because there will be double insertion while running the code*/

	fmt.Println("Record inserted successfully")

	fmt.Println("Fajr:", r.Data.Timings.Fajr)
	fmt.Println("Dhuhr:", r.Data.Timings.Dhuhr)
	fmt.Println("Asr:", r.Data.Timings.Asr)
	fmt.Println("Maghrib:", r.Data.Timings.Maghrib)
	fmt.Println("Isha:", r.Data.Timings.Isha)
}
