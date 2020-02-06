package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

const userKey string = "XXX"

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
	startDate := `2019-05-25`
	//endDate := `2019-05-31`

	//website := fmt.Sprintf(`https://api.meteostat.net/v1/history/daily?station=%s&start=%s&end=%s&key=%s`, cityInfo[0].Station, //startDate, endDate, userKey)
	//r := GetHTTPRequest(website)
	
	fmt.Println(startDate)
}

func GetHTTPRequest(website string) (r string) {
	response, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		r = err.Error()
	} else {
		r = string(body)
	}

	return
}