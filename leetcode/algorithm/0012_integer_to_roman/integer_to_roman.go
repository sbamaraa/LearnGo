package integer_to_roman

import "bytes"

func intToRoman(num int) string {
	var res bytes.Buffer

	// How many thousands
	for i := 0; i < num / 1000; i++ {
		res.WriteString("M")
	}

	// Express the hundreds
	switch (num % 1000) / 100 {
	case 1: res.WriteString("C")
	case 2: res.WriteString("CC")
	case 3: res.WriteString("CCC")
	case 4: res.WriteString("CD")
	case 5: res.WriteString("D")
	case 6: res.WriteString("DC")
	case 7: res.WriteString("DCC")
	case 8: res.WriteString("DCCC")
	case 9: res.WriteString("CM")
	}

	// Express the tens
	switch (num % 100) / 10 {
	case 1: res.WriteString("X")
	case 2: res.WriteString("XX")
	case 3: res.WriteString("XXX")
	case 4: res.WriteString("XL")
	case 5: res.WriteString("L")
	case 6: res.WriteString("LX")
	case 7: res.WriteString("LXX")
	case 8: res.WriteString("LXXX")
	case 9: res.WriteString("XC")
	}

	// Express the digit
	switch num % 10 {
	case 1: res.WriteString("I")
	case 2: res.WriteString("II")
	case 3: res.WriteString("III")
	case 4: res.WriteString("IV")
	case 5: res.WriteString("V")
	case 6: res.WriteString("VI")
	case 7: res.WriteString("VII")
	case 8: res.WriteString("VIII")
	case 9: res.WriteString("IX")
	}

	return res.String()
}
