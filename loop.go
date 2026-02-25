package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 10; i++{
		fmt.Print(i )
	}

	//while (aslinya for juga)
	n := 0
	for n < 10 {
		fmt.Println("whileKW",n)
		n++
	}

	
	items := []string {"pepaya", "pisang", "jeruk", "semangka"}
	//urutannya selalu index dan value dari slice
	for index,value := range items {
		fmt.Println(index, value)
	}

	//bisa pake _ kalo ingin mengabaikan index
	for _,val := range items {
		fmt.Println(val)
	}

}
