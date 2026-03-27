// Package controllers – shipment controller handling GET /shipments and GET /shipments/:id.
package controllers

import (
	"net/http"
	"supply-chain-monitor/services"

	"github.com/gin-gonic/gin"
)

// ShipmentController handles HTTP requests for shipment data.
type ShipmentController struct {
	svc *services.ShipmentService
}

// NewShipmentController creates a ShipmentController using the given service.
func NewShipmentController(svc *services.ShipmentService) *ShipmentController {
	return &ShipmentController{svc: svc}
}

// GetAll handles GET /shipments
//
// Supported query parameters:
//
//	page        – page number (default 1)
//	limit       – records per page (default 50, max 100)
//	origin      – filter by origin city (case-insensitive)
//	destination – filter by destination city
//	carrier     – filter by carrier name
//	mode        – filter by transport mode
//	status      – "delayed" | "ontime" (empty = all)
//
// Example: GET /shipments?origin=Delhi&status=delayed&page=2&limit=25
func (c *ShipmentController) GetAll(ctx *gin.Context) {
	pp, err := parsePageParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := parseShipmentFilter(ctx)

	resp, err := c.svc.ListFiltered(filter, pp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query shipments"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetByID handles GET /shipments/:id
func (c *ShipmentController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	shipment, ok := c.svc.GetByID(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "shipment not found", "shipment_id": id})
		return
	}
	ctx.JSON(http.StatusOK, shipment)
}
