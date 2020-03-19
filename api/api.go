package main

import (
	"fmt"
	"database/sql"
	"sync" // deadlock testing
	"sync/atomic" // deadlock testing
	DLock "../deadlock"
)

// const userKey string = "XXXXXXXXXXX"

// type cityStation struct {
// 	City       string
// 	Station    string
// }

// var cityInfo = [5]cityStation{
// 	cityStation{"London", "03772"},
// 	cityStation{"Cambridge", "03571"},
// 	cityStation{"Nottingham", "03354"},
// 	cityStation{"Manchester", "03334"},
// 	cityStation{"Leeds", "EGNM0"},
// }

func main() {
	//startDate := `2019-05-25`
	//endDate := `2019-05-31`

	//website := fmt.Sprintf(`https://api.meteostat.net/v1/history/daily?station=%s&start=%s&end=%s&key=%s`, cityInfo[0].Station, //startDate, endDate, userKey)
	//r := GetHTTPRequest(website)
	
	fmt.Println("START OF THE API.go FILE")
	//deadlockTest()

	injectionTest("Nottingham")
	

	fmt.Println("END OF THE API.go FILE")
}

func restore() func() {
	opts := DLock.Opts
	return func() {
		DLock.Opts = opts
	}
}
 
func deadlockTest() {
	defer restore()()
	DLock.Opts.DeadlockTimeout = 0
	var deadlocks uint32
	DLock.Opts.OnPotentialDeadlock = func() {
		atomic.AddUint32(&deadlocks, 1)
	}
	var a DLock.RWMutex // RWMutex allows for multiple access calls, whereas writers have to wait for each other
	var b DLock.Mutex // Allows for only one goroutine to access variable at given time, i.e. Mutual Exclusion
	var wg sync.WaitGroup // Waits for goroutines to finish
	wg.Add(1)
	go func() {
		defer wg.Done() // Done means the gorotuine has finished and will remove one counter from WaitGroup
		a.Lock()
		b.Lock()
		b.Unlock()
		a.Unlock()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		b.Lock()
		a.RLock()
		a.RUnlock()
		b.Unlock()
	}()
	wg.Wait()
	if atomic.LoadUint32(&deadlocks) != 1 {
		fmt.Println("expected 1 deadlock, detected", deadlocks)
	}
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