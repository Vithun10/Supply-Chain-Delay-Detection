// Package utils provides helper utilities for the supply-chain monitor.
package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"supply-chain-monitor/models"
	"supply-chain-monitor/utils/internal/dateparser"
)

// LoadCSV reads the supply-chain CSV dataset at filePath and returns a slice of
// Shipment records.  Any row that cannot be parsed is skipped with a warning.
func LoadCSV(filePath string) ([]models.Shipment, error) {
	Logger.Printf("Loading dataset from %s", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("csv_loader: open file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Read and discard the header row
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("csv_loader: read header: %w", err)
	}
	idx := buildColumnIndex(header)

	var shipments []models.Shipment
	lineNum := 1

	for {
		lineNum++
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Logger.Printf("csv_loader: skip malformed line %d: %v", lineNum, err)
			continue
		}

		s, parseErr := parseRow(row, idx)
		if parseErr != nil {
			Logger.Printf("csv_loader: skip line %d: %v", lineNum, parseErr)
			continue
		}
		shipments = append(shipments, s)
	}

	Logger.Printf("Loaded %d shipment records", len(shipments))
	return shipments, nil
}

// columnIndex maps lowercase column names to their position in a CSV row.
type columnIndex map[string]int

func buildColumnIndex(header []string) columnIndex {
	idx := make(columnIndex, len(header))
	for i, name := range header {
		idx[name] = i
	}
	return idx
}

// get returns the value at column name or "" if missing.
func (idx columnIndex) get(row []string, col string) string {
	i, ok := idx[col]
	if !ok || i >= len(row) {
		return ""
	}
	return row[i]
}

// parseRow converts a CSV row into a Shipment using column positions from idx.
func parseRow(row []string, idx columnIndex) (models.Shipment, error) {
	// Required string fields
	id := idx.get(row, "shipment_id")
	if id == "" {
		return models.Shipment{}, fmt.Errorf("missing shipment_id")
	}

	distKM, _ := strconv.ParseFloat(idx.get(row, "distance_km"), 64)
	weatherSev, _ := strconv.ParseFloat(idx.get(row, "weather_severity"), 64)
	trafficCond, _ := strconv.ParseFloat(idx.get(row, "traffic_condition"), 64)

	expectedDate, err := dateparser.Parse(idx.get(row, "expected_delivery_date"))
	if err != nil {
		return models.Shipment{}, fmt.Errorf("parse expected_delivery_date: %w", err)
	}

	deliveredDate, err := dateparser.Parse(idx.get(row, "delivered_date"))
	if err != nil {
		return models.Shipment{}, fmt.Errorf("parse delivered_date: %w", err)
	}

	return models.Shipment{
		ShipmentID:           id,
		Origin:               idx.get(row, "origin"),
		Destination:          idx.get(row, "destination"),
		DistanceKM:           distKM,
		Carrier:              idx.get(row, "carrier"),
		Mode:                 idx.get(row, "mode"),
		ExpectedDeliveryDate: expectedDate,
		DeliveredDate:        deliveredDate,
		WeatherSeverity:      weatherSev,
		TrafficCondition:     trafficCond,
	}, nil
}
