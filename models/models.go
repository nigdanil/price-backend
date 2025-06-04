package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleClient  = "client"
)

type PaginatedProductsResponse struct {
	Products []GalleryProduct `json:"products"`
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
	Total    int64            `json:"total"`
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"type:text;check:role IN ('admin','manager','client');not null" json:"role"`
	CreatedAt    time.Time `gorm:"default:now()" json:"created_at"`
}

type UserAuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Action    string    `gorm:"not null" json:"action"`
	Method    string    `gorm:"not null" json:"method"`
	Endpoint  string    `gorm:"not null" json:"endpoint"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}

type GalleryProduct struct {
	ID         string `gorm:"primaryKey"`
	Title      string
	ProductURL string `gorm:"uniqueIndex"`
	ImageURL   string
	DataUpload *time.Time
	IsActive   bool
	CategoryID string
}

type GalleryCategory struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

type CategoryURL struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	CategoryURL string `gorm:"uniqueIndex"`
	IsActive    bool
	LastChecked *time.Time
	Source      string
	CategoryID  string
}

type GalleryProductPrice struct {
	ID         uint      `json:"id"`
	ProductURL string    `json:"product_url"`
	Price      float64   `json:"price"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID string    `json:"category_id"`
}
