package main

import (
	"fmt"
	"math/rand"
	"time"
)

type numbers []int

func newNumbers() numbers {
	numberList := numbers{}

	for i := 0; i <= 100; i++ {
		numberList = append(numberList, i)
	}

	return numberList
}

func newRandomNumbers() numbers {

	source := rand.NewSource(time.Now().UnixNano()) // create new source with the value of now in nanoseconds
	r := rand.New(source)                           // add source to rand struct

	numberList := numbers{}

	for i := 0; i <= 100; i++ {
		numberList = append(numberList, r.Intn(1000))
	}
	return numberList
}

func (n numbers) print() {
	for _, number := range n {
		fmt.Println(number)
	}
}

func (n numbers) evenOrOdd() {

	for i, number := range n {
		if n[i]%2 == 0 {
			fmt.Println(number, "IS EVEN,")
		}
		if n[i]%2 != 0 {
			fmt.Println(number, "IS ODD,")
		}
	}
}
