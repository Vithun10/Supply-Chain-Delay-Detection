// Package concurrency – shipment channel helpers.
package concurrency

import "supply-chain-monitor/models"

// ShipmentProducer sends shipments onto a channel and closes it when done.
// It is intended to be launched as a goroutine.
//
//	ch := make(chan models.Shipment, bufferSize)
//	go ShipmentProducer(ch, shipments)
func ShipmentProducer(ch chan<- models.Shipment, shipments []models.Shipment) {
	for _, s := range shipments {
		ch <- s
	}
	close(ch)
}

// ShipmentConsumer drains the channel and collects processed shipments.
// This is a convenience helper for ad-hoc pipeline construction outside the
// WorkerPool; the main pipeline uses WorkerPool.Process instead.
func ShipmentConsumer(ch <-chan models.Shipment) []models.Shipment {
	var result []models.Shipment
	for s := range ch {
		result = append(result, s)
	}
	return result
}
