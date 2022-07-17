package main

var (
	abc          = "AĄBCČDEĘĖFGHIĮYJKLMNOPRSŠTUŲŪVZŽ"
	data         = "CKŲRG UZJZ UŲČIGĄ ĄĘKRURVG SĘKRUČĄ ČFJČZCS RA ĄUGCRVG CZŲRGČ TČJMCR S ŽARKUS"
	start        = "TELIA"
	letter1 rune = 'C'
	letter2 rune = 'T'
	count        = 32
)

func main() {
	println("pirma")
	abcRunes := []rune(abc)

	encIndex := 0
	resIndex := 0
	abcMap := map[rune]int{}
	for i, ab := range abcRunes {
		if ab == letter1 {
			encIndex = i
		}
		if ab == letter2 {
			resIndex = i
		}
		abcMap[ab] = i
	}

	diff := 0
	if encIndex < resIndex {
		diff = resIndex - encIndex
	} else {
		diff = encIndex - resIndex
	}

	runes := []rune(data)
	result := make([]rune, len(runes))
	for i, run := range runes {
		if run == 32 { // tarpas
			result[i] = run
			continue
		}
		position := abcMap[run]
		newPos := (position + diff) % count
		result[i] = abcRunes[newPos]
	}

	println(string(result), diff)
}
