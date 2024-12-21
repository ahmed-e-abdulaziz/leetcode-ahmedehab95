func canConstruct(ransomNote string, magazine string) bool {
    if len(magazine) < len(ransomNote) {
        return false
    }
    freq := make([]int, 26)
    for _, c := range magazine {
        freq[c-'a']++
    }

    for _, c := range ransomNote {
        freq[c-'a']--
		if freq[c-'a']<0 {
			return false
		}
    }
    return true
}