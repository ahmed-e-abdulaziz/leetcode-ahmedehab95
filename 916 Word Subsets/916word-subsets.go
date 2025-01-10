func wordSubsets(words1 []string, words2 []string) []string {
	words2Freq := calcFreq(words2[0])
	for i := 1; i < len(words2); i++ {
        wordFreq := calcFreq(words2[i])
		for j := 0; j < len(words2[i]); j++ {
            letter := words2[i][j]
			words2Freq[letter] = max(words2Freq[letter], wordFreq[letter])
		}
	}

	res := []string{}
	for _, word := range words1 {
		freq := calcFreq(word)
		isUniversal := true
		for w2Letter, w2LetterFreq := range words2Freq {
			if w1Freq, ok := freq[w2Letter]; !ok || w1Freq < w2LetterFreq {
				isUniversal = false
			}
		}
		if isUniversal {
			res = append(res, word)
		}
	}

	return res
}

func calcFreq(s string) map[byte]int {
	freq := map[byte]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	return freq
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}