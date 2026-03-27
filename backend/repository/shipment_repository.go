// Package repository provides in-memory storage and query functions for shipments.
package repository

import (
	"strings"
	"supply-chain-monitor/models"
	"sync"
)

// ShipmentRepository holds the in-memory collection of processed shipments and
// provides thread-safe query methods used by the service layer.
type ShipmentRepository struct {
	mu        sync.RWMutex
	shipments []models.Shipment
	byID      map[string]*models.Shipment // index for O(1) lookups by ShipmentID
}

// NewShipmentRepository returns an empty repository.
func NewShipmentRepository() *ShipmentRepository {
	return &ShipmentRepository{
		byID: make(map[string]*models.Shipment),
	}
}

// BulkLoad replaces all stored shipments atomically.  It rebuilds the ID index
// in the same call.  This is called once at startup after concurrent processing.
func (r *ShipmentRepository) BulkLoad(shipments []models.Shipment) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shipments = shipments
	r.byID = make(map[string]*models.Shipment, len(shipments))
	for i := range r.shipments {
		r.byID[r.shipments[i].ShipmentID] = &r.shipments[i]
	}
}

// GetAllShipments returns a copy of the full shipment slice.
func (r *ShipmentRepository) GetAllShipments() []models.Shipment {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]models.Shipment, len(r.shipments))
	copy(result, r.shipments)
	return result
}

// GetShipmentByID returns the shipment with the given ID, or (nil, false) if
// not found.
func (r *ShipmentRepository) GetShipmentByID(id string) (*models.Shipment, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.byID[id]
	if !ok {
		return nil, false
	}
	// Return a copy to avoid callers mutating shared state.
	cp := *s
	return &cp, true
}

// FilterAndPage applies optional filter predicates to the full shipment slice,
// then returns a single bounds-checked page.
//
//   - filter    – any combination of Origin, Destination, Carrier, Mode, Status
//     (empty string = "no constraint" for that field)
//   - pp        – validated page / limit parameters
//
// Returns (page slice, total matched records, error).
func (r *ShipmentRepository) FilterAndPage(
	filter models.ShipmentFilter,
	pp models.PageParams,
) ([]models.Shipment, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// ── Apply filters ─────────────────────────────────────────────────────────
	matched := make([]models.Shipment, 0, 512)
	for _, s := range r.shipments {
		if !matchesFilter(s, filter) {
			continue
		}
		matched = append(matched, s)
	}

	total := len(matched)

	// ── Paginate ──────────────────────────────────────────────────────────────
	start := pp.Offset()
	if start >= total {
		// Page beyond the last record – return empty slice (not an error).
		return []models.Shipment{}, total, nil
	}
	end := start + pp.Limit
	if end > total {
		end = total
	}

	return matched[start:end], total, nil
}

// matchesFilter returns true when the shipment satisfies all non-empty filter fields.
func matchesFilter(s models.Shipment, f models.ShipmentFilter) bool {
	if f.Origin != "" && !strings.EqualFold(s.Origin, f.Origin) {
		return false
	}
	if f.Destination != "" && !strings.EqualFold(s.Destination, f.Destination) {
		return false
	}
	if f.Carrier != "" && !strings.EqualFold(s.Carrier, f.Carrier) {
		return false
	}
	if f.Mode != "" && !strings.EqualFold(s.Mode, f.Mode) {
		return false
	}
	switch strings.ToLower(f.Status) {
	case "delayed":
		if !s.DelayDetected {
			return false
		}
	case "ontime":
		if s.DelayDetected {
			return false
		}
	}
	return true
}

// GetDelayedShipments returns all shipments where DelayDetected is true (unfiltered).
// Kept for internal use by the analytics service.
func (r *ShipmentRepository) GetDelayedShipments() []models.Shipment {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []models.Shipment
	for _, s := range r.shipments {
		if s.DelayDetected {
			result = append(result, s)
		}
	}
	return result
}

// GetHighRiskShipments returns all shipments classified as HIGH risk (unfiltered).
// Kept for internal use by services that don't need pagination.
func (r *ShipmentRepository) GetHighRiskShipments() []models.Shipment {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []models.Shipment
	for _, s := range r.shipments {
		if s.RiskLevel == models.RiskHigh {
			result = append(result, s)
		}
	}
	return result
}

// Clear removes all stored shipments and resets the ID index.
// Used when a new dataset is uploaded.
func (r *ShipmentRepository) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shipments = nil
	r.byID = make(map[string]*models.Shipment)
}

// Count returns the total number of stored shipments.
func (r *ShipmentRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.shipments)
}
