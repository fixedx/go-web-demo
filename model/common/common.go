package common

type CommonDbModel struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	CreatedAt int64 `gorm:"autoUpdateTime:milli" json:"createdAt"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt int64 `gorm:"index;autoUpdateTime:milli" json:"deletedAt"`
}
