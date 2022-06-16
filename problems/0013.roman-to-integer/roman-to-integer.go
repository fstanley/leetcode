package leetcode

var roman = map[string]map[string]int{
	"I": {
		"val": 1,
		"V":   4,
		"X":   9,
	},
	"V": {
		"val": 5,
	},
	"X": {
		"val": 10,
		"L":   40,
		"C":   90,
	},
	"L": {
		"val": 50,
	},
	"C": {
		"val": 100,
		"D":   400,
		"M":   900,
	},
	"D": {
		"val": 500,
	},
	"M": {
		"val": 1000,
	},
}

func romanToInt(s string) int {
	decimal := 0
	for i := 0; i < len(s); i++ {
		c0 := string(s[i])
		c1 := ""
		if i < len(s)-1 {
			c1 = string(s[i+1])
		}
		t, ok := roman[c0][c1]
		if !ok {
			decimal += roman[c0]["val"]
		} else {
			decimal += t
			i++
		}
	}
	return decimal
}
