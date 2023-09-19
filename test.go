package main

import (
	"fmt"
	"time"
)

func main() {
	// `now` is an instance of time.Time
	now := time.Now()
	fmt.Println(now.String())
}

