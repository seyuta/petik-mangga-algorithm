package main

import (
	"fmt"

	"github.com/seyuta/petik-mangga-algorithm/worker"
)

func main() {
	fmt.Println("musim mangga dimulai")
	worker.PetikMangga()

	fmt.Println("program berjalan")
	fmt.Scanln()
}
