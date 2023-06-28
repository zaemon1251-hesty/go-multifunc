package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	once.Do(f)
	once.Do(f)
	once.Do(f)
}

func f() {
	fmt.Println("hoge")
}
