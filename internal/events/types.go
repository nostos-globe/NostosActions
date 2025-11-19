package events

import "time"

type ContentLikedEvent struct {
	SourceID   uint      `json:"sourceId"`
	TargetID   uint      `json:"targetId"`
	TargetType string    `json:"targetType"` // trip | media | post
	CreatedAt  time.Time `json:"createdAt"`
}

type UserFollowedEvent struct {
	FollowerID uint      `json:"followerId"`
	FollowedID uint      `json:"followedId"`
	CreatedAt  time.Time `json:"createdAt"`
}
