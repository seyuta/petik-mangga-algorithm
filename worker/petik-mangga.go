package worker

import (
	"fmt"

	"github.com/seyuta/petik-mangga-algorithm/pohon"
)

func PetikMangga() {
	sourceDatas := make(chan []int)

	go func() {
		defer close(sourceDatas)

		pohon.Mangga(sourceDatas)

		fmt.Println("musim mangga telah usai sampai disini, silahkan tekan enter untuk pergi dari pohon")
	}()

	go func() {
		for data := range sourceDatas {
			fmt.Printf("%v mangga telah di petik\n", len(data))
			fmt.Println("==========================")
		}
	}()
}
