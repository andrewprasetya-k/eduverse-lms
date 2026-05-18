package storage

import (
	"fmt"
	"net/url"
	"strings"
)

// ObjectPathValidator validates and sanitizes object paths for storage operations
type ObjectPathValidator struct {
	maxPathLength int
}

// NewObjectPathValidator creates a new path validator
func NewObjectPathValidator(maxPathLength int) *ObjectPathValidator {
	if maxPathLength <= 0 {
		maxPathLength = 512 // Default max path length
	}
	return &ObjectPathValidator{
		maxPathLength: maxPathLength,
	}
}

// Validate checks if objectPath is safe for use in storage operations
// Returns error if path is invalid or unsafe
func (v *ObjectPathValidator) Validate(objectPath string) error {
	// Check empty
	if strings.TrimSpace(objectPath) == "" {
		return ErrInvalidPath
	}

	// Check length
	if len(objectPath) > v.maxPathLength {
		return fmt.Errorf("path exceeds maximum length of %d characters", v.maxPathLength)
	}

	// Check for absolute path
	if strings.HasPrefix(objectPath, "/") {
		return fmt.Errorf("path cannot be absolute (starts with /)")
	}

	// Check for directory traversal
	if strings.Contains(objectPath, "..") {
		return fmt.Errorf("path cannot contain .. (directory traversal)")
	}

	// Check for double slashes
	if strings.Contains(objectPath, "//") {
		return fmt.Errorf("path cannot contain // (double slashes)")
	}

	// Check for backslashes (not allowed in cloud storage paths)
	if strings.Contains(objectPath, "\\") {
		return fmt.Errorf("path cannot contain backslashes")
	}

	// Check for query/fragment characters
	disallowedChars := []string{"?", "#", "&", "=", "%"}
	for _, char := range disallowedChars {
		if strings.Contains(objectPath, char) {
			return fmt.Errorf("path contains disallowed character: %s", char)
		}
	}

	return nil
}

// SafeURL encodes path segments to prevent injection when building URLs
// Split by /, encode each segment, rejoin
func (v *ObjectPathValidator) SafeURL(objectPath string) string {
	// Split path into segments
	segments := strings.Split(objectPath, "/")

	// URL-encode each segment
	encoded := make([]string, len(segments))
	for i, segment := range segments {
		encoded[i] = url.QueryEscape(segment)
	}

	// Rejoin with forward slash
	return strings.Join(encoded, "/")
}
