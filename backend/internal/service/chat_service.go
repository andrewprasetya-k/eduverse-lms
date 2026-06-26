package service

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/repository"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	chatRoomTypeSubject = "subject"
	chatRefTypeSubject  = "subject"
	chatMessageTypeText = "text"
	maxChatMessageLimit = 50
	maxChatContentLen   = 5000
)

type ChatService interface {
	ListMyRooms(userID string, schoolID string) ([]dto.ChatRoomDTO, error)
	OpenSubjectClassRoom(userID string, schoolID string, subjectClassID string) (*dto.ChatRoomDTO, error)
	ListMessages(userID string, schoolID string, roomID string, limit int, before *time.Time) (*dto.ChatMessagesResponseDTO, error)
	CreateMessage(userID string, schoolID string, roomID string, content string) (*dto.ChatMessageDTO, error)
	MarkRead(userID string, schoolID string, roomID string, lastReadMessageID *string) error
	CanAccessSubjectClassChat(userID string, schoolID string, subjectClassID string) (bool, error)
	CanAccessRoom(userID string, schoolID string, roomID string) (bool, *repository.ChatRoomRow, error)
}

type chatService struct {
	repo             repository.ChatRepository
	subjectClassRepo repository.SubjectClassRepository
}

func NewChatService(repo repository.ChatRepository, subjectClassRepo repository.SubjectClassRepository) ChatService {
	return &chatService{
		repo:             repo,
		subjectClassRepo: subjectClassRepo,
	}
}

func (s *chatService) ListMyRooms(userID string, schoolID string) ([]dto.ChatRoomDTO, error) {
	studentRows, err := s.repo.ListStudentRooms(userID, schoolID)
	if err != nil {
		return nil, err
	}
	teacherRows, err := s.repo.ListTeacherRooms(userID, schoolID)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool, len(studentRows)+len(teacherRows))
	rows := make([]repository.ChatRoomRow, 0, len(studentRows)+len(teacherRows))
	for _, row := range append(studentRows, teacherRows...) {
		if row.RoomID == "" || seen[row.RoomID] {
			continue
		}
		seen[row.RoomID] = true
		rows = append(rows, row)
	}

	rooms := make([]dto.ChatRoomDTO, 0, len(rows))
	for _, row := range rows {
		unread, err := s.repo.UnreadCount(row.RoomID, userID)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, mapChatRoomRow(row, unread))
	}
	return rooms, nil
}

func (s *chatService) OpenSubjectClassRoom(userID string, schoolID string, subjectClassID string) (*dto.ChatRoomDTO, error) {
	allowed, err := s.CanAccessSubjectClassChat(userID, schoolID, subjectClassID)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, fmt.Errorf("forbidden: chat subject class access denied")
	}

	context, err := s.repo.SubjectClassContext(subjectClassID, schoolID)
	if err != nil {
		return nil, err
	}

	room, err := s.repo.GetRoomBySubjectClass(schoolID, subjectClassID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		room = &domain.ChatRoom{
			SchoolID:  schoolID,
			Name:      context.RoomName,
			Type:      chatRoomTypeSubject,
			RefType:   chatRefTypeSubject,
			RefID:     subjectClassID,
			CreatedBy: userID,
		}
		if err := s.repo.CreateSubjectClassRoom(room); err != nil {
			return nil, err
		}
		room, err = s.repo.GetRoomBySubjectClass(schoolID, subjectClassID)
		if err != nil {
			return nil, err
		}
		_ = s.repo.UpdateSubjectClassRoomID(subjectClassID, room.ID)
	}

	context.RoomID = room.ID
	context.RoomName = room.Name
	unread, err := s.repo.UnreadCount(room.ID, userID)
	if err != nil {
		return nil, err
	}
	roomDTO := mapChatRoomRow(*context, unread)
	return &roomDTO, nil
}

func (s *chatService) ListMessages(userID string, schoolID string, roomID string, limit int, before *time.Time) (*dto.ChatMessagesResponseDTO, error) {
	allowed, _, err := s.CanAccessRoom(userID, schoolID, roomID)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, fmt.Errorf("forbidden: chat room access denied")
	}

	if limit <= 0 || limit > maxChatMessageLimit {
		limit = maxChatMessageLimit
	}
	rows, err := s.repo.ListMessages(roomID, limit+1, before)
	if err != nil {
		return nil, err
	}

	hasMore := len(rows) > limit
	if hasMore {
		rows = rows[1:]
	}

	messages := make([]dto.ChatMessageDTO, 0, len(rows))
	for _, row := range rows {
		messages = append(messages, mapChatMessageRow(row, userID))
	}

	var nextBefore *string
	if hasMore && len(rows) > 0 {
		value := formatChatTime(rows[0].CreatedAt)
		nextBefore = &value
	}

	return &dto.ChatMessagesResponseDTO{
		Messages:   messages,
		NextBefore: nextBefore,
		HasMore:    hasMore,
	}, nil
}

