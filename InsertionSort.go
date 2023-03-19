package main

import "fmt"

func main() {
	fmt.Println(insertionSort())
}

func insertionSort() []int {
	var in int = 1
	var slice []int

	for in != 0 {
		fmt.Scan(&in)
		
		if in == 0 {
			break
		}

		var inserito bool
		for pos := 0; pos < len(slice); pos++ {
			if slice[pos] > in {

				slice = append(slice, 0)
				for i := len(slice)-2; i >= pos; i-- {
					slice[i+1] = slice[i]
				}
				slice[pos] = in

				inserito = true
				break
			}
		}
		if !inserito {
			slice = append(slice, in)
		}

		fmt.Println(slice)
	}
	return slice
}
