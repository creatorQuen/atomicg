package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

var totalOperations1 int32

func inc1() {
	atomic.AddInt32(&totalOperations1, 1) // атомарно
}

func main() {
	//r.FirstMain()
	//fmt.Println("-------------------------------")
	//SecondMain()
	//fmt.Println("-------------------------------\n")

	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pT))
	var ta, tb = T{1}, T{2}
	// store
	atomic.StorePointer(
		unsafePPT, unsafe.Pointer(&ta))
	fmt.Println(pT) // &{1}
	// load
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true
	// swap
	pa2 := atomic.SwapPointer(
		unsafePPT, unsafe.Pointer(&tb))
	fmt.Println((*T)(pa2) == &ta) // true
	fmt.Println(pT) // &{2}
	// compare and swap
	b := atomic.CompareAndSwapPointer(
		unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false
	b = atomic.CompareAndSwapPointer(
		unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true

	pT = &T{2}
	pTn = T{3}

	fmt.Println("----------------------------------------")

	p0 := new(int)   // p0 points to a zero int value.
	fmt.Println(p0)  // (a hex address string)
	fmt.Println(*p0) // 0
	fmt.Println("----------------------------------------")
	// x is a copy of the value at
	// the address stored in p0.
	x := *p0
	// Both take the address of x.
	// x, *p1 and *p2 represent the same value.
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false

	fmt.Println(p0)
	fmt.Println(p1)

	p3 := &*p0 // <=> p3 := &(*p0) <=> p3 := p0
	fmt.Println(p3)

	fmt.Println("----------------------------------------")
	//// Now, p3 and p0 store the same address.
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}

type T struct {x int}
var pT *T
var pTn T

func SecondMain()  {
	for i := 0; i < 1000; i++ {
		go inc1()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation = ", totalOperations1)
}


var (
	n uint64 = 97
	m uint64 = 1
	k int    = 2
)
const (
	a        = 3
	b uint64 = 4
	c uint32 = 5
	d int    = 6
)

func show(number uint64)  {
	fmt.Println(number)
}