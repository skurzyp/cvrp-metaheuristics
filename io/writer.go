// Package io provides utilities for reading and writing VRP data.
package io

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CSVWriter wraps a CSV file writer for structured result output.
type CSVWriter struct {
	file   *os.File
	writer *csv.Writer
}

// NewCSVWriter creates a new CSVWriter and opens/creates the given file path.
func NewCSVWriter(filePath string) *CSVWriter {
	// Ensure directory exists
	if err := os.MkdirAll(getDir(filePath), os.ModePerm); err != nil {
		fmt.Printf("Error creating directories for %s: %v\n", filePath, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	writer := csv.NewWriter(file)
	return &CSVWriter{
		file:   file,
		writer: writer,
	}
}

// Write writes a single record (row) to the CSV file.
func (w *CSVWriter) Write(record []string) {
	if err := w.writer.Write(record); err != nil {
		fmt.Printf("Error writing CSV record: %v\n", err)
	}
}

// Flush ensures all buffered data is written to disk.
func (w *CSVWriter) Flush() error {
	w.writer.Flush()
	if err := w.writer.Error(); err != nil {
		fmt.Printf("Error flushing CSV data: %v\n", err)
		return err
	}
	return nil
}

// Close flushes the writer and closes the underlying file.
func (w *CSVWriter) Close() {
	_ = w.Flush()
	if err := w.file.Close(); err != nil {
		fmt.Printf("Error closing CSV file: %v\n", err)
	}
}

// getDir extracts the directory path from a file path.
func getDir(filePath string) string {
	lastSlash := -1
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' || filePath[i] == '\\' {
			lastSlash = i
			break
		}
	}
	if lastSlash == -1 {
		return "."
	}
	return filePath[:lastSlash]
}
