package dto

type LogResponseDTO struct {
	ID        string `json:"logId"`
	UserID    string `json:"userId"`
	UserName  string `json:"userName,omitempty"`
	Action    string `json:"action"`
	Metadata  string `json:"metadata"`
	CreatedAt string `json:"createdAt"`
}
