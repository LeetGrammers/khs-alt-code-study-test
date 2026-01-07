func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i, num := range nums {
        n := target - num
        if _, ok := hash[n]; ok {
            return []int{i, hash[n]}
        }
        hash[num] = i
    }
    return []int{}
}