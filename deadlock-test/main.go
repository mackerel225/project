package main

import (
	"sync" // deadlock testing
	"sync/atomic" // deadlock testing
	DLock "../deadlock"
)

func main() {
	deadlockTest()
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
		// return err
	}
}