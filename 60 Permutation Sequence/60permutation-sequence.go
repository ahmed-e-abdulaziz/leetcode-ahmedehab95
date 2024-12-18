import (
	"fmt"
	"strings"
)

var factorials = []int{
	0,
	1,
	2,
	6,
	24,
	120,
	720,
	5040,
	40320,
	362880,
	3628800,
}

func getPermutation(n int, k int) string {
	sb := strings.Builder{}
	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i + 1
	}

	// j is 0 based k
	j := k - 1
	for i := n - 1; i >= 0; i-- {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%d", set[0]))
			break
		}

		// iv = j / (n-1)!
		iv := j / factorials[i]

		// j1 = j0 % (n-1)!
		j = j % factorials[i]

		// Write the value at idx iv and remove it from the set
		v := set[iv]
		set = append(set[:iv], set[iv+1:]...)
		sb.WriteString(fmt.Sprintf("%d", v))
	}

	return sb.String()
}