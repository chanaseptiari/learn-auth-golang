package models

type MsAuth struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"type:varchar(255);uniqueIndex" binding:"required"`
	Password  string `json:"password" gorm:"type:varchar(255)" binding:"required" validate:"min=8,max=255"`
	Role      string `json:"role" gorm:"type:varchar(100)" binding:"required"`
	ActiveAt  int64  `json:"status"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MsAuth) TableName() string {
	return "ms_auth"
}
