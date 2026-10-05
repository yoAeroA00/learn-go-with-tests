package countdown

// package main

import (
	"fmt"
	"io"
	"time"
)

const (
	countStartNum = 3
	endLine       = "Go!"
)

// func countDownFrom(from int) iter.Seq[int] {
// 	return func(yield func(int) bool) {
// 		for i := from; i > 0; i-- {
// 			if !yield(i) {
// 				return
// 			}
// 		}
// 	}
// }

func Countdown(writer io.Writer, spy Sleeper) {
	// for i := range countDownFrom(3) {
	for i := countStartNum; i > 0; i-- {
		fmt.Fprintln(writer, i)
		spy.Sleep(time.Second * 1)
		// time.Sleep(time.Second * 1)
	}
	fmt.Fprint(writer, endLine)

	// fmt.Fprint(writer, "3\n2\n1\nGo!")
}

// func main() {
// 	Countdown(os.Stdout, &DSleeper{})
// }
