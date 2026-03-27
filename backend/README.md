# Concurrent Supply Chain Delay Detection and Monitoring System

A production-grade Go backend server that ingests a 500k-row CSV logistics dataset, detects shipment delays, computes risk scores, and exposes a full REST API via the [Gin](https://github.com/gin-gonic/gin) framework.

---

## Quick Start

```powershell
# 1. Place your dataset
copy supply_chain_dataset_full_500k.csv backend\data\

# 2. Install dependencies
cd backend
go mod tidy

# 3. Run the server
go run .\cmd\main.go
# Server starts at http://localhost:8080
```

---

## Project Structure

```
backend/
├── cmd/main.go                      ← Entry point
├── config/config.go                 ← Server configuration
├── models/
│   ├── shipment.go                  ← Shipment struct
│   └── analytics.go                 ← Analytics result types
├── utils/
│   ├── csv_loader.go                ← Loads CSV dataset
│   ├── logger.go                    ← Shared logger
│   ├── date_utils.go                ← Date math helpers
│   └── internal/dateparser/         ← Multi-layout date parser
├── engine/
│   ├── delay_engine.go              ← Delay detection logic
│   └── risk_engine.go               ← Risk score calculation
├── concurrency/
│   ├── worker_pool.go               ← 10-worker goroutine pool
│   └── shipment_channel.go          ← Channel producer/consumer
├── repository/
│   └── shipment_repository.go       ← Thread-safe in-memory store
├── services/
│   ├── shipment_service.go
│   ├── delay_service.go
│   ├── analytics_service.go
│   └── risk_service.go
├── controllers/
│   ├── shipment_controller.go
│   ├── analytics_controller.go
│   └── tracking_controller.go
├── routes/routes.go                 ← Gin route registration
├── tests/
│   ├── shipment_test.go             ← Unit + table + benchmark tests
│   └── analytics_test.go
└── data/
    └── supply_chain_dataset_full_500k.csv  ← (you supply this)
```

---

## REST API Reference

### Shipment APIs
| Method | Endpoint             | Description               |
|--------|----------------------|---------------------------|
| GET    | `/shipments`         | List all shipments        |
| GET    | `/shipments/:id`     | Get shipment by ID        |

### Delay APIs
| Method | Endpoint             | Description               |
|--------|----------------------|---------------------------|
| GET    | `/delays`            | All delayed shipments     |
| GET    | `/delays/high-risk`  | High-risk shipments only  |

### Analytics APIs
| Method | Endpoint                          | Description                  |
|--------|-----------------------------------|------------------------------|
| GET    | `/analytics/delay-rate`           | Overall delay rate           |
| GET    | `/analytics/top-delayed-routes`   | Top 20 delayed routes        |
| GET    | `/analytics/carrier-performance`  | Per-carrier delay stats      |
| GET    | `/analytics/avg-delivery-time`    | Avg delivery time by carrier |

### Tracking API
| Method | Endpoint                  | Description                      |
|--------|---------------------------|----------------------------------|
| GET    | `/track/:shipment_id`     | Customer-facing shipment status  |

---

## Risk Score Formula

```
risk_score = (weather_severity/10 × 0.4)
           + (traffic_condition/10 × 0.3)
           + (distance_km/5000     × 0.3)
```

| Score Range | Risk Level |
|-------------|------------|
| 0.0 – 0.3   | LOW        |
| 0.3 – 0.6   | MEDIUM     |
| > 0.6       | HIGH       |

---

## Testing

```powershell
# Unit + table tests
go test ./tests/...

# With coverage report
go test -cover ./tests/...

# Benchmarks
go test -bench=. ./tests/...
```

---

## Concurrency Architecture

```
CSV Loader → []Shipment
                 ↓
         WorkerPool (10 goroutines)
           ↓ goroutine 1 .. 10
     jobs channel (buffered)
           ↓
     DelayEngine.Process()
           ↓
     RiskEngine.Process()
           ↓
     ShipmentRepository.BulkLoad()
           ↓
        Gin REST API
```

`sync.WaitGroup` ensures all goroutines finish before the HTTP server starts.
`sync.RWMutex` protects concurrent reads of repository data once the API is live.

---

## Role Support

| Role     | Accessible Endpoints                                      |
|----------|-----------------------------------------------------------|
| Owner    | All endpoints                                             |
| Admin    | `/shipments/*`, `/delays/*`                               |
| Customer | `/track/:shipment_id` only                                |

> Authentication/authorization is not implemented in the backend. Enforce roles at the API gateway or frontend layer.
