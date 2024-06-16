package database

import "time"

const createUser = `INSERT IGNORE INTO Users (Username, Email, PasswordHash) 
       VALUES ('nevik', 'nevik@gmail.com', 'nevim') ON DUPLICATE KEY UPDATE UpdatedAt = time::now()
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

const defineUserTable = `DEFINE TABLE IF NOT EXISTS Users TYPE ANY SCHEMAFULL
	PERMISSIONS
		FOR select
			WHERE Email = $auth.id
		FOR create, update NONE
		FOR delete
			WHERE user = $auth.id OR $auth.admin = true
;
DEFINE FIELD Username ON Users TYPE string
	PERMISSIONS FULL
;
DEFINE FIELD Email ON Users TYPE string VALUE string::lowercase($value) ASSERT string::is::email($value)
	PERMISSIONS FULL
;
DEFINE FIELD PasswordHash ON Users TYPE string
	PERMISSIONS FULL
;
DEFINE FIELD CreatedAt ON Users TYPE datetime VALUE time::now() READONLY
	PERMISSIONS FULL
;
DEFINE FIELD UpdatedAt ON Users TYPE datetime VALUE time::now()
	PERMISSIONS FULL
;
DEFINE INDEX IndexEmail ON Users FIELDS Email UNIQUE;
DEFINE INDEX IndexUsername ON Users FIELDS Username UNIQUE;`
