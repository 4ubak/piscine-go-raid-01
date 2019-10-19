package piscine

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	for a := 1; a <= x; a++ {
		for b := 1 ; b <= y ; b++ {
			if (a == 1 && b == 1) || (a == 1 && b == y) || (a == x && b == 1) || (a == x && b == y) {
				z01.PrintRune('0')
			} else if b == 1 && (a != x || a != 1) {
				z01.PrintRune('|')
			} else if b == y && (a != 1 || a != x) {
				z01.PrintRune('|')
			} else if a == 1 && (b != 1 || b != y) {
				z01.PrintRune('-')
			} else if a == x && (b != 1 || b != y) {
				z01.PrintRune('-')
			} else {
				z01.PrintRune(' ')
			}
			
		}
		z01.PrintRune(10)
	}
}
