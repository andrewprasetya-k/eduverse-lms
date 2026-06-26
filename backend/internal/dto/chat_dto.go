package dto

type ChatLastMessageDTO struct {
	MessageID  string `json:"messageId"`
	SenderID   string `json:"senderId"`
	SenderName string `json:"senderName"`
	Content    string `json:"content"`
	CreatedAt  string `json:"createdAt"`
}

type ChatRoomDTO struct {
	RoomID         string              `json:"roomId"`
	SubjectClassID string              `json:"subjectClassId"`
	SubjectID      string              `json:"subjectId"`
	SubjectName    string              `json:"subjectName"`
	SubjectCode    string              `json:"subjectCode"`
	ClassID        string              `json:"classId"`
	ClassName      string              `json:"className"`
	ClassCode      string              `json:"classCode"`
	RoomName       string              `json:"roomName"`
	LastMessage    *ChatLastMessageDTO `json:"lastMessage"`
	LastMessageAt  *string             `json:"lastMessageAt"`
	UnreadCount    int64               `json:"unreadCount"`
	CanSend        bool                `json:"canSend"`
}

type ChatMessageDTO struct {
	MessageID   string `json:"messageId"`
	RoomID      string `json:"roomId"`
	SenderID    string `json:"senderId"`
	SenderName  string `json:"senderName"`
	SenderRole  string `json:"senderRole"`
	Content     string `json:"content"`
	MessageType string `json:"messageType"`
	CreatedAt   string `json:"createdAt"`
	IsMine      bool   `json:"isMine"`
}

type ChatRoomsResponseDTO struct {
	Rooms []ChatRoomDTO `json:"rooms"`
}

type ChatRoomResponseDTO struct {
	Room ChatRoomDTO `json:"room"`
}

type ChatMessagesResponseDTO struct {
	Messages   []ChatMessageDTO `json:"messages"`
	NextBefore *string          `json:"nextBefore"`
	HasMore    bool             `json:"hasMore"`
}

type CreateChatMessageDTO struct {
	Content string `json:"content" binding:"required"`
}

type MarkChatRoomReadDTO struct {
	LastReadMessageID *string `json:"lastReadMessageId,omitempty" binding:"omitempty,uuid"`
}
