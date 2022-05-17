package main

import (
	"log"
	"semester-project-smooth-sort/sort"
	"semester-project-smooth-sort/utils"
	"time"
)

// CYCLES - Количество повторений
const CYCLES = 10

func main() {
	smoothRandomResults := make([]utils.Result, 0)
	smoothSortedResults := make([]utils.Result, 0)
	heapRandomResults := make([]utils.Result, 0)
	heapSortedResults := make([]utils.Result, 0)

	// Проверка на сгенерированных данных
	for data := range utils.ReadData() {
		smoothTimings := make([]time.Duration, 0)
		heapTimings := make([]time.Duration, 0)

		// Для каждого набора данных
		for nums := range data.Nums {

			// Тестируем HeapSort
			timer := utils.NewTimer()
			for i := 0; i < CYCLES; i++ {
				// создаем копию
				numsCopy := make([]int, len(nums), len(nums))
				copy(numsCopy, nums)

				// сортируем
				timer.Start()
				sort.HeapSort(numsCopy)
				timer.Stop()

				if !utils.IsSorted(numsCopy) {
					panic("doesn't have sorted")
				}
			}
			heapTimings = append(heapTimings, timer.Passed()/CYCLES)

			// Тестируем SmoothSort
			timer = utils.NewTimer()
			for i := 0; i < CYCLES; i++ {
				// создаем копию
				numsCopy := make([]int, len(nums), len(nums))
				copy(numsCopy, nums)

				// сортируем
				timer.Start()
				sort.SmoothSort(numsCopy)
				timer.Stop()

				if !utils.IsSorted(numsCopy) {
					panic("doesn't have sorted")
				}
			}
			smoothTimings = append(smoothTimings, timer.Passed()/CYCLES)
		}

		smoothRandomResults = append(smoothRandomResults, utils.Result{
			Count: data.Count,
			Dur:   utils.Average(smoothTimings),
		})
		log.Println("smooth_random", data.Count, utils.Average(smoothTimings))

		heapRandomResults = append(heapRandomResults, utils.Result{
			Count: data.Count,
			Dur:   utils.Average(heapTimings),
		})
		log.Println("heap_random", data.Count, utils.Average(heapTimings))
	}

	// Проверка на отсортированных данных
	for data := range utils.GetSortedData() {
		smoothTimings := make([]time.Duration, 0)
		heapTimings := make([]time.Duration, 0)

		// Для каждого набора данных
		for nums := range data.Nums {

			// Тестируем HeapSort
			timer := utils.NewTimer()
			for i := 0; i < CYCLES; i++ {
				// создаем копию
				numsCopy := make([]int, len(nums), len(nums))
				copy(numsCopy, nums)

				// сортируем
				timer.Start()
				sort.HeapSort(numsCopy)
				timer.Stop()

				if !utils.IsSorted(numsCopy) {
					panic("doesn't have sorted")
				}
			}
			heapTimings = append(heapTimings, timer.Passed()/CYCLES)

			// Тестируем SmoothSort
			timer = utils.NewTimer()
			for i := 0; i < CYCLES; i++ {
				// создаем копию
				numsCopy := make([]int, len(nums), len(nums))
				copy(numsCopy, nums)

				// сортируем
				timer.Start()
				sort.SmoothSort(numsCopy)
				timer.Stop()

				if !utils.IsSorted(numsCopy) {
					panic("doesn't have sorted")
				}
			}
			smoothTimings = append(smoothTimings, timer.Passed()/CYCLES)
		}

		smoothSortedResults = append(smoothSortedResults, utils.Result{
			Count: data.Count,
			Dur:   utils.Average(smoothTimings),
		})
		log.Println("smooth_sorted", data.Count, utils.Average(smoothTimings))

		heapSortedResults = append(heapSortedResults, utils.Result{
			Count: data.Count,
			Dur:   utils.Average(heapTimings),
		})
		log.Println("heap_sorted", data.Count, utils.Average(heapTimings))
	}

	utils.WriteResult("heap_random.csv", heapRandomResults)
	utils.WriteResult("heap_sorted.csv", heapSortedResults)
	utils.WriteResult("smooth_random.csv", smoothRandomResults)
	utils.WriteResult("smooth_sorted.csv", smoothSortedResults)
}
