// Package routes wires together all controllers and registers them with Gin.
package routes

import (
	"supply-chain-monitor/controllers"
	"supply-chain-monitor/middleware"

	"github.com/gin-gonic/gin"
)

// AppControllers groups all controllers
type AppControllers struct {
	Shipment  *controllers.ShipmentController
	Analytics *controllers.AnalyticsController
	Tracking  *controllers.TrackingController
	Dataset   *controllers.DatasetController
	Auth      *controllers.AuthController // NEW
}

// RegisterRoutes maps every API endpoint
func RegisterRoutes(r *gin.Engine, ctrl AppControllers) {

	// ── Auth Routes (Public — no token needed) ──────────────────────────────
	r.POST("/register", ctrl.Auth.Register)
	r.POST("/login", ctrl.Auth.Login)

	// ── Dataset Upload API (Protected) ─────────────────────────────────────
	r.POST("/upload-dataset", middleware.AuthMiddleware(), ctrl.Dataset.UploadDataset)

	// ── Shipment APIs (Protected) ───────────────────────────────────────────
	shipments := r.Group("/shipments", middleware.AuthMiddleware())
	{
		shipments.GET("", ctrl.Shipment.GetAll)
		shipments.GET("/:id", ctrl.Shipment.GetByID)
	}

	// ── Delay APIs (Protected) ──────────────────────────────────────────────
	delays := r.Group("/delays", middleware.AuthMiddleware())
	{
		delays.GET("", ctrl.Analytics.GetDelays)
		delays.GET("/high-risk", ctrl.Analytics.GetHighRisk)
	}

	// ── Analytics APIs (Protected) ──────────────────────────────────────────
	analytics := r.Group("/analytics", middleware.AuthMiddleware())
	{
		analytics.GET("/delay-rate", ctrl.Analytics.DelayRate)
		analytics.GET("/top-delayed-routes", ctrl.Analytics.TopDelayedRoutes)
		analytics.GET("/carrier-performance", ctrl.Analytics.CarrierPerformance)
		analytics.GET("/avg-delivery-time", ctrl.Analytics.AvgDeliveryTime)
	}

	// ── Tracking API (Protected) ────────────────────────────────────────────
	r.GET("/track/:shipment_id", middleware.AuthMiddleware(), ctrl.Tracking.Track)
}
