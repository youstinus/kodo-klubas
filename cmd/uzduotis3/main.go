package main

var (
	counts   = map[rune]int{'A': 10, 'B': 6, 'C': 2, 'D': 7, 'E': 3, 'F': 15, 'G': 3, 'H': 3, 'I': 7, 'J': 7} // A10 B6 C2 D7 E3 F15 G3 H3 I7 J7
	komandos = [][]string{
		{
			"ABF",
			"FA",
			"C",
			"JIA",
		},
		{
			"ABC",
			"A",
			"A",
			"B",
			"B",
		},
		{
			"FG",
			"G",
			"G",
			"CAB",
		},
		{
			"B",
		},
		{
			"IJ",
			"JIA",
			"ABC",
			"BC",
		},
		{
			"CD",
			"C",
		},
		{

			"DEI",
		},
		{
			"ABCDD",
			"AJ",
		},
		{
			"HG",
			"G",
			"AB",
		},
		{
			"BC",
			"CJ",
			"J",
		},
	}
)

func main() {
	println("trecia")

	dauginam := 1

	for _, komanda := range komandos {

		unique := map[rune]int{}

		for _, narys := range komanda {

			for _, patiekalas := range narys {
				num := counts[patiekalas]
				if num > 0 {
					counts[patiekalas]--
					unique[patiekalas]++
				}
			}
		}

		kiekis := len(unique)
		if kiekis > 0 {
			dauginam *= kiekis
		}
	}

	println(dauginam)
}
