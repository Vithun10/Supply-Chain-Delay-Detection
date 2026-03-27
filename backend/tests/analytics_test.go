// Package tests – analytics service tests.
package tests

import (
	"supply-chain-monitor/engine"
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
	"supply-chain-monitor/services"
	"testing"
)

// buildAnalyticsRepo creates a repository loaded with a small known dataset.
func buildAnalyticsRepo() *repository.ShipmentRepository {
	de := engine.NewDelayEngine()
	re := engine.NewRiskEngine()

	raw := []models.Shipment{
		makeShipment("A001", "Mumbai", "Delhi", 1400, 8, 7, "2024-02-01", "2024-02-04"),    // delayed 3d, high risk
		makeShipment("A002", "Mumbai", "Delhi", 1400, 2, 2, "2024-02-01", "2024-01-30"),    // on-time, low risk
		makeShipment("A003", "Chennai", "Kolkata", 1700, 9, 9, "2024-02-05", "2024-02-10"), // delayed 5d, high risk
		makeShipment("A004", "Pune", "Hyderabad", 800, 1, 1, "2024-02-10", "2024-02-10"),   // on-time, low risk
		makeShipment("A005", "Mumbai", "Delhi", 1400, 5, 5, "2024-02-15", "2024-02-17"),    // delayed 2d, medium risk
	}
	for i := range raw {
		de.Process(&raw[i])
		re.Process(&raw[i])
	}
	repo := repository.NewShipmentRepository()
	repo.BulkLoad(raw)
	return repo
}

// TestAnalyticsService_DelayRate checks counts and rate calculation.
func TestAnalyticsService_DelayRate(t *testing.T) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	result := svc.DelayRate()

	if result.TotalShipments != 5 {
		t.Errorf("TotalShipments: want 5, got %d", result.TotalShipments)
	}
	// A001, A003, A005 are delayed
	if result.DelayedShipments != 3 {
		t.Errorf("DelayedShipments: want 3, got %d", result.DelayedShipments)
	}
	wantRate := 60.0
	if result.DelayRate != wantRate {
		t.Errorf("DelayRate: want %.1f, got %.1f", wantRate, result.DelayRate)
	}
}

// TestAnalyticsService_TopDelayedRoutes ensures sorting by delay count.
func TestAnalyticsService_TopDelayedRoutes(t *testing.T) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	routes := svc.TopDelayedRoutes(10)

	if len(routes) == 0 {
		t.Fatal("expected at least one route")
	}
	// Mumbai→Delhi has 2 delays (A001, A005); should be first
	top := routes[0]
	if top.Origin != "Mumbai" || top.Destination != "Delhi" {
		t.Errorf("expected Mumbai→Delhi as top route, got %s→%s", top.Origin, top.Destination)
	}
	if top.DelayCount != 2 {
		t.Errorf("expected DelayCount=2, got %d", top.DelayCount)
	}
}

// TestAnalyticsService_CarrierPerformance validates output structure.
func TestAnalyticsService_CarrierPerformance(t *testing.T) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	carriers := svc.CarrierPerformance()

	if len(carriers) == 0 {
		t.Error("expected non-empty carrier performance list")
	}
	for _, cp := range carriers {
		if cp.Carrier == "" {
			t.Error("carrier name must not be empty")
		}
		if cp.DelayRate < 0 || cp.DelayRate > 100 {
			t.Errorf("delay rate out of range: %.2f", cp.DelayRate)
		}
	}
}

// TestAnalyticsService_AvgDeliveryTime validates avg delivery time result.
func TestAnalyticsService_AvgDeliveryTime(t *testing.T) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	times := svc.AvgDeliveryTime()

	if len(times) == 0 {
		t.Error("expected non-empty avg delivery time result")
	}
	for _, at := range times {
		if at.AvgDeliveryDays < 0 {
			t.Errorf("negative avg delivery days for carrier %s", at.Carrier)
		}
	}
}

// BenchmarkDelayRate measures analytics computation speed.
func BenchmarkDelayRate(b *testing.B) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = svc.DelayRate()
	}
}

// BenchmarkTopDelayedRoutes measures route aggregation speed.
func BenchmarkTopDelayedRoutes(b *testing.B) {
	repo := buildAnalyticsRepo()
	svc := services.NewAnalyticsService(repo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = svc.TopDelayedRoutes(10)
	}
}
