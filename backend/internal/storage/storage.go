package storage

import (
	"context"
	"io"
)

// Provider defines the interface for file storage operations
type Provider interface {
	// Upload stores a file and returns its public URL
	// objectPath: target path in storage (e.g., "schools/uuid/file.pdf")
	// content: file data to upload
	// contentType: MIME type of the file (e.g., "application/pdf")
	Upload(ctx context.Context, objectPath string, content io.Reader, contentType string) (publicURL string, err error)

	// Delete removes a file from storage
	// objectPath: path to the file to delete
	Delete(ctx context.Context, objectPath string) error

	// HealthCheck verifies storage is available
	HealthCheck(ctx context.Context) error
}
