// Package services – shipment service (list and lookup).
package services

import (
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
)

// ShipmentService exposes read operations on the shipment store.
type ShipmentService struct {
	repo *repository.ShipmentRepository
}

// NewShipmentService creates a ShipmentService backed by the given repository.
func NewShipmentService(repo *repository.ShipmentRepository) *ShipmentService {
	return &ShipmentService{repo: repo}
}

// GetAll returns every shipment in the system (unfiltered, unpaginated).
// Used internally by the analytics layer.
func (s *ShipmentService) GetAll() []models.Shipment {
	return s.repo.GetAllShipments()
}

// GetByID returns the shipment for id, or (nil, false) if not found.
func (s *ShipmentService) GetByID(id string) (*models.Shipment, bool) {
	return s.repo.GetShipmentByID(id)
}

// ListFiltered applies the given filter and pagination to the shipment store
// and returns a structured PaginatedResponse.
func (s *ShipmentService) ListFiltered(
	filter models.ShipmentFilter,
	pp models.PageParams,
) (models.PaginatedResponse, error) {
	page, total, err := s.repo.FilterAndPage(filter, pp)
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
