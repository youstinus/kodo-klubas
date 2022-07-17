package main

import (
	"strconv"
	"strings"
	"time"
)

var (
	data = `A10 08000915-4 10201415-9 15001530-10 16301645-3
B5 08150845-1 10251115-2 14001530-2 16301645-4 17001710-5
C4 10151020-3 10451215-3 14001630-1 16301645-2 17101725-3
D5 09300955-3 11201225-2 15101530-2 15301650-4
E7 09000945-8 12001255-7 14101530-2 15301820-1
F3 09000905-2 11201355-6 15101530-3 15401740-3 17501900-4
G3 08001015-4 11201305-3 13201430-2 15301540-2
H5 08201035-4 11201245-1 15101530-1 16201740-7 18001820-3
I8 07450915-1 11201135-6 15101530-6 16501740-8
J6 09100915-2 11201345-2 15101530-7 15351640-4 17201755-2`
)

// 9825
func main() {
	println("penkta")
	lines := strings.Split(data, "\n")
	sum := 0
	for _, line := range lines {
		sections := strings.Split(line, " ")
		code := sections[0]
		codeNum := code[1:]
		times := sections[1:]
		numGlobal, _ := strconv.Atoi(codeNum)
		for _, meet := range times {
			parts := strings.Split(meet, "-")
			countStr := parts[1]
			timesStr := parts[0]
			num, _ := strconv.Atoi(countStr)
			startStr := timesStr[:4]
			endStr := timesStr[4:]
			start, _ := time.Parse("1504", startStr)
			end, _ := time.Parse("1504", endStr)
			durat := int(end.Sub(start).Minutes())
			sum += durat * min(num, numGlobal)
		}
	}
	println(sum)
}

func min(one, two int) int {
	if one < two {
		return one
	}

	return two
}
