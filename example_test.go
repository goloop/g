package g_test

import (
	"fmt"

	g "github.com/goloop/g/v2"
)

func ExampleIf() {
	max := g.If(3 > 2, 3, 2)
	fmt.Println(max)
	// Output: 3
}

func ExampleIfFunc() {
	// Only the chosen branch runs, so the nil pointer is never dereferenced.
	var p *int
	v := g.IfFunc(p != nil,
		func() int { return *p },
		func() int { return -1 },
	)
	fmt.Println(v)
	// Output: -1
}

func ExampleIn() {
	fmt.Println(g.In(5, 1, 2, 3, 4, 5))
	fmt.Println(g.In("x", "a", "b", "c"))
	// Output:
	// true
	// false
}

func ExampleClamp() {
	fmt.Println(g.Clamp(15, 0, 10))
	fmt.Println(g.Clamp(-3, 0, 10))
	fmt.Println(g.Clamp(7, 0, 10))
	// Output:
	// 10
	// 0
	// 7
}

func ExampleSum() {
	fmt.Println(g.Sum(3, 5, 7, 1, 9, 2))
	// Output: 27
}

func ExampleMaxOk() {
	if v, ok := g.MaxOk(3, 1, 2); ok {
		fmt.Println(v)
	}
	var empty []int
	_, ok := g.MaxOk(empty...)
	fmt.Println(ok)
	// Output:
	// 3
	// false
}

func ExampleRange() {
	fmt.Println(g.Range(5))
	fmt.Println(g.Range(1, 10, 2))
	fmt.Println(g.Range(10, 0, -2))
	// Output:
	// [0 1 2 3 4]
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}

func ExampleSort() {
	xs := []int{5, 2, 7, 8, 1, 9}
	g.Sort(xs)
	fmt.Println(xs)
	g.Sort(xs, true)
	fmt.Println(xs)
	// Output:
	// [1 2 5 7 8 9]
	// [9 8 7 5 2 1]
}

func ExampleStringToDate() {
	t, err := g.StringToDate("2023-07-07")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format("2006-01-02"))
	// Output: 2023-07-07
}
