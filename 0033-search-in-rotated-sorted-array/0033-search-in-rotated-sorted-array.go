func search(nums []int, target int) int {
    n := len(nums)
    l := 0
    r := n
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] == target {
            return mid
        }
        // 왼쪽 정렬됨
        if nums[mid] > nums[l]{
            if nums[mid] > target && nums[l] <= target{
                r = mid
            } else {
                l = mid+1
            }
        // 오른쪽 정렬됨
        } else {
            if nums[mid] < target && nums[r-1] >= target {
                l = mid+1
            } else {
                r = mid
            }
        }
    }
    return -1
}