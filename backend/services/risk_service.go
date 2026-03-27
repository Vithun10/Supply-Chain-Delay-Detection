// Package services – risk service (query helpers for risk-level filtering).
package services

import (
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
)

// RiskService provides risk-related query helpers.
type RiskService struct {
	repo *repository.ShipmentRepository
}

// NewRiskService creates a RiskService backed by the given repository.
func NewRiskService(repo *repository.ShipmentRepository) *RiskService {
	return &RiskService{repo: repo}
}

// GetByRiskLevel returns all shipments matching a particular risk level.
func (r *RiskService) GetByRiskLevel(level models.RiskLevel) []models.Shipment {
	all := r.repo.GetAllShipments()
	var result []models.Shipment
	for _, s := range all {
		if s.RiskLevel == level {
			result = append(result, s)
		}
	}
	return result
}

// RiskSummary aggregates counts per risk level.
func (r *RiskService) RiskSummary() map[models.RiskLevel]int {
	counts := map[models.RiskLevel]int{
		models.RiskLow:    0,
		models.RiskMedium: 0,
		models.RiskHigh:   0,
	}
	for _, s := range r.repo.GetAllShipments() {
		counts[s.RiskLevel]++
	}
	return counts
}
