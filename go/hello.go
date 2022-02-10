package main

import (
	"fmt"
	// 	"io/ioutil"
	m "math"
	//	"net/http"
	//	"net/http"
	//	"os"
	//	"strconv"
)

func main() {
	fmt.Println("hello, world\n")
	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	renturn x + y, x * y
}

// source: https://learnxinyminutes.com/docs/go/