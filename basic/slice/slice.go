package main

import "fmt"

func main() {
	// [...] represent array which length is not defined
	var months = [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "Oktober", "November", "Desember"}

	var slice1 = months[4:7]

	fmt.Println(slice1)

	// len slice
	fmt.Println(len(slice1))

	// cap slice
	fmt.Println(cap(slice1))

	// update array also update the slice
	// months[5] = "Diubah"
	fmt.Println(slice1)

	fmt.Println("=====================================")
	fmt.Println(slice1)

	// update slice also update the array, update slice1 index 0 means update months index 4 based on [4:7]
	// slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	fmt.Println("=====================================")
	var slice2 = months[2:3]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bulan Baru")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(months)
	fmt.Println(slice2)

	fmt.Println("==================- make -===================")
	// make(type, length, capacity)
	newSlice := make([]string, 0, 1)

	// newSlice[0] = "Muhammad"
	// this will increase the len
	newSlice = append(newSlice, "a")
	newSlice = append(newSlice, "b")
	newSlice = append(newSlice, "c")
	newSlice = append(newSlice, "d")
	newSlice = append(newSlice, "e")
	newSlice = append(newSlice, "f")
	newSlice = append(newSlice, "g")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	copySlice[0] = "Muhammad"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice), cap(newSlice))
	fmt.Println(copySlice)
	fmt.Println(len(copySlice), cap(copySlice))

	fmt.Println("==================- -===================")

	iniArray := [5]int{1, 2, 3, 4, 5} // or [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(len(iniArray))
	fmt.Println(iniSlice)
	fmt.Println(len(iniSlice))
}
