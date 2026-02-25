package main

import(
	"fmt"
)

func main() {
	hari := "senin"
	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("Hari kerja melelahkan")
	case "sabtu", "minggu":
		fmt.Println("libuurrr")
	default:
		fmt.Println("tidak valid")
	}

	//kalau mau lanjut ke case berikutnya dia harus pake fallthrough
}
