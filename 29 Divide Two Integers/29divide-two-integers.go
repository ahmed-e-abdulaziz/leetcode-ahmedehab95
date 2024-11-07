func divide(dividend int, divisor int) int {
    sign := true

    if divisor == 1 {
        return dividend
    }

    //Extract sign ("+" or "-")
    if (dividend >= 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        sign = false
    }

    //Calculate absolute value
    dividend = abs(dividend)
    divisor = abs(divisor)

    result := 0

    for dividend >= divisor { // This for loop will handle if desired result is odd
        temp := divisor
        count := 1

        // This for loop left shift temp(i.e, divisor) until its greater than dividend
        // but it multiply temp with 2 every time so we need the outer for loop to handle odd desird cases
        // example : 10/3 --> for this inner loop will multiply 3 with 2 every time
        // 3*2, (next at -> 3*4 -- break) leads to result = 2 afterwords outer-for-loop just 
        // increment result with 1, handeling desired odd results
        for dividend >= (temp << 1) {
            temp <<= 1
            count <<= 1
        }
        dividend -= temp
        result += count
    }

    // Put sign back in result
    if !sign {
        return -result
    }

    // handle edge cases
    if result > 2147483647 {
        return 2147483647
    } else {
        return result
    }
}

// function to calculate absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}