package handler

import (
	"backend/internal/dto"
	"backend/internal/middleware"
	"backend/internal/realtime"
	"backend/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	service service.ChatService
	hub     *realtime.Hub
}

func NewChatHandler(service service.ChatService, hub *realtime.Hub) *ChatHandler {
	return &ChatHandler{service: service, hub: hub}
}

func (h *ChatHandler) ListRooms(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	rooms, err := h.service.ListMyRooms(userID, schoolID, c.Query("search"))
	if err != nil {
		HandleError(c, err)
		return
	}
	if rooms == nil {
		rooms = make([]dto.ChatRoomDTO, 0)
	}
	c.JSON(http.StatusOK, dto.ChatRoomsResponseDTO{Rooms: rooms})
}

func (h *ChatHandler) ListMembers(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var excludeRoomID *string
	if value := c.Query("excludeRoomId"); value != "" {
		excludeRoomID = &value
	}
	members, err := h.service.ListMembers(userID, schoolID, c.Query("search"), excludeRoomID)
	if err != nil {
		HandleError(c, err)
		return
	}
	if members == nil {
		members = make([]dto.ChatMemberDTO, 0)
	}
	c.JSON(http.StatusOK, dto.ChatMembersResponseDTO{Members: members})
}

func (h *ChatHandler) OpenSchoolRoom(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	room, err := h.service.OpenSchoolRoom(userID, schoolID)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ChatRoomResponseDTO{Room: *room})
}

func (h *ChatHandler) OpenDirectMessage(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.OpenDirectMessageDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	room, err := h.service.OpenDirectMessage(userID, schoolID, input.TargetUserID)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ChatRoomResponseDTO{Room: *room})
}

func (h *ChatHandler) CreateGroupRoom(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.CreateChatGroupDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	room, err := h.service.CreateGroupRoom(userID, schoolID, input.RoomName, input.MemberUserIDs)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.ChatRoomResponseDTO{Room: *room})
}

func (h *ChatHandler) GetGroupInfo(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	group, err := h.service.GetGroupInfo(userID, schoolID, c.Param("roomId"))
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ChatGroupInfoResponseDTO{Group: *group})
}

func (h *ChatHandler) RenameGroupRoom(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.UpdateChatGroupDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	room, err := h.service.RenameGroupRoom(userID, schoolID, c.Param("roomId"), input.RoomName)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.ChatRoomResponseDTO{Room: *room})
}

func (h *ChatHandler) LeaveGroupRoom(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	if err := h.service.LeaveGroupRoom(userID, schoolID, c.Param("roomId")); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Keluar dari grup berhasil"})
}

func (h *ChatHandler) AddGroupMembers(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.AddChatGroupMembersDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	if err := h.service.AddGroupMembers(userID, schoolID, c.Param("roomId"), input.MemberUserIDs); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Anggota grup ditambahkan"})
}

func (h *ChatHandler) RemoveGroupMember(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	if err := h.service.RemoveGroupMember(userID, schoolID, c.Param("roomId"), c.Param("userId")); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Anggota grup dikeluarkan"})
}

func (h *ChatHandler) ListMessages(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	var before *time.Time
	if raw := c.Query("before"); raw != "" {
		parsed, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid before timestamp"})
			return
		}
		before = &parsed
	}

	response, err := h.service.ListMessages(userID, schoolID, c.Param("roomId"), limit, before)
	if err != nil {
		HandleError(c, err)
		return
	}
	if response.Messages == nil {
		response.Messages = make([]dto.ChatMessageDTO, 0)
	}
	c.JSON(http.StatusOK, response)
}

func (h *ChatHandler) CreateMessage(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.CreateChatMessageDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	message, err := h.service.CreateMessage(userID, schoolID, c.Param("roomId"), input.Content)
	if err != nil {
		HandleError(c, err)
		return
	}
	h.broadcastNewMessage(userID, schoolID, c.Param("roomId"), *message)
	h.broadcastRoomUpdated(userID, schoolID, c.Param("roomId"), "new_message")
	c.JSON(http.StatusCreated, message)
}

func (h *ChatHandler) broadcastNewMessage(userID string, schoolID string, roomID string, message dto.ChatMessageDTO) {
	if h.hub == nil {
		return
	}
	recipients, err := h.service.ListRealtimeRecipients(userID, schoolID, roomID)
	if err != nil {
		return
	}
	for _, recipientID := range recipients {
		payload := message
		payload.IsMine = recipientID == message.SenderID
		h.hub.BroadcastToUser(schoolID, recipientID, realtime.Event{
			Type:     realtime.EventTypeNewMessage,
			RoomID:   roomID,
			SchoolID: schoolID,
			Payload:  payload,
		})
	}
}

func (h *ChatHandler) broadcastMessageRead(userID string, schoolID string, roomID string, receipt dto.ChatReadReceiptDTO) {
	if h.hub == nil {
		return
	}
	recipients, err := h.service.ListRealtimeRecipients(userID, schoolID, roomID)
	if err != nil {
		return
	}
	h.hub.BroadcastToUsers(schoolID, recipients, realtime.Event{
		Type:     realtime.EventTypeMessageRead,
		RoomID:   roomID,
		SchoolID: schoolID,
		Payload:  receipt,
	})
}

func (h *ChatHandler) broadcastRoomUpdated(userID string, schoolID string, roomID string, reason string) {
	if h.hub == nil {
		return
	}
	recipients, err := h.service.ListRealtimeRecipients(userID, schoolID, roomID)
	if err != nil {
		return
	}
	h.hub.BroadcastToUsers(schoolID, recipients, realtime.Event{
		Type:     realtime.EventTypeRoomUpdated,
		RoomID:   roomID,
		SchoolID: schoolID,
		Payload: gin.H{
			"reason": reason,
		},
	})
}

func (h *ChatHandler) GetReadSummary(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	summary, err := h.service.GetReadSummary(userID, schoolID, c.Param("roomId"))
	if err != nil {
		HandleError(c, err)
		return
	}
	if summary.Members == nil {
		summary.Members = make([]dto.ChatReadMemberDTO, 0)
	}
	c.JSON(http.StatusOK, summary)
}

func (h *ChatHandler) MarkRead(c *gin.Context) {
	userID := middleware.GetUserID(c)
	schoolID, ok := getChatActiveSchoolID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	var input dto.MarkChatRoomReadDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	receipt, err := h.service.MarkRead(userID, schoolID, c.Param("roomId"), input.LastReadMessageID)
	if err != nil {
		HandleError(c, err)
		return
	}
	h.broadcastMessageRead(userID, schoolID, c.Param("roomId"), *receipt)
	h.broadcastRoomUpdated(userID, schoolID, c.Param("roomId"), "message_read")
	c.JSON(http.StatusOK, dto.MarkChatRoomReadResponseDTO{
		Message:           "Chat room marked as read",
		RoomID:            receipt.RoomID,
		UserID:            receipt.UserID,
		LastReadMessageID: receipt.LastReadMessageID,
		LastReadAt:        receipt.LastReadAt,
	})
}

func getChatActiveSchoolID(c *gin.Context) (string, bool) {
	if sid, exists := c.Get("school_id"); exists {
		if value, ok := sid.(string); ok && value != "" {
			return value, true
		}
	}
	if value := c.GetHeader("SchoolId"); value != "" {
		return value, true
	}
	return "", false
}
