package main

import (
	"fmt"
)

func main() {
	targetX1 := 235
	targetY1 := -118
	targetX2 := 259
	targetY2 := -62
	p1 := partOne(targetX1, targetY1, targetX2, targetY2)

	p2 := partTwo(targetX1, targetY1, targetX2, targetY2)

	fmt.Println()
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func partOne(tx1 int, ty1 int, tx2 int, ty2 int) int {
	max := 0
	for x := 1; x < tx2; x++ {
		for y := ty1; y < 1000; y++ {
			h := fire(tx1, ty1, tx2, ty2, x, y)
			if h > max {
				max = h
			}
		}
	}
	return max
}

func partTwo(tx1 int, ty1 int, tx2 int, ty2 int) int {
	c := 0
	for x := 1; x <= tx2; x++ {
		for y := ty1; y < 1000; y++ {
			h := fire(tx1, ty1, tx2, ty2, x, y)
			if h > ty1-1 {
				c++
			}
		}
	}
	return c
}

func fire(tx1 int, ty1 int, tx2 int, ty2 int, velx int, vely int) int {
	posx := 0
	posy := 0
	maxy := ty1 - 1
	tmaxy := ty1 - 1
	for {
		//fmt.Println(posx, posy, tmaxy)
		posx += velx
		posy += vely
		if velx > 0 {
			velx -= 1
		}
		vely -= 1
		if posy > tmaxy {
			tmaxy = posy
		}
		if posx > tx2 {
			break
		}
		if posy < ty1 {
			break
		}
		if posx >= tx1 && posx <= tx2 {
			if posy > ty2 {
				// still above
				continue
			}
			if posy < ty1 {
				// fell below
				break
			}
			// hit
			maxy = tmaxy
			break
		}
		if posy <= ty2 && posy >= ty1 {
			if posx < tx1 {
				// still left
				continue
			}
			if posx > tx2 {
				// to right
				break
			}
			// hit
			maxy = tmaxy
			break
		}
	}
	return maxy
}
