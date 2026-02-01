func maxDistance(nums1 []int, nums2 []int) int {
    i := 0
    j := 0
    ans := 0
    for i < len(nums1) && j < len(nums2){
        if nums1[i] <= nums2[j] {
            j++
        } else {
            i++
        }
        ans = max(ans, j-i-1)
    }
    return ans
}