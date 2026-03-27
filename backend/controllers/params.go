// Package controllers – pagination query-param helper shared by all controllers.
package controllers

import (
	"fmt"
	"strconv"
	"supply-chain-monitor/models"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage  = 1
	defaultLimit = 50
	maxLimit     = 100
)

// parsePageParams reads ?page and ?limit from the Gin context, applies defaults
// and enforces the maximum limit.  Returns (PageParams, nil) or ("", HTTP 400
// error) which the caller should write back with ctx.JSON.
func parsePageParams(ctx *gin.Context) (models.PageParams, error) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "50")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return models.PageParams{}, fmt.Errorf("invalid page value %q: must be a positive integer", pageStr)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return models.PageParams{}, fmt.Errorf("invalid limit value %q: must be a positive integer", limitStr)
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	return models.PageParams{Page: page, Limit: limit}, nil
}

// parseShipmentFilter reads optional ?origin, ?destination, ?carrier, ?mode,
// ?status query parameters from the Gin context.
func parseShipmentFilter(ctx *gin.Context) models.ShipmentFilter {
	return models.ShipmentFilter{
		Origin:      ctx.Query("origin"),
		Destination: ctx.Query("destination"),
		Carrier:     ctx.Query("carrier"),
		Mode:        ctx.Query("mode"),
		Status:      ctx.Query("status"),
	}
}
