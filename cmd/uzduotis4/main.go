package main

import (
	"fmt"
)

var (
	matricos1 = [][]string{
		{
			"ABCDS",
			"FZHIJ",
			"KLMNO",
			"PQTER",
			"UVWXG",
		},
		{
			"AZCDO",
			"FGHJI",
			"KVMNT",
			"PQRES",
			"ULWXB",
		},
		{
			"ABCDS",
			"FZHIJ",
			"KPMNO",
			"LGRET",
			"UVWXQ",
		},
	}

	// 1573
	// 110592
	// 46656
	matricos = [][]string{
		{ // 13. 4 * 4 * 3 * 2 = 96
			// 12. 3 2 3 4 = 72
			"@   S",
			"     ",
			" L   ",
			"  T  ",
			"    G",
		},
		{ // 11. 2 * 4 * 1 * 4 = 32
			// 10. 2 3 4 1 = 24
			"@    ",
			" G   ",
			"    T",
			"    S",
			" L   ",
		},
		{ // 11. 4 * 3 * 3 * 1 = 36
			// 10. 3 1 3 3 = 27
			"@   S",
			"     ",
			"     ",
			"LG  T",
			"     ",
		},
	}
)

func main() {
	fmt.Println("ketvirta")
	fmt.Println("1200")
}
