package zigzag_conversion

import "bytes"

func convert(s string, numRows int) string {
	if numRows < 0 {
		panic("Whyyyyy?")
	}
	
	if numRows == 1 {
		return s
	}

	bufs := make([]bytes.Buffer, numRows)

	var direction int
	for bc, i := 0, 0; i < len(s); i++ {
		bufs[bc].WriteByte(s[i])

		if bc == 0 {
			direction = 1 // go down
		} else if bc == numRows - 1 {
			direction = -1 // go up
		}

		bc = bc + direction

	}

	var res bytes.Buffer
	for _, buf := range bufs {
		res.Write(buf.Bytes())
	}

	return res.String()
}
