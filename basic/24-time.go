package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	s1 := strconv.FormatInt(now, 10)
	fmt.Printf("%T", now)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%v", s1)
}
