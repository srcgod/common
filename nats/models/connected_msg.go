package models

import "time"

type UserConnectedMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserConnectedMsg) GetUserID() int64        { return e.UserID }
func (e UserConnectedMsg) GetTimestamp() time.Time { return e.Timestamp }

type UserDisconnectedMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserDisconnectedMsg) GetUserID() int64        { return e.UserID }
func (e UserDisconnectedMsg) GetTimestamp() time.Time { return e.Timestamp }

type UserOnlineMsg struct {
	IsOnline bool `json:"is_online"`
}

// PresenceStatusEvent — единый формат события статуса пользователя в NATS.
// Публикуется в фиксированный subject presence.events.status.
type PresenceStatusEvent struct {
	UserID    int64     `json:"user_id"`
	IsOnline  bool      `json:"is_online"`
	Timestamp time.Time `json:"timestamp"`
}

// PresenceHeartbeatMsg — heartbeat от gateway к presence-service.
// Используется, чтобы продлевать online TTL и обновлять last_seen.
type PresenceHeartbeatMsg struct {
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}
