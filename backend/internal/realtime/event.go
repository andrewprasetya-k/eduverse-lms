package realtime

type Event struct {
	Type     string `json:"type"`
	RoomID   string `json:"roomId"`
	SchoolID string `json:"schoolId"`
	Payload  any    `json:"payload"`
}

const EventTypeNewMessage = "new_message"
