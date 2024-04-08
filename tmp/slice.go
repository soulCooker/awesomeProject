package tmp

import "fmt"

type SliceItem struct {
	val     int32
	trigger bool
	text    string
}

func test() {
	arr := [10][]SliceItem{}
	s := arr[0:6]
	s = append(s, make([]SliceItem, 6, 6))
	fmt.Printf("val:%+v, len:%d, cap:%d\n", s, len(s), cap(s))
	for _, i := range s {
		fmt.Printf("val:%+v, len:%d, cap:%d\n", i, len(i), cap(i))
	}
	fmt.Println()

	a := arr[6:]
	fmt.Printf("val:%+v, len:%d, cap:%d\n", a, len(a), cap(a))
	for _, i := range a {
		fmt.Printf("val:%+v, len:%d, cap:%d\n", i, len(i), cap(i))
	}
	fmt.Println()

	b := a[:2]
	fmt.Printf("val:%+v, len:%d, cap:%d\n", b, len(b), cap(b))
	for _, i := range b {
		fmt.Printf("val:%+v, len:%d, cap:%d\n", i, len(i), cap(i))
	}
	fmt.Println()

	b = append(b, make([]SliceItem, 3, 3), make([]SliceItem, 3, 3), make([]SliceItem, 3, 3), make([]SliceItem, 3, 3))

	fmt.Printf("val:%+v, len:%d, cap:%d\n", b, len(b), cap(b))
	for _, i := range b {
		fmt.Printf("val:%+v, len:%d, cap:%d\n", i, len(i), cap(i))
	}
	fmt.Println()

	fmt.Printf("val:%+v, len:%d, cap:%d\n", a, len(a), cap(a))
	for _, i := range a {
		fmt.Printf("val:%+v, len:%d, cap:%d\n", i, len(i), cap(i))
	}

	fmt.Println(a[cap(a)+1])
}
