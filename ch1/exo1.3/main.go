// main.go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for n := 0; n < 10000; n++ {
		var s string
		sep := " "
		s = os.Args[0]
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
		}
	}
	elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
	fmt.Printf("%v elapsed\n", elapsed)
}
