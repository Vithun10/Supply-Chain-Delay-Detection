// Package models contains the core domain types for the supply-chain monitor.
package models

import "time"

// RiskLevel classifies a shipment's computed risk.
type RiskLevel string

const (
	RiskLow    RiskLevel = "LOW"
	RiskMedium RiskLevel = "MEDIUM"
	RiskHigh   RiskLevel = "HIGH"
)

// Shipment represents a single logistics record loaded from the CSV dataset.
type Shipment struct {
	ShipmentID           string    `json:"shipment_id"`
	Origin               string    `json:"origin"`
	Destination          string    `json:"destination"`
	DistanceKM           float64   `json:"distance_km"`
	Carrier              string    `json:"carrier"`
	Mode                 string    `json:"mode"`
	ExpectedDeliveryDate time.Time `json:"expected_delivery_date"`
	DeliveredDate        time.Time `json:"delivered_date"`
	WeatherSeverity      float64   `json:"weather_severity"`  // 0–10 scale
	TrafficCondition     float64   `json:"traffic_condition"` // 0–10 scale

	// Computed fields – populated by the processing pipeline
	DelayDays     float64   `json:"delay_days"`
	DelayDetected bool      `json:"delay_detected"`
	RiskScore     float64   `json:"risk_score"`
	RiskLevel     RiskLevel `json:"risk_level"`
}

// Route is a helper type used in analytics (origin→destination pair).
type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// String returns a human-readable route label.
func (r Route) String() string {
	return r.Origin + " → " + r.Destination
}
