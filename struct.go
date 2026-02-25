package main

import(
	"fmt"
)

type MhsInterf interface{
	hitungYudisium() string
}

type MhsS1 struct{
	nim int
	ipk float64
}
type MhsS2 struct{
	nim int
	ipk float64
}

func (mhs MhsS1) hitungYudisium() string{
	if mhs.ipk >= 3.5 {
		return "Cumlaude"
	}
	return "AWIKWOK"
}

func (mhs MhsS2) hitungYudisium() string{
	return "Cumlaude"
}

func prosesMhs(mhs MhsInterf){
	fmt.Println("Yudisium:", mhs.hitungYudisium())
}

func main(){
	var mhs1 MhsS1 = MhsS1{
		nim: 12222,
		ipk: 3.12,
	}

	var mhs2 MhsS2 = MhsS2{
		nim: 12222,
		ipk: 3.12,
	}

	prosesMhs(mhs1)
	prosesMhs(mhs2)
}
