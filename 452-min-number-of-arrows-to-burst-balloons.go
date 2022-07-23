package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	//fmt.Printf("%v\n", points)

	res := 0
	i := 0
	for i < len(points) {
		res++
		s, e := points[i][0], points[i][1]
		for i++; i < len(points); i++ {
			if points[i][0] > e {
				break
			}
			if s < points[i][0] {
				s = points[i][0]
			}
			if e > points[i][1] {
				e = points[i][1]
			}
		}
	}
	return res
}

func findMinArrowShots2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	maxright := points[0][1]
	res := 1
	for _, p := range points {
		if p[0] > maxright {
			maxright = p[1]
			res++
		}
	}

	return res
}
