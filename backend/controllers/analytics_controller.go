// Package controllers – analytics controller.
package controllers

import (
	"net/http"
	"supply-chain-monitor/services"

	"github.com/gin-gonic/gin"
)

// AnalyticsController handles analytics and delay API endpoints.
type AnalyticsController struct {
	analyticsSvc *services.AnalyticsService
	delaySvc     *services.DelayService
}

// NewAnalyticsController creates an AnalyticsController.
func NewAnalyticsController(analyticsSvc *services.AnalyticsService, delaySvc *services.DelayService) *AnalyticsController {
	return &AnalyticsController{analyticsSvc: analyticsSvc, delaySvc: delaySvc}
}

// DelayRate handles GET /analytics/delay-rate
func (a *AnalyticsController) DelayRate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, a.analyticsSvc.DelayRate())
}

// TopDelayedRoutes handles GET /analytics/top-delayed-routes
func (a *AnalyticsController) TopDelayedRoutes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"routes": a.analyticsSvc.TopDelayedRoutes(20),
	})
}

// CarrierPerformance handles GET /analytics/carrier-performance
func (a *AnalyticsController) CarrierPerformance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"carriers": a.analyticsSvc.CarrierPerformance(),
	})
}

// AvgDeliveryTime handles GET /analytics/avg-delivery-time
func (a *AnalyticsController) AvgDeliveryTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"avg_delivery_times": a.analyticsSvc.AvgDeliveryTime(),
	})
}

// GetDelays handles GET /delays
//
// Supported query parameters:
//
//	page   – page number (default 1)
//	limit  – records per page (default 50, max 100)
//	origin, destination, carrier, mode – optional filters
//
// Example: GET /delays?carrier=BlueDart&page=1&limit=25
func (a *AnalyticsController) GetDelays(ctx *gin.Context) {
	pp, err := parsePageParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := parseShipmentFilter(ctx)

	resp, err := a.delaySvc.GetDelayedPaged(filter, pp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query delayed shipments"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetHighRisk handles GET /delays/high-risk
//
// Supported query parameters:
//
//	page  – page number (default 1)
//	limit – records per page (default 50, max 100)
//
// Example: GET /delays/high-risk?page=2&limit=100
func (a *AnalyticsController) GetHighRisk(ctx *gin.Context) {
	pp, err := parsePageParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.delaySvc.GetHighRiskPaged(pp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query high-risk shipments"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
