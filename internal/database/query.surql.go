package database

import "time"

const createUser = `INSERT INTO Users (Username, Email, PasswordHash) 
       VALUES ($Username, $Email, $PasswordHash) ON DUPLICATE KEY UPDATE UpdatedAt = time::now()
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

const (
	defineUserTable = `DEFINE TABLE IF NOT EXISTS Users TYPE ANY SCHEMAFULL
	PERMISSIONS
		FOR create, select, update, delete 
			WHERE user = $auth.id OR $auth.admin = true
;
DEFINE FIELD IF NOT EXISTS Username ON Users TYPE string
	PERMISSIONS FULL
;
DEFINE FIELD IF NOT EXISTS Email ON Users TYPE string VALUE string::lowercase($value) ASSERT string::is::email($value)
	PERMISSIONS FULL
;
DEFINE FIELD IF NOT EXISTS PasswordHash ON Users TYPE string
	PERMISSIONS FULL
;
DEFINE FIELD IF NOT EXISTS CreatedAt ON Users TYPE datetime VALUE time::now() READONLY
	PERMISSIONS FULL
;
DEFINE FIELD IF NOT EXISTS UpdatedAt ON Users TYPE datetime VALUE time::now()
	PERMISSIONS FULL
;
DEFINE INDEX IF NOT EXISTS IndexEmail ON Users FIELDS Email UNIQUE;
DEFINE INDEX IF NOT EXISTS IndexUsername ON Users FIELDS Username UNIQUE;`
	relationalTable = `DEFINE TABLE assigned_to SCHEMAFULL TYPE RELATION IN tag OUT sticky
    PERMISSIONS
        FOR create, select, update, delete 
            WHERE in.owner == $auth.id AND out.author == $auth.id;`
)
