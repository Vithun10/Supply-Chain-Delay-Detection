// Package engine implements the delay detection logic for the supply-chain monitor.
package engine

import (
	"supply-chain-monitor/models"
	"supply-chain-monitor/utils"
)

// DelayEngine detects shipment delays by comparing expected vs actual delivery dates.
type DelayEngine struct{}

// NewDelayEngine creates and returns a new DelayEngine instance.
func NewDelayEngine() *DelayEngine {
	return &DelayEngine{}
}

// Process annotates a shipment pointer in-place with DelayDays and DelayDetected.
// It must be called before RiskEngine.Process so risk scoring may account for delay.
func (e *DelayEngine) Process(s *models.Shipment) {
	// Guard against zero-value dates which indicate missing CSV data.
	if s.ExpectedDeliveryDate.IsZero() || s.DeliveredDate.IsZero() {
		return
	}

	delayDays := utils.DaysBetween(s.ExpectedDeliveryDate, s.DeliveredDate)

	if delayDays > 0 {
		s.DelayDays = delayDays
		s.DelayDetected = true
		utils.Logger.Printf("Delay detected for shipment %s (%.1f days)", s.ShipmentID, delayDays)
	} else {
		s.DelayDays = 0
		s.DelayDetected = false
	}
}
