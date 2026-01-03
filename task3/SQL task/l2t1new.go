//根据示例扩展用户模型，增加以下字段：
//Phone：电话号码（字符串，唯一索引）
//LastLoginAt：最后登录时间（时间类型）
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:64;not null"`
    Email     string    `gorm:"size:128;uniqueIndex;not null"`
    Age       uint8     `gorm:"not null"`
    Status    string    `gorm:"size:16;default:active;index"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    phone     string    `gorm:"uniqueindex"`
    LastLoginAt *time.Time
}
