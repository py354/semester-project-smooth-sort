package utils

import (
	"time"
)

// Average - функция для вычисления среднего времени
func Average(durs []time.Duration) time.Duration {
	if len(durs) == 0 {
		return 0
	}

	s := time.Duration(0)
	for _, d := range durs {
		s += d
	}
	return s / time.Duration(len(durs))
}

// GetSortedData - передает канал для передачи отсортированных наборов данных (по 100, 1000...)
func GetSortedData() chan Data {
	c := make(chan Data, 5)

	// выполняется параллельно
	go func() {
		for _, count := range counts {
			data := Data{
				Count: count,
				Nums:  make(chan []int, 1),
			}

			data.Nums <- getSortedArray(count)
			close(data.Nums)

			c <- data
		}
		close(c)
	}()

	return c
}

func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i += 1 {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
