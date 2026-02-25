package main

import(
	"fmt"
	"time"
)

func sayHello(){
	for i:=0; i<5; i++{
		fmt.Println("Hello GO!")
		time.Sleep(1 * time.Second)
	}
}


func main() {
    go sayHello()

	for i:=0; i<5; i++{
		fmt.Println("Hello friend!")
		time.Sleep(1 * time.Second)
	}
}
