package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	VaultID   int       `json:"vaultId"`
	Heading   string    `json:"heading"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	LastEdit  time.Time `json:"last_edit"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

type Vault struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"created_at"`
}
