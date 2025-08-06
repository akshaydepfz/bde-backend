package models
import("time")

type BDEUser struct {
	ID             uint      `json:"id"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	PasswordHash   string    `json:"-"`
	DrivingLicense string   `json:"driving_license,omitempty"` // nullable
	Role           string    `json:"role"`
	JoinDate       time.Time `json:"join_date"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

