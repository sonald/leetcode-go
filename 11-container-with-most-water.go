package main

func maxArea(height []int) int {
    var i, j = 0, len(height)-1
    var max = 0
    for i <= j {
        var h, area = 0, 0
        if height[i] > height[j] {
            h = height[j]
            area = (j - i) * h
            j--
        } else {
            h = height[i]
            area = (j - i) * h
            i++
        }
        if area > max {
            max = area
        }
    }

    return max
}

