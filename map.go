package main

import(
	"fmt"
	"math/rand"
)

func getStatus() string {
	if rand.Float64() >= 0.5{
		return "WADUH"
	} else {
		return "OK"
	}
}

func main(){
	//seed
	


	nilai := make(map[string]int)
	nilai["Budi"] = 95
	nilai["Wati"] = 98
	nilai["Jima"] = 92
	nilai["X"] = 92
	nilaiBudi := nilai["Budi"]
	fmt.Println(nilaiBudi)
	fmt.Println(nilai["Budi"])

	nilaiX := nilai["X"]
	fmt.Println(nilaiX) //0 karena emg nilainya atau key "X" gaada

	//solution:
	//comma OK
	//hasil ini variabel lokal aja
	if hasil, ok:= nilai["X"]; ok {
		fmt.Println("ada, dan nilainya:", hasil)
	} else {
		fmt.Println("Nilai tdk ditemukan")
	}

	fmt.Println(getStatus())
	if val := getStatus(); val == "OK"{
		fmt.Println("BERHASILLLL")
	} else {
		fmt.Println("GAGAL")
	}

	for k, v := range nilai {
		fmt.Println(k,v)
	}

	for k:= range nilai {
		fmt.Println(k)
	}
	for _,v:= range nilai {
		fmt.Println(v)
	}
}
