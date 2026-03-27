// Package tests – unit, table, and benchmark tests for the shipment domain.
package tests

import (
	"supply-chain-monitor/engine"
	"supply-chain-monitor/models"
	"supply-chain-monitor/repository"
	"supply-chain-monitor/services"
	"supply-chain-monitor/utils"
	"testing"
	"time"
)

// ── Helpers ───────────────────────────────────────────────────────────────────

func mustParse(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}

func makeShipment(id, origin, dest string, distKM, weather, traffic float64, expected, delivered string) models.Shipment {
	return models.Shipment{
		ShipmentID:           id,
		Origin:               origin,
		Destination:          dest,
		DistanceKM:           distKM,
		Carrier:              "TestCarrier",
		Mode:                 "road",
		WeatherSeverity:      weather,
		TrafficCondition:     traffic,
		ExpectedDeliveryDate: mustParse(expected),
		DeliveredDate:        mustParse(delivered),
	}
}

// ── DelayEngine Tests ─────────────────────────────────────────────────────────

func TestDelayEngine_NoDelay(t *testing.T) {
	eng := engine.NewDelayEngine()
	s := makeShipment("SHP001", "Mumbai", "Delhi", 1400, 3, 4, "2024-01-10", "2024-01-09")
	eng.Process(&s)
	if s.DelayDetected {
		t.Errorf("expected no delay, got DelayDetected=true")
	}
	if s.DelayDays != 0 {
		t.Errorf("expected DelayDays=0, got %.1f", s.DelayDays)
	}
}

func TestDelayEngine_WithDelay(t *testing.T) {
	eng := engine.NewDelayEngine()
	s := makeShipment("SHP002", "Chennai", "Kolkata", 1700, 5, 6, "2024-01-10", "2024-01-13")
	eng.Process(&s)
	if !s.DelayDetected {
		t.Error("expected DelayDetected=true, got false")
	}
	if s.DelayDays != 3 {
		t.Errorf("expected DelayDays=3.0, got %.1f", s.DelayDays)
	}
}

// Table-driven tests for the delay engine.
func TestDelayEngine_Table(t *testing.T) {
	eng := engine.NewDelayEngine()
	tests := []struct {
		name          string
		expected      string
		delivered     string
		wantDelayed   bool
		wantDelayDays float64
	}{
		{"on-time", "2024-03-01", "2024-03-01", false, 0},
		{"early", "2024-03-10", "2024-03-08", false, 0},
		{"1 day late", "2024-03-01", "2024-03-02", true, 1},
		{"5 days late", "2024-03-01", "2024-03-06", true, 5},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s := makeShipment("X", "A", "B", 100, 1, 1, tc.expected, tc.delivered)
			eng.Process(&s)
			if s.DelayDetected != tc.wantDelayed {
				t.Errorf("[%s] DelayDetected: want %v got %v", tc.name, tc.wantDelayed, s.DelayDetected)
			}
			if s.DelayDays != tc.wantDelayDays {
				t.Errorf("[%s] DelayDays: want %.0f got %.0f", tc.name, tc.wantDelayDays, s.DelayDays)
			}
		})
	}
}

// ── RiskEngine Tests ──────────────────────────────────────────────────────────

func TestRiskEngine_HighRisk(t *testing.T) {
	eng := engine.NewRiskEngine()
	s := makeShipment("SHP010", "A", "B", 5000, 10, 10, "2024-01-01", "2024-01-02")
	eng.Process(&s)
	if s.RiskLevel != models.RiskHigh {
		t.Errorf("expected HIGH risk, got %s (score %.3f)", s.RiskLevel, s.RiskScore)
	}
}

func TestRiskEngine_LowRisk(t *testing.T) {
	eng := engine.NewRiskEngine()
	s := makeShipment("SHP011", "A", "B", 10, 0, 0, "2024-01-01", "2024-01-01")
	eng.Process(&s)
	if s.RiskLevel != models.RiskLow {
		t.Errorf("expected LOW risk, got %s", s.RiskLevel)
	}
}

