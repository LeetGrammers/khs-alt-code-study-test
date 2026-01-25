import(
    "slices"
)
func minimumDifference(nums []int, k int) int {
    ans := 1000000
    slices.Sort(nums)
    for i := 0; i < len(nums)-k+1; i++ {
        ans = min(ans, nums[i+(k-1)] - nums[i])
    }
    return ans
}