package main

import(
	"fmt"
	"sync"
)
var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
   defer wg.Done()
   mu.Lock()
   //untuk memastikan unlock selalu terjadi	
   defer mu.Unlock() 
   counter++ //critical section
}

func main() {
   var wg sync.WaitGroup
   for i := 0; i < 5; i++ {
       wg.Add(1)
       go increment(&wg)
   }
   wg.Wait()
   fmt.Println("Counter:", counter)
}
