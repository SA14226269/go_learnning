package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	foo := make([]int, 5, 10)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 90
	fmt.Println(foo,len(foo),cap(foo))
    
	a := make([]int, 32)
	fmt.Println(cap(a),len(a))
	b := a[1:16]
	a = append(a,1)
	a[2] = 42
	fmt.Println(a,b)
	fmt.Println(cap(a),len(a))
	fmt.Println(cap(b),len(b))

	test1 := make([]int, 1024)
	test1 = append(test1, 1) //1.25 * cap(1024)
	fmt.Println(cap(test1),len(test1))
}
