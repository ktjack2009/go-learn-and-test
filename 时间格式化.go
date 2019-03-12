package main

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func main() {
	// Format := "2006-01-02 15:04:05"
	// Format := "2006-01-02 15:04:05"
	// Format := "2006-01-02 15:04:05"
	// Format := "2006-01-02 15:04:05"
	// Format := "2006-01-02"
	// a := time.Now()
	// f := a.Format(Format)
	// fmt.Println(f)

	// var b JsonTime
	// fmt.Println(b)
	a, _ := time.Parse("2006-01-02", "2019-03-04")
	fmt.Println(a)
}
