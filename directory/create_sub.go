package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("Japan/Tokyo/Shinjuku", 0777)
	if err != nil {
		fmt.Println(err)
	}
}
