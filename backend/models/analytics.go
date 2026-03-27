// Package models – analytics result types and shared pagination/filter types.
package models

// DelayRateResult is the payload returned by GET /analytics/delay-rate.
type DelayRateResult struct {
	TotalShipments   int     `json:"total_shipments"`
	DelayedShipments int     `json:"delayed_shipments"`
	DelayRate        float64 `json:"delay_rate_percent"`
}

// RouteDelayStat ranks a route by number of delays.
type RouteDelayStat struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	DelayCount  int    `json:"delay_count"`
}

// CarrierPerformance holds delay statistics per carrier.
type CarrierPerformance struct {
	Carrier          string  `json:"carrier"`
	TotalShipments   int     `json:"total_shipments"`
	DelayedShipments int     `json:"delayed_shipments"`
	DelayRate        float64 `json:"delay_rate_percent"`
	AvgDelayDays     float64 `json:"avg_delay_days"`
}

// AvgDeliveryTime summarises mean delivery durations per carrier.

type AvgDeliveryTime struct {
	Carrier         string  `json:"carrier"`
	AvgDeliveryDays float64 `json:"avg_delivery_days"`
}

// ── Pagination & Filtering ───────────────────────────────────────────────────

// PaginatedResponse is the standard JSON envelope returned by all list endpoints.
type PaginatedResponse struct {
	Page         int         `json:"page"`
	Limit        int         `json:"limit"`
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	Data         interface{} `json:"data"`
}

// ShipmentFilter holds optional filter criteria passed in via query parameters.
// Empty strings mean "no filter" for that field.
type ShipmentFilter struct {
	Origin      string // ?origin=
	Destination string // ?destination=
	Carrier     string // ?carrier=
	Mode        string // ?mode=
	// Status can be "delayed" or "ontime" (empty = return all)
	Status string // ?status=
}

// PageParams holds validated pagination parameters.
type PageParams struct {
	Page  int // 1-based
	Limit int // records per page, capped at 100
}

// Offset returns the zero-based start index for slicing.
func (p PageParams) Offset() int { return (p.Page - 1) * p.Limit }

// TotalPages computes the number of pages for a given total record count.
func TotalPages(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	pages := total / limit
	if total%limit != 0 {
		pages++
	}
	return pages
}
