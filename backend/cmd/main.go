// main.go is the application entry point for the supply-chain monitor server.
package main

import (
	"time"

	"supply-chain-monitor/concurrency"
	"supply-chain-monitor/config"
	"supply-chain-monitor/controllers"
	"supply-chain-monitor/middleware"
	"supply-chain-monitor/repository"
	"supply-chain-monitor/routes"
	"supply-chain-monitor/services"
	"supply-chain-monitor/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.DefaultConfig()

	// ─────────────────────────────────────────────────────────────
	// 1. Load dataset from CSV
	// ─────────────────────────────────────────────────────────────
	shipments, err := utils.LoadCSV(cfg.DataFilePath)
	if err != nil {
		utils.Logger.Fatalf("Failed to load dataset: %v", err)
	}

	// ─────────────────────────────────────────────────────────────
	// 2. Process shipments concurrently
	// ─────────────────────────────────────────────────────────────
	pool := concurrency.NewWorkerPool(cfg.WorkerCount)
	processed := pool.Process(shipments)

	// ─────────────────────────────────────────────────────────────
	// 3. Store results in repository
	// ─────────────────────────────────────────────────────────────
	repo := repository.NewShipmentRepository()
	repo.BulkLoad(processed)
	utils.Logger.Printf("Repository loaded with %d shipments", repo.Count())

	// ─────────────────────────────────────────────────────────────
	// 4. Initialize user repository (NEW)
	// ─────────────────────────────────────────────────────────────
	userRepo := repository.NewUserRepository()

	// ─────────────────────────────────────────────────────────────
	// 5. Initialize services
	// ─────────────────────────────────────────────────────────────
	shipmentSvc := services.NewShipmentService(repo)
	delaySvc := services.NewDelayService(repo)
	analyticsSvc := services.NewAnalyticsService(repo)
	authSvc := services.NewAuthService(userRepo) // NEW

	// ─────────────────────────────────────────────────────────────
	// 6. Initialize controllers
	// ─────────────────────────────────────────────────────────────
	shipmentCtrl := controllers.NewShipmentController(shipmentSvc)
	analyticsCtrl := controllers.NewAnalyticsController(analyticsSvc, delaySvc)
	trackingCtrl := controllers.NewTrackingController(shipmentSvc)
	datasetCtrl := controllers.NewDatasetController(repo)
	authCtrl := controllers.NewAuthController(authSvc) // NEW

	// ─────────────────────────────────────────────────────────────
	// 7. Start async background worker (Goroutine) NEW
	// ─────────────────────────────────────────────────────────────
	go func() {
		for {
			utils.Logger.Printf("[Async Worker] System running — shipments in repo: %d", repo.Count())
			time.Sleep(60 * time.Second)
		}
	}()

	// ─────────────────────────────────────────────────────────────
	// 8. Create Gin server
	// ─────────────────────────────────────────────────────────────
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ─────────────────────────────────────────────────────────────
	// 9. Root API endpoint
	// ─────────────────────────────────────────────────────────────
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "Supply Chain Delay Detection System",
			"version": "1.0",
			"status":  "Running",
		})
	})

	// ─────────────────────────────────────────────────────────────
	// 10. Register routes
	// ─────────────────────────────────────────────────────────────
	routes.RegisterRoutes(r, routes.AppControllers{
		Shipment:  shipmentCtrl,
		Analytics: analyticsCtrl,
		Tracking:  trackingCtrl,
		Dataset:   datasetCtrl,
		Auth:      authCtrl, // NEW
	})

	// ─────────────────────────────────────────────────────────────
	// 11. Start server
	// ─────────────────────────────────────────────────────────────
	utils.Logger.Printf("Server starting on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		utils.Logger.Fatalf("Server failed: %v", err)
	}

	_ = middleware.AuthMiddleware // ensure middleware package is referenced
}
