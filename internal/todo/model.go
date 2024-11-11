package todo

import "time"

type Todo struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
