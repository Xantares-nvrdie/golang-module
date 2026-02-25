package main;

import(
	"fmt"
)

func main(){
	arr1 := [3]int{10 ,20 ,30}
	arr1[2] = 999
	fmt.Println(arr1)

	var matrix [3][3] int
	matrix[0][0] = 999
	fmt.Println(matrix)

	arrDy := [][]int{}
	arrDy = append(arrDy, []int{10,20}, []int{30,40})
	fmt.Println(arrDy)

	arrDyCOPY := arrDy[0:2]
	fmt.Println(arrDyCOPY)

	//prove kalo slice tadi itu dia shared memory,
	//jadi si arrDyCopy ini dia sama dengan arrDy
	arrDy[0][1] = 88
	fmt.Println(arrDyCOPY)


	//supaya copy
	subSlice := make([]int, len(arrDy))
	copy(subSlice, arrDy)
}
