package sort

import (
	"fmt"
	"testing"
)

func TestBy_Sort(t *testing.T) {

	var people = []People{
		{"Rose", 58, 166},
		{"Daisley", 78, 184},
		{"Lumiya", 65, 179},
		{"Sola", 68, 177},
	}

	// 这些部分可以移动到people中
	heightFunc := func(p1, p2 *People) bool {
		return p1.Height < p2.Height
	}
	ageASCFunc := func(p1, p2 *People) bool {
		return p1.Age < p2.Age
	}
	ageDescFunc := func(p1, p2 *People) bool {
		return p1.Age > p2.Age
	}

	By(heightFunc).Sort(people)
	fmt.Println("By height ASC:", people)
	By(ageDescFunc).Sort(people)
	fmt.Println("By age DESC:", people)
	By(ageASCFunc).Sort(people)
	fmt.Println("By age ASC:", people)
}
