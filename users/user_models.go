package users

import (
	"time"
	"fmt"
	"github.com/jmoiron/sqlx"
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


const CREATE_USERS = `CREATE TABLE IF NOT EXISTS users (
id serial,
created_at timestamp NOT NULL DEFAULT NOW(),
updated_at timestamp NOT NULL DEFAULT NOW(),
last_login timestamp NOT NULL DEFAULT NOW(),
username VARCHAR(25) NOT NULL,
password VARCHAR(100) NOT NULL,
email    VARCHAR(100) NOT NULL DEFAULT '',
state    VARCHAR(50) NOT NULL,
PRIMARY KEY (id),
UNIQUE (username),
UNIQUE (state)
);`


/*
 one to one: User has one user_profile
*/

const CREATE_USER_PROFILE = `CREATE TABLE IF NOT EXISTS user_profile (
id serial,
user_id int NOT NULL,
created_at timestamp NOT NULL DEFAULT NOW(),
updated_at timestamp NOT NULL DEFAULT NOW(),
first_name VARCHAR(30) NOT NULL DEFAULT '',
last_name VARCHAR(30) NOT NULL DEFAULT '',
guid VARCHAR(40) NOT NULL,
nickname VARCHAR(30) NOT NULL DEFAULT '',
UNIQUE (guid),
PRIMARY KEY (user_id),
CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);`

/*
 one to one: User has one user_credentials
*/

const CREATE_USER_CREDENTIALS = `CREATE TABLE IF NOT EXISTS user_credentials (
id serial,
user_id int NOT NULL,
created_at timestamp NOT NULL DEFAULT NOW(),
updated_at timestamp NOT NULL DEFAULT NOW(),
access_token VARCHAR(1500) NOT NULL,
refresh_token VARCHAR(1500) NOT NULL,
expiration timestamp NOT NULL DEFAULT NOW(),
type VARCHAR(40),
PRIMARY KEY (user_id),
CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id),
UNIQUE (access_token),
UNIQUE (refresh_token)
);`

const CLEAR_ALL_TABLES = `DELETE FROM users; DELETE FROM user_profile; DELETE FROM user_credentials;`


func MigrateUsers(db *sqlx.DB) error {
	fmt.Println("Creating table users")
	_, err := db.Exec(CREATE_USERS)
	if err != nil {
		return err
	}

	fmt.Println("Creating table user_profiles")
	_, err2 := db.Exec(CREATE_USER_PROFILE)
	if err2 != nil {
		return err2
	}

	fmt.Println("Creating table user_credentials")
	_, err3 := db.Exec(CREATE_USER_CREDENTIALS)
	if err3 != nil {
		return err3
	}
	return nil
}

func RemoveUsersTables(db *sqlx.DB) error {
	fmt.Println("Clearing all users tables")
	_, err := db.Exec(CLEAR_ALL_TABLES)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
