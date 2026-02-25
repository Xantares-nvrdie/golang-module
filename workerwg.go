package main

import(
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// wg di-passing by reference
	// karena decrement counter jika fungsi sudah selesai
	defer wg.Done() 
	fmt.Printf("Worker %d mulai\n", id)
	time.Sleep(time.Second) //1 set simulasi kerja
	fmt.Printf("Worker %d selesai\n", id)
}

func main() {
	//import sync untuk WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)  // inc counter
		go worker(i, &wg)
	}
	// block sampai counter =  0 (semua worker selesai)
	wg.Wait() 
	fmt.Println("Semua worker selesai")
}
