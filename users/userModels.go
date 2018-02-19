package users

import (
	"time"
)

// Data structures for receiving Yahoo JSON api calls
type UserResult struct {
	Profile struct {
		Firstname   string          `json:"givenName"`
		Lastname    string  		`json:"familyName"`
		Guid        string  		`json:"guid"`
		Nickname    string  		`json:"nickName"`
		MemberSince string  		`json:"memberSince"`
		Created     string  		`json:"created"`
		ProfileMode string  		`json:"profileMode"`
		ProfilePermission string 	`json:"profilePermission"`
	} `json:"profile"`
}

// Data structures for receiving Yahoo JSON api calls
type CredentialResult struct {
		AccessToken   string        `json:"access_token"`
		RefreshToken  string  		`json:"refresh_token"`
		ExpiresIn     int  		    `json:"expires_in"`
		TokenType     string  		`json:"token_type"`
		Guid		  string  		`json:"xoauth_yahoo_guid"`
}

// Client Request Struct
type UserRequest struct {
		Username 	string
		Password 	string
		Guid 		string
}

// Database Objects for storage
type User struct {
	ID              int
	CreatedAt       time.Time  	`db:"created_at" json:"created_at"`
	LastLogin		time.Time  	`db:"last_login" json:"last_login"`
	UpdatedAt       time.Time  	`db:"updated_at" json:"updated_at"`
	Username		string		`db:"username" json:"username"`
	Password		string		`db:"password" json:"password"`
	Email           string      `db:"email" json:"email"`
	State 			string		`db:"state" json:"state"`
}

type UserProfile struct {
	ID              int
	CreatedAt       time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time   `db:"updated_at" json:"updated_at"`
	UserID		    int			`db:"user_id" json:"user_id"`
	Firstname       string      `db:"first_name" json:"first_name"`
	Lastname       	string      `db:"last_name" json:"last_name"`
	Guid            string      `db:"guid" json:"guid"`
	Nickname        string		`db:"nickname" json:"nickname"`
}

type UserCredential struct {
	ID              int
	UserID		    int			`db:"user_id" json:"user_id"`
	CreatedAt       time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt       time.Time	`db:"updated_at" json:"updated_at"`
	AccessToken     string		`db:"access_token" json:"access_token"`
	RefreshToken    string		`db:"refresh_token" json:"refresh_token"`
	Expiration      time.Time 	`db:"expiration" json:"expiration"`
	Type            string		`db:"type" json:"type"`
}

type Credentials interface {
	checkToken() bool
	refreshToken() string
}