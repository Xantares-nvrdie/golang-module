package main

import(
	"fmt"
)
func multiReturn(a int, b int) (int, int){
	return a + b, a * b
}

func tambah(a, b int) int {
	return a + b
}

func kali(a, b int) int {
	return a * b
}

func applyOperation(a int, b int, operation func(int, int) int) int{
	return operation(a, b)
}

func main(){
	var a, b int
	a = 5
	b = 3

	fmt.Println(multiReturn(a,b))

	x, y := multiReturn(a,b)
	fmt.Println(x,y)

	fmt.Println(applyOperation(9,9, tambah))

}
