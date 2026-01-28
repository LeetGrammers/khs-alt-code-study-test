func characterReplacement(s string, k int) int {
    hash := make(map[byte]int)
    l := 0
    r := 0
    ans := 0
    for r < len(s) {
        hash[s[r]]++
        r++
        for !helper(hash, k, r, l) {
            hash[s[l]]--
            l++
        }
        ans = max(ans, r-l) // count in hash value
    }
    return ans
}

func helper(hash map[byte]int, k, r, l int) bool {
    maxVal := 0
    for key, _ := range hash {
        maxVal = max(maxVal, hash[key])
    }
    return (((r-l) - maxVal) <= k)
}
