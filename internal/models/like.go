package models

type Like struct {
	LikeID     uint   `gorm:"primaryKey;autoIncrement" json:"like_id"`
	SourceID   uint   `json:"source_id"`
	TargetID   uint   `json:"target_id"`
	TargetType string `json:"target_type"`
}
