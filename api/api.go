package main

import (
	"fmt"
	"os"
	"database/sql"
	// "log"
	// "io/ioutil"
)

const userKey string = "XXXXXXXXXXX"

type cityStation struct {
	City       string
	Station    string
}

var cityInfo = [5]cityStation{
	cityStation{"London", "03772"},
	cityStation{"Cambridge", "03571"},
	cityStation{"Nottingham", "03354"},
	cityStation{"Manchester", "03334"},
	cityStation{"Leeds", "EGNM0"},
}

func main() {
	fmt.Println(os.Getenv("GO15VENDOREXPERIMENT"));
	//startDate := `2019-05-25`
	//endDate := `2019-05-31`

	//website := fmt.Sprintf(`https://api.meteostat.net/v1/history/daily?station=%s&start=%s&end=%s&key=%s`, cityInfo[0].Station, //startDate, endDate, userKey)
	//r := GetHTTPRequest(website)
	
	fmt.Println("START OF THE API.go FILE")

	injectionTest("Nottingham")

	fmt.Println("END OF THE API.go FILE")
}

func injectionTest(city string) {
	db, err := sql.Open("postgres", "postgresql://test:test@test")
	if err != nil {
		// return err
	}

	var count int

	row := db.QueryRow("SELECT COUNT(*) FROM t WHERE city=" + city) //nolint:safesql
	if err := row.Scan(&count); err != nil {
		// return err
	}

	row = db.QueryRow("SELECT COUNT(*) FROM t WHERE city=?", city)
	if err := row.Scan(&count); err != nil {
		// return err
	}

	return
}

// func GetHTTPRequest(website string) (r string) {
// 	response, err := http.Get(website)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		r = err.Error()
// 	} else {
// 		r = string(body)
// 	}

// 	return
// }