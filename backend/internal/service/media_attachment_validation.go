package service

import (
	"backend/internal/repository"
	"fmt"
	"strings"
)

func validateAttachableMedia(mediaRepo repository.MediaRepository, mediaIDs []string, schoolID string, userID string, isAdmin bool) error {
	_, err := prepareAttachableMediaIDs(mediaRepo, mediaIDs, schoolID, userID, isAdmin)
	return err
}

func prepareAttachableMediaIDs(mediaRepo repository.MediaRepository, mediaIDs []string, schoolID string, userID string, isAdmin bool) ([]string, error) {
	uniqueIDs, err := uniqueNonEmptyIDs(mediaIDs)
	if err != nil {
		return nil, err
	}
	if len(uniqueIDs) == 0 {
		return uniqueIDs, nil
	}

	medias, err := mediaRepo.GetByIDs(uniqueIDs)
	if err != nil {
		return nil, err
	}
	if len(medias) != len(uniqueIDs) {
		return nil, fmt.Errorf("invalid media attachment")
	}

	seen := make(map[string]bool, len(medias))
	for _, media := range medias {
		seen[media.ID] = true
		if media.SchoolID != schoolID {
			return nil, fmt.Errorf("forbidden: media does not belong to current school")
		}
		if !isAdmin && media.OwnerID != userID {
			return nil, fmt.Errorf("forbidden: media cannot be attached by current user")
		}
	}

	for _, id := range uniqueIDs {
		if !seen[id] {
			return nil, fmt.Errorf("invalid media attachment")
		}
	}

	return uniqueIDs, nil
}

func uniqueNonEmptyIDs(ids []string) ([]string, error) {
	seen := make(map[string]bool, len(ids))
	unique := make([]string, 0, len(ids))
	for _, id := range ids {
		trimmed := strings.TrimSpace(id)
		if trimmed == "" {
			return nil, fmt.Errorf("invalid media attachment")
		}
		if seen[trimmed] {
			continue
		}
		seen[trimmed] = true
		unique = append(unique, trimmed)
	}
	return unique, nil
}
