// Package engine implements the risk scoring logic for the supply-chain monitor.
package engine

import (
	"supply-chain-monitor/models"
)

const (
	maxDistance = 5000.0 // km – used to normalise the distance factor to [0,1]

	weightWeather  = 0.4
	weightTraffic  = 0.3
	weightDistance = 0.3
)

// RiskEngine calculates a risk score for each shipment based on environmental
// and operational factors.
type RiskEngine struct{}

// NewRiskEngine returns a new RiskEngine instance.
func NewRiskEngine() *RiskEngine {
	return &RiskEngine{}
}

// Process annotates a shipment pointer in-place with a RiskScore and RiskLevel.
//
// Formula:
//
//	risk_score = (weather_severity/10 * 0.4) + (traffic_condition/10 * 0.3) + (distance_factor * 0.3)
//
// All raw scores are normalised to the [0,1] range before weighting.
func (e *RiskEngine) Process(s *models.Shipment) {
	// Normalise each factor to [0, 1]
	weatherFactor := clamp(s.WeatherSeverity/10.0, 0, 1)
	trafficFactor := clamp(s.TrafficCondition/10.0, 0, 1)
	distanceFactor := clamp(s.DistanceKM/maxDistance, 0, 1)

	score := (weatherFactor * weightWeather) +
		(trafficFactor * weightTraffic) +
		(distanceFactor * weightDistance)

	s.RiskScore = score
	s.RiskLevel = classifyRisk(score)
}

// classifyRisk maps a continuous score to a discrete risk level.
func classifyRisk(score float64) models.RiskLevel {
	switch {
	case score > 0.6:
		return models.RiskHigh
	case score > 0.3:
		return models.RiskMedium
	default:
		return models.RiskLow
	}
}

// clamp restricts v to the range [lo, hi].
func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
