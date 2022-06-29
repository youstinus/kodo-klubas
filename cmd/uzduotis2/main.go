package main

import (
	"fmt"
	"strconv"
)

var (
	floors = []int{4, 5, 6, 7}
	floor  = 2
	cnt    = 0
	data   = "3U2U3D1D5U7U2D4D3U2U1D1D2D4U5D2U1U1D2D3D1D1D1U1D2D5U6U1U2U4D2D1D2U1D2D1U4U1D1D2D3D8D2D1U3U2U3D1D3U7U2D4D3U2U1D1D2D3D4D2U1U1D2D3D8D2U1U1D2D5U6U1U2U4D2D1D2U1D2D1D4U1D1D2D3D1D3U1D"
)

func main() {
	fmt.Println("antra")

	tmp := 0
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			num, _ := strconv.Atoi(string(data[i]))
			tmp = num
			continue
		}

		if data[i] == byte('D') {
			if floor-tmp < -1 {
				continue
			}

			floor -= tmp

			if contains(floors, floor) {
				cnt++
			}
		} else if data[i] == byte('U') {
			if floor+tmp > 9 {
				continue
			}

			floor += tmp

			if contains(floors, floor) {
				cnt++
			}
		} else {
			fmt.Println("fail", string(data[i]), byte('U'))
		}
	}

	fmt.Println(floor, cnt)
}

func contains(ints []int, num int) bool {
	for _, a := range ints {
		if num == a {
			return true
		}
	}

	return false
}
