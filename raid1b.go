package piscine

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x <= 0 || y <= 0 {

		} else {
			for a := 1; a <= y; a++ {
				for b := 1; b <= x; b++ {
					if a == 1 && b == 1 {
						fmt.Printf("%s", "/")
					} else if a == 1 && b == x {
						z01.PrintRune(92)
					} else if a == y && b == 1 {
						z01.PrintRune(92)
					} else if a == y && b == x {
						fmt.Printf("%s", "/")
					} else if b == 1 && (a != y || a != 1) {
						fmt.Printf("%s", "*")
					} else if b == x && (a != 1 || a != y) {
						fmt.Printf("%s", "*")
					} else if a == 1 && (b != 1 || b != x) {
						fmt.Printf("%s", "*")
					} else if a == y && (b != 1 || b != x) {
						fmt.Printf("%s", "*")
					} else {
						fmt.Printf("%s", " ")
					}
	
				}
				fmt.Println()
			}
		}
	}
}
