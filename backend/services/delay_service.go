// Package services – delay service (delayed and high-risk shipments).
package services

import (
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
)

// DelayService provides operations focussed on delayed shipments.
type DelayService struct {
	repo *repository.ShipmentRepository
}

// NewDelayService creates a DelayService backed by the given repository.
func NewDelayService(repo *repository.ShipmentRepository) *DelayService {
	return &DelayService{repo: repo}
}

// GetDelayed returns all delayed shipments (unfiltered, unpaginated).
// Used by the analytics service internally.
func (d *DelayService) GetDelayed() []models.Shipment {
	return d.repo.GetDelayedShipments()
}

// GetHighRisk returns all HIGH-risk shipments (unfiltered, unpaginated).
func (d *DelayService) GetHighRisk() []models.Shipment {
	return d.repo.GetHighRiskShipments()
}

// GetDelayedPaged returns a paginated page of delayed shipments.
// An optional ShipmentFilter can further narrow the result.
func (d *DelayService) GetDelayedPaged(
	filter models.ShipmentFilter,
	pp models.PageParams,
) (models.PaginatedResponse, error) {
	// Force the status filter to "delayed", overriding any caller-supplied value.
	filter.Status = "delayed"
	page, total, err := d.repo.FilterAndPage(filter, pp)
	if err != nil {
		return models.PaginatedResponse{}, err
	}
	return models.PaginatedResponse{
		Page:         pp.Page,
		Limit:        pp.Limit,
		TotalRecords: total,
		TotalPages:   models.TotalPages(total, pp.Limit),
		Data:         page,
	}, nil
}

// GetHighRiskPaged returns a paginated page of HIGH-risk shipments.
func (d *DelayService) GetHighRiskPaged(
	pp models.PageParams,
) (models.PaginatedResponse, error) {
	// High-risk filter: status is not directly "delayed", but RiskLevel HIGH.
	// Use the repository helper that already knows about RiskLevel.
	all := d.repo.GetHighRiskShipments()
	total := len(all)

	start := pp.Offset()
	if start >= total {
		return models.PaginatedResponse{
			Page: pp.Page, Limit: pp.Limit,
			TotalRecords: total,
			TotalPages:   models.TotalPages(total, pp.Limit),
			Data:         []models.Shipment{},
		}, nil
	}
	end := start + pp.Limit
	if end > total {
		end = total
	}
	return models.PaginatedResponse{
		Page:         pp.Page,
		Limit:        pp.Limit,
		TotalRecords: total,
		TotalPages:   models.TotalPages(total, pp.Limit),
		Data:         all[start:end],
	}, nil
}
