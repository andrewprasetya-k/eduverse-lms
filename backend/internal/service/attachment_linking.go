package service

import (
	"backend/internal/domain"
	"fmt"
)

func replaceSourceAttachments(attService AttachmentService, schoolID string, sourceType domain.SourceType, sourceID string, mediaIDs []string) error {
	uniqueIDs, err := uniqueNonEmptyIDs(mediaIDs)
	if err != nil {
		return err
	}

	if err := attService.ReplaceBySource(schoolID, string(sourceType), sourceID, uniqueIDs); err != nil {
		return fmt.Errorf("failed to link media attachments")
	}
	return nil
}
