package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(1999, time.November, 23, 0, 0, 0, 0, time.UTC)
	out, err := ADtoBS(date)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	fmt.Println(BSDaysAfter(out, -6))
}
