package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var roll = flag.Int("d", 20, "Sides of a dice to roll")

func numberGen(n int) (int error) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(n-1) + 1)
	return
}

func main() {
	flag.Parse()
	numberGen(*roll)
}
