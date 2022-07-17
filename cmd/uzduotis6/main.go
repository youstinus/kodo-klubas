package main

var (
	mp          = map[rune]int{}
	mp1         = map[rune]int{}
	mp2         = map[rune]int{}
	mp3         = map[rune]int{}
	mas         = "AĄBCČDEĘĖFGHIĮYJKLMNOPRSŠTUŲŪVZŽ"
	masr        = []rune(mas)
	mas1        = "AĄBCČDEĘĖFGHIĮYJKLMNOPRSŠTUŲŪVZŽ"
	mas2        = "AĄBCČDEĘĖFGHIĮYJKLMNOPRSŠTUŲŪVZŽ"
	mas3        = "AĄBCČDEĘĖFGHIĮYJKLMNOPRSŠTUŲŪVZŽ"
	lenght      = 32
	x           = 1
	first       = mas[x-1]
	x1          = 1
	x2          = 1
	x3          = 1
	question    = "ĄJŲSIVBOČĄCĮRPĮYPŽEHĖKZIVĘYNZŪAEĮČŽDUUDĘĄMIAĖRKIGEŲRIŠČĖUYKTŲMĮŠR"
	questionr   = []rune(question)
	ques2       = "VBOČĄCĮRPĮYPŽEHĖKZIVĘYNZŪAEĮČŽDUUDĘĄMIAĖRKIGEŲRIŠČĖUYKTŲMĮŠR"
	ques2r      = []rune(ques2)
	encryp      = "ĄJŲSI"
	encrypr     = []rune(encryp)
	result      = "TELIA"
	resultr     = []rune(result)
	result2     = "TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽ"
	resultr2    = []rune(result2)
	mqp         = map[rune]int{}
	resultfull  = "TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽDUOTĮIRSKATINATOLIAUJUDĖTIĮPRIEKĮ"
	resultfullr = []rune(resultfull)
	asd         = "TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽČUOTĮIRSKATINATOLIAUJUDĖTIĮPRIEKI"
	asdr        = []rune(asd)
)

func main() {
	println("sesta")
	for i, a := range masr {
		mp[a] = i + 1
	}
	// sudek(32, 31, 1)
	// return
	// pradzia(30, 11, 2)
	// return

	// println(string(first))
	for x33 := 1; x33 <= len(masr); x33++ {
		for x22 := 1; x22 <= len(masr); x22++ {
			for x11 := 1; x11 <= len(masr); x11++ {
				pradzia(x11, x22, x33)
				// writeme(x11, x22, x33)
			}
		}
	}
}

// TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽČUOTĮIRSKATINATOLIAUJUDĖTIĮPRIEKI

func pradzia(c1, c2, c3 int) {
	i1, i2, i3 := c1, c2, c3
	var runes []rune
	for i := 0; i < len(resultr); i++ {
		sudek(c1, c2, c3)
		let := resultr[i]
		num1 := mp[let]
		// if c1 == 1 && c2 == 5 && c3 == 1 {
		// 	println("we")
		// }
		for k, v := range mp1 {
			if v == num1 {
				let = k
				break
			}
		}
		// c1 = c1%32 + 1
		// if c1 == i1 {
		// 	c2 = c2%32 + 1
		// 	if c2 == i2 {
		// 		c3 = c3%32 + 1
		// 	}
		// }
		// sudek(c1, c2, c3)
		num2 := mp[let]
		for k, v := range mp2 {
			if v == num2 {
				let = k
				break
			}
		}
		// c1 = c1%32 + 1
		// if c1 == i1 {
		// 	c2 = c2%32 + 1
		// 	if c2 == i2 {
		// 		c3 = c3%32 + 1
		// 	}
		// }
		// sudek(c1, c2, c3)
		num3 := mp[let]
		for k, v := range mp3 {
			if v == num3 {
				let = k
				break
			}
		}
		me := questionr[i]
		if let != me {
			return
		}
		c1 = c1%32 + 1
		if c1 == i1 {
			c2 = c2%32 + 1
			if c2 == i2 {
				c3 = c3%32 + 1
			}
		}
		runes = append(runes, let)
	}
	// c1 rotated
	//if string(runes) == question {
	println(string(runes), writeme(i1, i2, i3), i1, i2, i3)
	writeme(i1, i2, i3)
	//}
	/*if c1 != 1 && c2 != 1 && c3 != 1 {
		println(c1, c2, c3)
	}*/
	//writeme(i1, i2, i3)
}

// "ĄJŲSIVBOČĄCĮRPĮYPŽEHĖKZIVĘYNZŪAEĮČŽDUUDĘĄMIAĖRKIGEŲRIŠČĖUYKTŲMĮŠR"
// "TELIA"
//
func writeme(c1, c2, c3 int) string {
	i1, i2 := c1, c2
	// asd := "TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽDUOTĮIRSKATINATOLIAUJUDĖTIĮPRIEKĮ"
	// asd = "TELIAKODOKLUBASSVEIKINAĮVEIKUSUŽČUOTĮIRSKATINATOLIAUJUDĖTIĮPRIEKI"
	// asd = asd
	var runes []rune
	for i := 0; i < len(questionr); i++ {
		sudek(c1, c2, c3)
		let := questionr[i]

		num3 := mp3[let]
		for k, v := range mp {
			if v == num3 {
				let = k
				break
			}
		}
		num2 := mp2[let]
		for k, v := range mp {
			if v == num2 {
				let = k
				break
			}
		}
		num1 := mp1[let]
		for k, v := range mp {
			if v == num1 {
				let = k
				break
			}
		}

		c1 = c1%32 + 1
		if c1 == i1 {
			c2 = c2%32 + 1
			if c2 == i2 {
				c3 = c3%32 + 1
			}
		}

		runes = append(runes, let)
	}

	return string(runes)
	// if string(runes) == asd {
	// 	println(string(runes))
	// }
}

func sudek(c1, c2, c3 int) {
	for i := 0; i < len(masr); i++ {
		num := (32+i-c1+1)%32 + 1
		let := masr[i]
		mp1[let] = num
	}
	for i := 0; i < len(masr); i++ {
		num := (32+i-c2+1)%32 + 1
		let := masr[i]
		mp2[let] = num
	}
	for i := 0; i < len(masr); i++ {
		num := (32+i-c3+1)%32 + 1
		let := masr[i]
		mp3[let] = num
	}
}
