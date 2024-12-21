func numJewelsInStones(jewels string, stones string) int {
    stonesMap := make(map[byte]int, 0)
    for i:=0;i<len(stones);i++ {
        stonesMap[stones[i]]++
    }

    sum:=0
    for i:=0;i<len(jewels);i++ {
        sum+=stonesMap[jewels[i]]
    }
    return sum
}