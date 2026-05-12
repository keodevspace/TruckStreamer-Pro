package processor

import (
	"fmt"
	"sync"
)

// ProcessStats handles heavy calculations in a concurrent thread (Goroutine)
func ProcessStats(truckId string, speeds []float64, wg *sync.WaitGroup) {
	defer wg.Done()

	var total float64
	for _, s := range speeds {
		total += s
	}

	average := total / float64(len(speeds))
	fmt.Printf("[ANALYTICS] Truck: %s | Average Speed: %.2f km/h | Status: Processed\n", truckId, average)
}