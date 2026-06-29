package realtime

import (
	"backend/internal/service"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	hub         *Hub
	chatService service.ChatService
	upgrader    websocket.Upgrader
}

func NewWebSocketHandler(hub *Hub, chatService service.ChatService) *WebSocketHandler {
	return &WebSocketHandler{
		hub:         hub,
		chatService: chatService,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(*http.Request) bool {
				return true
			},
		},
	}
}

func (h *WebSocketHandler) Chat(c *gin.Context) {
	tokenValue := extractHandshakeToken(c)
	if tokenValue == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := parseUserIDFromToken(tokenValue)
	if err != nil || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	schoolID := strings.TrimSpace(c.Query("schoolId"))
	if schoolID == "" {
		schoolID = strings.TrimSpace(c.GetHeader("SchoolId"))
	}
	if schoolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
		return
	}

	allowed, err := h.chatService.CanAccessSchoolChat(userID, schoolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify chat access"})
		return
	}
	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: chat school access denied"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := NewClient(h.hub, conn, userID, schoolID)
	h.hub.Register(client)
	client.ReadLoop()
}

func extractHandshakeToken(c *gin.Context) string {
	if value := strings.TrimSpace(c.Query("token")); value != "" {
		return value
	}
	authHeader := c.GetHeader("Authorization")
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func parseUserIDFromToken(tokenValue string) (string, error) {
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}
	userID, _ := claims["user_id"].(string)
	return userID, nil
}
