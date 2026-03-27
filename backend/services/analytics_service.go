// Package services – analytics service computing logistics insights.
package services

import (
	"sort"
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
	"supply-chain-monitor/utils"
)

// AnalyticsService computes aggregated logistics analytics from shipment data.
type AnalyticsService struct {
	repo *repository.ShipmentRepository
}

// NewAnalyticsService creates an AnalyticsService backed by the given repository.
func NewAnalyticsService(repo *repository.ShipmentRepository) *AnalyticsService {
	return &AnalyticsService{repo: repo}
}

// DelayRate returns the overall delay rate as a structured result.
func (a *AnalyticsService) DelayRate() models.DelayRateResult {
	utils.Logger.Println("Analytics computed: delay rate")
	shipments := a.repo.GetAllShipments()
	total := len(shipments)
	delayed := 0
	for _, s := range shipments {
		if s.DelayDetected {
			delayed++
		}
	}
	rate := 0.0
	if total > 0 {
		rate = float64(delayed) / float64(total) * 100
	}
	return models.DelayRateResult{
		TotalShipments:   total,
		DelayedShipments: delayed,
		DelayRate:        rate,
	}
}

// TopDelayedRoutes returns the top N routes by delay count, sorted descending.
func (a *AnalyticsService) TopDelayedRoutes(n int) []models.RouteDelayStat {
	counts := make(map[string]*models.RouteDelayStat)

	for _, s := range a.repo.GetAllShipments() {
		if !s.DelayDetected {
			continue
		}
		key := s.Origin + "|" + s.Destination
		if _, ok := counts[key]; !ok {
			counts[key] = &models.RouteDelayStat{
				Origin:      s.Origin,
				Destination: s.Destination,
			}
		}
		counts[key].DelayCount++
	}

	stats := make([]models.RouteDelayStat, 0, len(counts))
	for _, v := range counts {
		stats = append(stats, *v)
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].DelayCount > stats[j].DelayCount
	})
	if n > 0 && n < len(stats) {
		stats = stats[:n]
	}
	return stats
}

// CarrierPerformance returns per-carrier delay and delivery statistics.
func (a *AnalyticsService) CarrierPerformance() []models.CarrierPerformance {
	type accumulator struct {
		total     int
		delayed   int
		totalDays float64
	}
	acc := make(map[string]*accumulator)

	for _, s := range a.repo.GetAllShipments() {
		if _, ok := acc[s.Carrier]; !ok {
			acc[s.Carrier] = &accumulator{}
		}
		acc[s.Carrier].total++
		if s.DelayDetected {
			acc[s.Carrier].delayed++
			acc[s.Carrier].totalDays += s.DelayDays
		}
	}

	result := make([]models.CarrierPerformance, 0, len(acc))
	for carrier, a := range acc {
		rate := 0.0
		avgDays := 0.0
		if a.total > 0 {
			rate = float64(a.delayed) / float64(a.total) * 100
		}
		if a.delayed > 0 {
			avgDays = a.totalDays / float64(a.delayed)
		}
		result = append(result, models.CarrierPerformance{
			Carrier:          carrier,
			TotalShipments:   a.total,
			DelayedShipments: a.delayed,
			DelayRate:        rate,
			AvgDelayDays:     avgDays,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].DelayRate > result[j].DelayRate
	})
	return result
}

// AvgDeliveryTime returns average delivery duration per carrier (in days).
func (a *AnalyticsService) AvgDeliveryTime() []models.AvgDeliveryTime {
	type acc struct {
		total int
		days  float64
	}
	accMap := make(map[string]*acc)

	for _, s := range a.repo.GetAllShipments() {
		if s.ExpectedDeliveryDate.IsZero() || s.DeliveredDate.IsZero() {
			continue
		}
		if _, ok := accMap[s.Carrier]; !ok {
			accMap[s.Carrier] = &acc{}
		}
		dur := s.DeliveredDate.Sub(s.ExpectedDeliveryDate).Hours() / 24
		// We measure actual span, not delay; use absolute value.
		if dur < 0 {
			dur = -dur
		}
		accMap[s.Carrier].days += dur
		accMap[s.Carrier].total++
	}

	result := make([]models.AvgDeliveryTime, 0, len(accMap))
	for carrier, a := range accMap {
		avg := 0.0
		if a.total > 0 {
			avg = a.days / float64(a.total)
		}
		result = append(result, models.AvgDeliveryTime{
			Carrier:         carrier,
			AvgDeliveryDays: avg,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Carrier < result[j].Carrier
	})
	return result
}
