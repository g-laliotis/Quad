package quad

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ { // i rows and j columns
		for j := 1; j <= x; j++ {
			if j == 1 && (i == 1 || i == y) {
				z01.PrintRune('A')
			} else if j == x && (i == 1 || i == y) {
				z01.PrintRune('C')
			} else if ((i == 1 || i == y) && (j > 1 && j < x)) || (i > 1 || i < y) && (j == 1 || j == x) {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
