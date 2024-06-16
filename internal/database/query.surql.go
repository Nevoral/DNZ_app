package database

import "time"

const createUser = `INSERT INTO Users (Username, Email, PasswordHash, CreatedAt, UpdatedAt) 
VALUES ($Username, $Email, $PasswordHash, $CreatedAt, $UpdatedAt) ON DUPLICATE KEY UPDATE 
	UpdatedAt = $input.UpdatedAt
	RETURN id, Username;
`

type (
	CreateUserParam struct {
		Username     string
		Email        string
		PasswordHash string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	CreateUserResult struct {
		Id       string `json:"id"`
		Username string `json:"Username"`
	}
)