func (s *chatService) CreateMessage(userID string, schoolID string, roomID string, content string) (*dto.ChatMessageDTO, error) {
	allowed, _, err := s.CanAccessRoom(userID, schoolID, roomID)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, fmt.Errorf("forbidden: chat room access denied")
	}

	content = strings.TrimSpace(content)
	if content == "" {
		return nil, fmt.Errorf("chat message content is required")
	}
	if len([]rune(content)) > maxChatContentLen {
		return nil, fmt.Errorf("chat message content exceeds %d characters", maxChatContentLen)
	}

	message := domain.ChatMessage{
		RoomID:  roomID,
		UserID:  userID,
		Content: content,
		Type:    chatMessageTypeText,
	}
	if err := s.repo.CreateMessage(&message); err != nil {
		return nil, err
	}

	row, err := s.repo.GetMessageByID(message.ID, roomID)
	if err != nil {
		return nil, err
	}
	mapped := mapChatMessageRow(*row, userID)
	return &mapped, nil
}

func (s *chatService) MarkRead(userID string, schoolID string, roomID string, lastReadMessageID *string) error {
	allowed, _, err := s.CanAccessRoom(userID, schoolID, roomID)
	if err != nil {
		return err
	}
	if !allowed {
		return fmt.Errorf("forbidden: chat room access denied")
	}

	if lastReadMessageID != nil && *lastReadMessageID != "" {
		if _, err := s.repo.GetMessageByID(*lastReadMessageID, roomID); err != nil {
			return err
		}
	}
	return s.repo.UpsertReadReceipt(roomID, userID, lastReadMessageID)
}

func (s *chatService) CanAccessSubjectClassChat(userID string, schoolID string, subjectClassID string) (bool, error) {
	if userID == "" || schoolID == "" || subjectClassID == "" {
		return false, nil
	}

	studentOK, err := s.subjectClassRepo.UserEnrolledInSubjectClassAsRole(userID, schoolID, subjectClassID, "student")
	if err != nil {
		return false, err
	}
	if studentOK {
		return true, nil
	}

	teacherOK, err := s.subjectClassRepo.TeacherOwnsSubjectClass(userID, schoolID, subjectClassID)
	if err != nil {
		return false, err
	}
	return teacherOK, nil
}

func (s *chatService) CanAccessRoom(userID string, schoolID string, roomID string) (bool, *repository.ChatRoomRow, error) {
	room, err := s.repo.GetRoomContext(roomID, schoolID)
	if err != nil {
		return false, nil, err
	}
	allowed, err := s.CanAccessSubjectClassChat(userID, schoolID, room.SubjectClassID)
	if err != nil {
		return false, nil, err
	}
	return allowed, room, nil
}

func mapChatRoomRow(row repository.ChatRoomRow, unread int64) dto.ChatRoomDTO {
	var lastMessage *dto.ChatLastMessageDTO
	var lastMessageAt *string
	if row.LastMessageID != nil && row.LastContent != nil && row.LastMessageAt != nil {
		createdAt := formatChatTime(*row.LastMessageAt)
		lastMessageAt = &createdAt
		lastMessage = &dto.ChatLastMessageDTO{
			MessageID: *row.LastMessageID,
			Content:   *row.LastContent,
			CreatedAt: createdAt,
		}
		if row.LastSenderID != nil {
			lastMessage.SenderID = *row.LastSenderID
		}
		if row.LastSenderName != nil {
			lastMessage.SenderName = *row.LastSenderName
		}
	}

	return dto.ChatRoomDTO{
		RoomID:         row.RoomID,
		SubjectClassID: row.SubjectClassID,
		SubjectID:      row.SubjectID,
		SubjectName:    row.SubjectName,
		SubjectCode:    row.SubjectCode,
		ClassID:        row.ClassID,
		ClassName:      row.ClassName,
		ClassCode:      row.ClassCode,
		RoomName:       row.RoomName,
		LastMessage:    lastMessage,
		LastMessageAt:  lastMessageAt,
		UnreadCount:    unread,
		CanSend:        true,
	}
}

func mapChatMessageRow(row repository.ChatMessageRow, currentUserID string) dto.ChatMessageDTO {
	return dto.ChatMessageDTO{
		MessageID:   row.MessageID,
		RoomID:      row.RoomID,
		SenderID:    row.SenderID,
		SenderName:  row.SenderName,
		SenderRole:  row.SenderRole,
		Content:     row.Content,
		MessageType: row.Type,
		CreatedAt:   formatChatTime(row.CreatedAt),
		IsMine:      row.SenderID == currentUserID,
	}
}

func formatChatTime(value time.Time) string {
	return value.UTC().Format(time.RFC3339)
}
