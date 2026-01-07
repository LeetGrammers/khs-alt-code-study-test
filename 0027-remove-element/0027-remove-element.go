func removeElement(nums []int, val int) int {
    answer := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[answer] = nums[i]
            answer += 1
        }
    }
    return answer
}