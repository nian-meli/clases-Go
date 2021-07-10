package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Println("Good night, dear!")

	for i := 0; i < 6; i++ {
		fmt.Print(emoji.KissingFace,
			emoji.FaceBlowingAKiss)
	}
}
