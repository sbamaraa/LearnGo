package divide_two_integers

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("duuuh")
	}

	positiveQ := true
	if dividend < 0 && divisor > 0 {
		positiveQ = false
	}
	if dividend > 0 && divisor < 0 {
		positiveQ = false
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	q := 0
	r := dividend
	for (r >= divisor) {
		cnt := 0
		for r >= (divisor << (cnt + 1)) {
			cnt++
		}

		q += 1 << cnt
		r -= (divisor << cnt)
	}



	if positiveQ {
		if ((1 << 31) - 1) < q {
			return (1 << 31) - 1;
		}

		return q
	} else {
		if q < -(1 << 31) {
			return -(1 << 31);
		}
		return q * -1;
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