// ── DateUtils Tests ───────────────────────────────────────────────────────────

func TestDaysBetween(t *testing.T) {
	a := mustParse("2024-01-01")
	b := mustParse("2024-01-06")
	got := utils.DaysBetween(a, b)
	if got != 5 {
		t.Errorf("expected 5, got %.1f", got)
	}
}

func TestIsDelayed(t *testing.T) {
	expected := mustParse("2024-01-10")
	delivered := mustParse("2024-01-12")
	if !utils.IsDelayed(expected, delivered) {
		t.Error("expected IsDelayed to return true")
	}
}

// ── Repository Tests ──────────────────────────────────────────────────────────

func populatedRepo() *repository.ShipmentRepository {
	eng1 := engine.NewDelayEngine()
	eng2 := engine.NewRiskEngine()

	shipments := []models.Shipment{
		makeShipment("R001", "A", "B", 100, 2, 2, "2024-01-01", "2024-01-01"),
		makeShipment("R002", "C", "D", 4000, 9, 9, "2024-01-01", "2024-01-05"),
		makeShipment("R003", "E", "F", 200, 1, 1, "2024-01-01", "2024-01-03"),
	}
	for i := range shipments {
		eng1.Process(&shipments[i])
		eng2.Process(&shipments[i])
	}

	repo := repository.NewShipmentRepository()
	repo.BulkLoad(shipments)
	return repo
}

func TestRepository_GetByID(t *testing.T) {
	repo := populatedRepo()
	s, ok := repo.GetShipmentByID("R001")
	if !ok {
		t.Fatal("expected to find R001")
	}
	if s.ShipmentID != "R001" {
		t.Errorf("expected R001, got %s", s.ShipmentID)
	}
}

func TestRepository_GetDelayed(t *testing.T) {
	repo := populatedRepo()
	delayed := repo.GetDelayedShipments()
	if len(delayed) == 0 {
		t.Error("expected at least one delayed shipment")
	}
}

func TestRepository_GetHighRisk(t *testing.T) {
	repo := populatedRepo()
	highRisk := repo.GetHighRiskShipments()
	for _, s := range highRisk {
		if s.RiskLevel != models.RiskHigh {
			t.Errorf("expected RiskHigh, got %s for %s", s.RiskLevel, s.ShipmentID)
		}
	}
}

// ── ShipmentService Tests ─────────────────────────────────────────────────────

func TestShipmentService_GetAll(t *testing.T) {
	repo := populatedRepo()
	svc := services.NewShipmentService(repo)
	all := svc.GetAll()
	if len(all) != 3 {
		t.Errorf("expected 3 shipments, got %d", len(all))
	}
}

// ── Benchmarks ────────────────────────────────────────────────────────────────

func BenchmarkDelayEngine(b *testing.B) {
	eng := engine.NewDelayEngine()
	s := makeShipment("BNC001", "A", "B", 500, 5, 5, "2024-01-01", "2024-01-04")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Process(&s)
	}
}

func BenchmarkRiskEngine(b *testing.B) {
	eng := engine.NewRiskEngine()
	s := makeShipment("BNC002", "A", "B", 2500, 7, 6, "2024-01-01", "2024-01-02")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Process(&s)
	}
}

func BenchmarkRepository_BulkLoad(b *testing.B) {
	eng1 := engine.NewDelayEngine()
	eng2 := engine.NewRiskEngine()
	// Prepare 10,000 shipments
	shipments := make([]models.Shipment, 10_000)
	for i := range shipments {
		shipments[i] = makeShipment("X", "A", "B", float64(i%5000), float64(i%10), float64(i%10), "2024-01-01", "2024-01-02")
		eng1.Process(&shipments[i])
		eng2.Process(&shipments[i])
	}
	repo := repository.NewShipmentRepository()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.BulkLoad(shipments)
	}
}
