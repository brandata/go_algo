package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(13)
	fmt.Println("My favourite number is", rand.Intn(10))
}
