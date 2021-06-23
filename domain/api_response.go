package domain

import "time"

type ResponseRedeemLicense struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	User struct {
		ID       string    `json:"id"`
		FullName string    `json:"full_name"`
		Email    string    `json:"email"`
		UsedAt   time.Time `json:"used_at"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
