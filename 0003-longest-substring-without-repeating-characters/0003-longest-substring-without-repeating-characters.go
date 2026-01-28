func lengthOfLongestSubstring(s string) int {
    ans := 0
    l := 0
    r := 0
    hash := make(map[byte]bool)
    for r < len(s) {
        if !hash[s[r]] {
            hash[s[r]] = true
            r++
        } else {
            delete(hash, s[l])
            l++
        }
        ans = max(ans, r-l)
    }
    return ans
}