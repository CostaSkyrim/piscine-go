package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

type players struct {
	hand []int
}

func DealAPackOfCards(deck []int) {
	p1 := players{}
	p2 := players{}
	p3 := players{}
	p4 := players{}

	for i := 1; i <= 12; i++ {
		if i < 4 {
			p1.hand = append(p1.hand, i)
			fmt.Printf("Player 1: %v", p1.hand)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else if i < 7 {
			p2.hand = append(p2.hand, i)
			fmt.Printf("Player 2: %v", p2.hand)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else if i < 10 {
			p3.hand = append(p3.hand, i)
			fmt.Printf("Player 3: %v", p3.hand)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			p4.hand = append(p4.hand, i)
			fmt.Printf("Player 4: %v", p4.hand)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
