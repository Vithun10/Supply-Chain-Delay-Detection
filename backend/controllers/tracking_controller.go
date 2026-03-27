// Package controllers – tracking controller for GET /track/:shipment_id.
package controllers

import (
	"net/http"
	"supply-chain-monitor/services"

	"github.com/gin-gonic/gin"
)

// TrackingController handles the customer-facing tracking endpoint.
type TrackingController struct {
	svc *services.ShipmentService
}

// NewTrackingController creates a TrackingController.
func NewTrackingController(svc *services.ShipmentService) *TrackingController {
	return &TrackingController{svc: svc}
}

// Track handles GET /track/:shipment_id
// It returns a simplified tracking view suitable for the customer role.
func (t *TrackingController) Track(ctx *gin.Context) {
	id := ctx.Param("shipment_id")
	shipment, ok := t.svc.GetByID(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "shipment not found", "shipment_id": id})
		return
	}

	// Customer view – only expose tracking-relevant fields.
	ctx.JSON(http.StatusOK, gin.H{
		"shipment_id":            shipment.ShipmentID,
		"origin":                 shipment.Origin,
		"destination":            shipment.Destination,
		"carrier":                shipment.Carrier,
		"mode":                   shipment.Mode,
		"delay_detected":         shipment.DelayDetected,
		"delay_days":             shipment.DelayDays,
		"risk_level":             shipment.RiskLevel,
		"expected_delivery_date": shipment.ExpectedDeliveryDate,
		"delivered_date":         shipment.DeliveredDate,
	})
}
