package repository

import (
	"backend/internal/domain"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatRepository interface {
	ListSchoolRooms(userID string, schoolID string) ([]ChatRoomRow, error)
	ListChatMembers(schoolID string, search string) ([]ChatMemberRow, error)
	GetSchoolRoom(schoolID string) (*domain.ChatRoom, error)
	CreateSchoolRoom(room *domain.ChatRoom) error
	GetRoomContext(roomID string, schoolID string) (*ChatRoomRow, error)
	UserIsActiveSchoolMember(userID string, schoolID string) (bool, error)
	UsersAreActiveSchoolMembers(userIDs []string, schoolID string) (map[string]bool, error)
	UserIsActiveRoomMember(userID string, roomID string) (bool, error)
	CreateGroupRoomWithMembers(schoolID string, roomName string, creatorID string, memberUserIDs []string) (*domain.ChatRoom, error)
	ListMessages(roomID string, limit int, before *time.Time) ([]ChatMessageRow, error)
	CreateMessage(message *domain.ChatMessage) error
	GetMessageByID(messageID string, roomID string) (*ChatMessageRow, error)
	UpsertReadReceipt(roomID string, userID string, messageID *string) error
	UnreadCount(roomID string, userID string) (int64, error)
}

type chatRepository struct {
	db *gorm.DB
}

type ChatRoomRow struct {
	RoomID         string     `gorm:"column:room_id"`
	RoomName       string     `gorm:"column:room_name"`
	RoomType       string     `gorm:"column:room_type"`
	RoomRefType    *string    `gorm:"column:room_ref_type"`
	RoomRefID      *string    `gorm:"column:room_ref_id"`
	SchoolID       string     `gorm:"column:school_id"`
	SchoolName     string     `gorm:"column:school_name"`
	LastMessageID  *string    `gorm:"column:last_message_id"`
	LastSenderID   *string    `gorm:"column:last_sender_id"`
	LastSenderName *string    `gorm:"column:last_sender_name"`
	LastContent    *string    `gorm:"column:last_content"`
	LastMessageAt  *time.Time `gorm:"column:last_message_at"`
}

type ChatMemberRow struct {
	UserID   string `gorm:"column:user_id"`
	FullName string `gorm:"column:full_name"`
	Email    string `gorm:"column:email"`
	Roles    string `gorm:"column:roles"`
}

type ChatMessageRow struct {
	MessageID  string    `gorm:"column:message_id"`
	RoomID     string    `gorm:"column:room_id"`
	SenderID   string    `gorm:"column:sender_id"`
	SenderName string    `gorm:"column:sender_name"`
	SenderRole string    `gorm:"column:sender_role"`
	Content    string    `gorm:"column:content"`
	Type       string    `gorm:"column:message_type"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) ListSchoolRooms(userID string, schoolID string) ([]ChatRoomRow, error) {
	var rows []ChatRoomRow
	err := r.db.Raw(chatRoomListSelect()+`
		JOIN edv.school_users scu
			ON scu.scu_usr_id = ?
			AND scu.scu_sch_id = cr.room_sch_id
			AND scu.deleted_at IS NULL
		JOIN edv.users active_user
			ON active_user.usr_id = scu.scu_usr_id
			AND active_user.deleted_at IS NULL
		WHERE cr.room_sch_id = ?
			AND cr.room_type = 'group'
			AND cr.deleted_at IS NULL
			AND s.deleted_at IS NULL
			AND (
				(cr.room_ref_type = 'school' AND cr.room_ref_id = ?)
				OR (
					cr.room_ref_type IS NULL
					AND cr.room_ref_id IS NULL
					AND EXISTS (
						SELECT 1
						FROM edv.chat_room_members crm
						WHERE crm.crm_room_id = cr.room_id
							AND crm.crm_usr_id = ?
							AND crm.left_at IS NULL
					)
				)
			)
		ORDER BY COALESCE(lm.created_at, cr.created_at) DESC
	`, userID, schoolID, schoolID, userID).Scan(&rows).Error
	return rows, err
}

func (r *chatRepository) ListChatMembers(schoolID string, search string) ([]ChatMemberRow, error) {
	var rows []ChatMemberRow
	searchPattern := "%" + search + "%"
	err := r.db.Raw(`
		SELECT
			u.usr_id AS user_id,
			COALESCE(u.usr_nama_lengkap, '') AS full_name,
			u.usr_email AS email,
			COALESCE(string_agg(DISTINCT rol.rol_name, ',' ORDER BY rol.rol_name), '') AS roles
		FROM edv.school_users scu
		JOIN edv.users u
			ON u.usr_id = scu.scu_usr_id
			AND u.deleted_at IS NULL
		LEFT JOIN edv.user_roles ur ON ur.urol_scu_id = scu.scu_id
		LEFT JOIN edv.roles rol ON rol.rol_id = ur.urol_rol_id
		WHERE scu.scu_sch_id = ?
			AND scu.deleted_at IS NULL
			AND (
				? = ''
				OR u.usr_nama_lengkap ILIKE ?
				OR u.usr_email ILIKE ?
			)
		GROUP BY u.usr_id, u.usr_nama_lengkap, u.usr_email
		ORDER BY u.usr_nama_lengkap ASC, u.usr_email ASC
		LIMIT 50
	`, schoolID, search, searchPattern, searchPattern).Scan(&rows).Error
	return rows, err
}

func (r *chatRepository) GetSchoolRoom(schoolID string) (*domain.ChatRoom, error) {
	var room domain.ChatRoom
	err := r.db.Where("room_sch_id = ? AND room_type = ? AND room_ref_type = ? AND room_ref_id = ? AND deleted_at IS NULL", schoolID, "group", "school", schoolID).First(&room).Error
	return &room, err
}

func (r *chatRepository) CreateSchoolRoom(room *domain.ChatRoom) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "room_sch_id"}, {Name: "room_ref_type"}, {Name: "room_ref_id"}},
		DoNothing: true,
	}).Create(room).Error
}

func (r *chatRepository) GetRoomContext(roomID string, schoolID string) (*ChatRoomRow, error) {
	var row ChatRoomRow
	err := r.db.Raw(chatRoomListSelect()+`
		WHERE cr.room_id = ?
			AND cr.room_sch_id = ?
			AND cr.room_type = 'group'
			AND cr.deleted_at IS NULL
			AND s.deleted_at IS NULL
		LIMIT 1
	`, roomID, schoolID).Scan(&row).Error
	if err != nil {
		return nil, err
	}
	if row.RoomID == "" {
		return nil, gorm.ErrRecordNotFound
	}
	return &row, nil
}

func (r *chatRepository) UserIsActiveSchoolMember(userID string, schoolID string) (bool, error) {
	var count int64
	err := r.db.Table("edv.school_users AS scu").
		Joins("JOIN edv.users u ON u.usr_id = scu.scu_usr_id AND u.deleted_at IS NULL").
		Where("scu.scu_usr_id = ? AND scu.scu_sch_id = ? AND scu.deleted_at IS NULL", userID, schoolID).
		Count(&count).Error
	return count > 0, err
}

func (r *chatRepository) UsersAreActiveSchoolMembers(userIDs []string, schoolID string) (map[string]bool, error) {
	result := make(map[string]bool, len(userIDs))
	if len(userIDs) == 0 {
		return result, nil
	}

	var rows []struct {
		UserID string `gorm:"column:user_id"`
	}
	err := r.db.Raw(`
		SELECT DISTINCT scu.scu_usr_id AS user_id
		FROM edv.school_users scu
		JOIN edv.users u
			ON u.usr_id = scu.scu_usr_id
			AND u.deleted_at IS NULL
		WHERE scu.scu_sch_id = ?
			AND scu.deleted_at IS NULL
			AND scu.scu_usr_id IN ?
	`, schoolID, userIDs).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		result[row.UserID] = true
	}
	return result, nil
}

func (r *chatRepository) UserIsActiveRoomMember(userID string, roomID string) (bool, error) {
	var count int64
	err := r.db.Table("edv.chat_room_members AS crm").
		Joins("JOIN edv.users u ON u.usr_id = crm.crm_usr_id AND u.deleted_at IS NULL").
		Where("crm.crm_room_id = ? AND crm.crm_usr_id = ? AND crm.left_at IS NULL", roomID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *chatRepository) CreateGroupRoomWithMembers(schoolID string, roomName string, creatorID string, memberUserIDs []string) (*domain.ChatRoom, error) {
	var created domain.ChatRoom
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw(`
			INSERT INTO edv.chat_rooms (room_sch_id, room_name, room_type, room_ref_type, room_ref_id, created_by)
			VALUES (?, ?, 'group', NULL, NULL, ?)
			RETURNING room_id, room_sch_id, room_name, room_type, created_by, created_at
		`, schoolID, roomName, creatorID).Scan(&created).Error; err != nil {
			return err
		}

		members := make([]domain.ChatRoomMember, 0, len(memberUserIDs))
		for _, memberID := range memberUserIDs {
			role := "member"
			if memberID == creatorID {
				role = "admin"
			}
			members = append(members, domain.ChatRoomMember{
				RoomID: created.ID,
				UserID: memberID,
				Role:   role,
			})
		}
		if len(members) == 0 {
			return nil
		}
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "crm_room_id"}, {Name: "crm_usr_id"}},
			DoNothing: true,
		}).Create(&members).Error
	})
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *chatRepository) ListMessages(roomID string, limit int, before *time.Time) ([]ChatMessageRow, error) {
	var rows []ChatMessageRow
	query := `
		SELECT *
		FROM (
			SELECT
				msg.msg_id AS message_id,
				msg.msg_room_id AS room_id,
				msg.msg_usr_id AS sender_id,
				COALESCE(u.usr_nama_lengkap, 'Pengguna') AS sender_name,
				COALESCE(sender_role.role_name, 'member') AS sender_role,
				msg.msg_content AS content,
				msg.msg_type AS message_type,
				msg.created_at AS created_at
			FROM edv.chat_messages msg
			JOIN edv.users u ON u.usr_id = msg.msg_usr_id AND u.deleted_at IS NULL
			LEFT JOIN LATERAL (
				SELECT r.rol_name AS role_name
				FROM edv.school_users scu
				JOIN edv.user_roles ur ON ur.urol_scu_id = scu.scu_id
				JOIN edv.roles r ON r.rol_id = ur.urol_rol_id
				JOIN edv.chat_rooms cr_role ON cr_role.room_id = msg.msg_room_id
				WHERE scu.scu_usr_id = msg.msg_usr_id
					AND scu.scu_sch_id = cr_role.room_sch_id
					AND scu.deleted_at IS NULL
					AND r.rol_name IN ('admin', 'teacher', 'student')
				ORDER BY CASE r.rol_name WHEN 'admin' THEN 1 WHEN 'teacher' THEN 2 WHEN 'student' THEN 3 ELSE 4 END
				LIMIT 1
			) sender_role ON true
			WHERE msg.msg_room_id = ?
				AND msg.msg_type = 'text'
				AND msg.deleted_at IS NULL
				AND (?::timestamp IS NULL OR msg.created_at < ?)
			ORDER BY msg.created_at DESC
			LIMIT ?
		) page
		ORDER BY created_at ASC
	`
	var beforeValue any
	if before != nil {
		beforeValue = *before
	}
	err := r.db.Raw(query, roomID, beforeValue, beforeValue, limit).Scan(&rows).Error
	return rows, err
}

func (r *chatRepository) CreateMessage(message *domain.ChatMessage) error {
	return r.db.Create(message).Error
}

func (r *chatRepository) GetMessageByID(messageID string, roomID string) (*ChatMessageRow, error) {
	var row ChatMessageRow
	err := r.db.Raw(`
		SELECT
			msg.msg_id AS message_id,
			msg.msg_room_id AS room_id,
			msg.msg_usr_id AS sender_id,
			COALESCE(u.usr_nama_lengkap, 'Pengguna') AS sender_name,
			COALESCE(sender_role.role_name, 'member') AS sender_role,
			msg.msg_content AS content,
			msg.msg_type AS message_type,
			msg.created_at AS created_at
		FROM edv.chat_messages msg
		JOIN edv.users u ON u.usr_id = msg.msg_usr_id AND u.deleted_at IS NULL
		LEFT JOIN LATERAL (
			SELECT r.rol_name AS role_name
			FROM edv.school_users scu
			JOIN edv.user_roles ur ON ur.urol_scu_id = scu.scu_id
			JOIN edv.roles r ON r.rol_id = ur.urol_rol_id
			JOIN edv.chat_rooms cr_role ON cr_role.room_id = msg.msg_room_id
			WHERE scu.scu_usr_id = msg.msg_usr_id
				AND scu.scu_sch_id = cr_role.room_sch_id
				AND scu.deleted_at IS NULL
				AND r.rol_name IN ('admin', 'teacher', 'student')
			ORDER BY CASE r.rol_name WHEN 'admin' THEN 1 WHEN 'teacher' THEN 2 WHEN 'student' THEN 3 ELSE 4 END
			LIMIT 1
		) sender_role ON true
		WHERE msg.msg_id = ?
			AND msg.msg_room_id = ?
			AND msg.msg_type = 'text'
			AND msg.deleted_at IS NULL
		LIMIT 1
	`, messageID, roomID).Scan(&row).Error
	if err != nil {
		return nil, err
	}
	if row.MessageID == "" {
		return nil, gorm.ErrRecordNotFound
	}
	return &row, nil
}

func (r *chatRepository) UpsertReadReceipt(roomID string, userID string, messageID *string) error {
	now := time.Now()
	receipt := domain.ChatReadReceipt{
		RoomID:            roomID,
		UserID:            userID,
		LastReadMessageID: messageID,
		LastReadAt:        now,
	}
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "rct_room_id"}, {Name: "rct_usr_id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"last_read_msg_id": messageID,
			"last_read_at":     now,
		}),
	}).Create(&receipt).Error
}

func (r *chatRepository) UnreadCount(roomID string, userID string) (int64, error) {
	var count int64
	err := r.db.Raw(`
		SELECT COUNT(*)
		FROM edv.chat_messages msg
		LEFT JOIN edv.chat_read_receipts rct
			ON rct.rct_room_id = msg.msg_room_id
			AND rct.rct_usr_id = ?
		WHERE msg.msg_room_id = ?
			AND msg.msg_usr_id <> ?
			AND msg.msg_type = 'text'
			AND msg.deleted_at IS NULL
			AND (rct.last_read_at IS NULL OR msg.created_at > rct.last_read_at)
	`, userID, roomID, userID).Scan(&count).Error
	return count, err
}

func chatRoomListSelect() string {
	return chatRoomContextSelect() + `
		LEFT JOIN LATERAL (
			SELECT msg.msg_id, msg.msg_usr_id, msg.msg_content, msg.created_at
			FROM edv.chat_messages msg
			WHERE msg.msg_room_id = cr.room_id
				AND msg.msg_type = 'text'
				AND msg.deleted_at IS NULL
			ORDER BY msg.created_at DESC
			LIMIT 1
		) lm ON true
		LEFT JOIN edv.users last_sender ON last_sender.usr_id = lm.msg_usr_id
	`
}

func chatRoomContextSelect() string {
	return `
		SELECT
			cr.room_id AS room_id,
			cr.room_name AS room_name,
			cr.room_type AS room_type,
			cr.room_ref_type AS room_ref_type,
			cr.room_ref_id AS room_ref_id,
			s.sch_id AS school_id,
			s.sch_name AS school_name,
			lm.msg_id AS last_message_id,
			lm.msg_usr_id AS last_sender_id,
			last_sender.usr_nama_lengkap AS last_sender_name,
			lm.msg_content AS last_content,
			lm.created_at AS last_message_at
		FROM edv.chat_rooms cr
		JOIN edv.schools s ON s.sch_id = cr.room_sch_id
	`
}
