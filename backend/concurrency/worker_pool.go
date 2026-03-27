// Package concurrency provides a worker pool that processes shipments concurrently.
package concurrency

import (
	"supply-chain-monitor/engine"
	"supply-chain-monitor/models"
	"supply-chain-monitor/utils"
	"sync"
)

// ProcessorFunc is the signature every processing stage must satisfy.
type ProcessorFunc func(s *models.Shipment)

// WorkerPool fans out shipment processing across a fixed number of goroutines.
// It applies the provided processors in order for every shipment it receives.
type WorkerPool struct {
	workerCount int
	processors  []ProcessorFunc
	delayEng    *engine.DelayEngine
	riskEng     *engine.RiskEngine
}

// NewWorkerPool creates a WorkerPool with workerCount goroutines wired up to
// the delay and risk engines.
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		delayEng:    engine.NewDelayEngine(),
		riskEng:     engine.NewRiskEngine(),
	}
}

// Process sends all shipments through the worker pool and returns the fully
// annotated slice.  The caller provides an input slice; the function returns a
// new slice (same backing array) with DelayDays, DelayDetected, RiskScore, and
// RiskLevel populated for every record.
//
// Internally it:
//  1. Spawns `workerCount` goroutines.
//  2. Feeds each shipment (by index) through a channel.
//  3. Each worker runs delay detection then risk scoring on its shipment.
//  4. Waits for all workers to finish before returning.
func (wp *WorkerPool) Process(shipments []models.Shipment) []models.Shipment {
	utils.Logger.Printf("Processing shipments concurrently with %d workers", wp.workerCount)

	// Channel carries the index of each shipment to be processed.
	jobs := make(chan int, wp.workerCount*2)

	var wg sync.WaitGroup

	// Launch workers
	for i := 0; i < wp.workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				wp.delayEng.Process(&shipments[idx])
				wp.riskEng.Process(&shipments[idx])
			}
		}()
	}

	// Feed jobs
	for i := range shipments {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	utils.Logger.Printf("Concurrent processing complete for %d shipments", len(shipments))
	return shipments
}
