package quad

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= x; i++ { // i rows and j columns
		for j := 1; j <= y; j++ {
			if i == 1 || i == x {
				if j == 1 || j == y {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if j == 1 || j == y {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
