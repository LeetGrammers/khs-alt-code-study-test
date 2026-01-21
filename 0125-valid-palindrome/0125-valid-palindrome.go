import(
    "strings"
)
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l := 0
    r := len(s)-1
    for l < r {
        for l < r {
            if (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9'){
                l++
                continue
            } else {
                break
            }
        }

        for l < r  {
            if (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
                r--
                continue
            } else {
                break
            }
        }

        if s[l] == s[r] {
            l++
            r--
        } else {
            return false
        }
    }
    return true
}