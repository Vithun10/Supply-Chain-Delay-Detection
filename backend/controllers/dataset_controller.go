// Package controllers handles HTTP requests for dataset uploads.
package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"supply-chain-monitor/concurrency"
	"supply-chain-monitor/repository"
	"supply-chain-monitor/utils"
)

// DatasetController manages dataset upload and processing
type DatasetController struct {
	Repo *repository.ShipmentRepository
}

// NewDatasetController creates a new DatasetController
func NewDatasetController(repo *repository.ShipmentRepository) *DatasetController {
	return &DatasetController{
		Repo: repo,
	}
}

// UploadDataset handles CSV dataset upload
func (dc *DatasetController) UploadDataset(c *gin.Context) {

	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CSV file is required",
		})
		return
	}

	// Save file to data folder
	savePath := filepath.Join("data", file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save uploaded dataset",
		})
		return
	}

	utils.Logger.Printf("New dataset uploaded: %s", savePath)

	// Load CSV using existing loader
	shipments, err := utils.LoadCSV(savePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dataset format",
		})
		return
	}

	utils.Logger.Printf("Loaded %d shipments from uploaded dataset", len(shipments))

	// Process shipments using worker pool
	pool := concurrency.NewWorkerPool(10)
	processed := pool.Process(shipments)

	// Replace repository data
	dc.Repo.Clear()
	dc.Repo.BulkLoad(processed)

	utils.Logger.Printf("Repository updated with %d new shipments", dc.Repo.Count())

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message":           "Dataset uploaded and processed successfully",
		"processed_records": len(processed),
	})
}
