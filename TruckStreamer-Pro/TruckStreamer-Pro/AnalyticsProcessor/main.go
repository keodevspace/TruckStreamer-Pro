package main

import (
	"AnalyticsProcessor/processor" // Adjust module name if necessary
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Analytics Processor started. Watching shared data...")

	// Infinite loop to simulate a background worker
	for {
		time.Sleep(15 * time.Second)

		var wg sync.WaitGroup

		// Simulated data batch extracted from the shared JSON file
		dataBatch := map[string][]float64{
			"Peterbilt_389": {75.5, 82.0, 79.2},
			"Mack_Anthem":    {68.4, 71.0, 70.5},
		}

		fmt.Println("--- Starting Parallel Calculation Cycle ---")

		for truckId, speeds := range dataBatch {
			wg.Add(1)
			// Executing the calculation in parallel using Goroutines
			go processor.ProcessStats(truckId, speeds, &wg)
		}

		wg.Wait()
		fmt.Println("--- Cycle Finished ---")
	}
}