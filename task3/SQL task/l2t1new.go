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
