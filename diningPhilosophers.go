package main

import(
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

func philosophers(id int, forks []sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()

	left := id
	right := (id + 1) % numPhilosophers

	first := left
	second := right
	if(left > right){
		first, second = right, left
	}

	for i := 1; i < 3; i++{
		fmt.Printf("Philosophers %d sedang berpikir\n", id)
		time.Sleep(time.Millisecond * 500)

		forks[first].Lock()
		forks[second].Lock()

		fmt.Printf("Philosophers %d sedang makan..\n", id)
		time.Sleep(time.Millisecond * 500)

		forks[first].Unlock()
		forks[second].Unlock()
	}
}

func main(){
	var wg sync.WaitGroup

	forks := make([]sync.Mutex, numPhilosophers)

	for i:=0; i < numPhilosophers;i++{
		wg.Add(1)
		go philosophers(i, forks, &wg)
	}
	wg.Wait()
	fmt.Println("Semua filsuf selesai makan.")
}
