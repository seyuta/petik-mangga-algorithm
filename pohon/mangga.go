package pohon

import (
	"fmt"
	"time"
)

func Mangga(dataChan chan<- []int) {
	var sourceDatas []int
	for i := 1; i <= 30; i++ {
		var data int
		sourceDatas = append(sourceDatas, data)
		fmt.Println("hari ke", i)
		fmt.Println("jumlah mangga =", len(sourceDatas))
		fmt.Println("==================")

		// process data with max buffer size
		if len(sourceDatas) >= 10 {
			// creating a slice from buffer with max buffer size
			processSlice := make([]int, 10)
			copy(processSlice, sourceDatas[0:10])

			// emptying buffer
			sourceDatas = sourceDatas[:0]

			// give data to channel
			dataChan <- processSlice
		}

		time.Sleep(100 * time.Millisecond)
	}

	if len(sourceDatas) > 0 {
		dataChan <- sourceDatas
	}
}
