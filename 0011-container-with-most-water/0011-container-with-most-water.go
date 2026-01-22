func maxArea(height []int) int {
    maxWater := 0
    curWater := 0
    l := 0
    r := len(height)-1
    for l < r {
        curWater = (r-l) * min(height[l], height[r])
        maxWater = max(maxWater, curWater)
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return maxWater
}