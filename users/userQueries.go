package users

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"gofant/api"
	"time"
)



// User queries
func GetUserFromPassword(db *sqlx.DB, username string, password string) (User, error) {
	var u User
	if dbErr := db.QueryRowx(selectUserFromUsername, username).StructScan(&u); dbErr != nil {
		return User{}, api.ApiError{
			Status: api.UsernameNotFound,
			Message: "Username not found",
		}
	}
	if validated := ValidatePassword([]byte(u.Password), password); validated == true {
		return u, nil
	}
	return u, api.ApiError{
		Status: api.IncorrectPassword,
		Message: "Incorrect password",
	}
}

func insertUserProfile(db *sqlx.DB, u User, up UserProfile) (UserProfile, error) {
	var upNew UserProfile
	if dbErr := db.QueryRowx(createUserProfile, time.Now(), u.ID, up.Firstname, up.Lastname, up.Guid, up.Nickname).StructScan(&upNew); dbErr != nil {
		return UserProfile{}, dbErr
	}
	return upNew, nil
}

func insertUserCredentials(db *sqlx.DB, u User, uc UserCredential) (UserCredential, error) {
	var ucNew UserCredential
	if dbErr := db.QueryRowx(createUserCredentials, time.Now(), u.ID, uc.AccessToken, uc.RefreshToken, uc.Expiration, uc.Type).StructScan(&ucNew); dbErr != nil {
		return UserCredential{}, dbErr
	}
	return ucNew, nil
}

func insertUser(db *sqlx.DB, username, password, state string) (User, error) {
	var userNew User

	if dbErr := db.QueryRowx(createUser, username, password, state).StructScan(&userNew); dbErr != nil {
		pqErr := dbErr.(*pq.Error)
		if pqErr.Code == "23505"  {
			return User{}, api.ApiError{
				Status: "23505",
				ErrorString: pqErr.Message,
				Message: fmt.Sprintf("Username [%s] already exists", username),
			}
		}
		return User{}, dbErr
	}
	return userNew, nil
}

func updateUser(db *sqlx.DB, user User, username, password, email string) (User, error) {
	var userUpdated User

	if username == "" {
		username = user.Username
	}

	//TODO Encrypt password
	if password == "" {
		password = user.Password
	}

	if email == "" {
		email = user.Email
	}

	if dbErr := db.QueryRowx(update_user, username, password, email).StructScan(&userUpdated); dbErr != nil {
		return User{}, nil
	}
	return userUpdated, nil
}


// UserCredential queries
func getUserCredentialsFromUser(db *sqlx.DB, u User) (UserCredential, error) {
	var uc UserCredential
	if dbErr := db.QueryRowx(selectUserCredentials, u.ID).StructScan(&uc); dbErr != nil {
		return UserCredential{}, dbErr
	}
	return uc, nil
}

func getUserProfileFromUser(db *sqlx.DB, u User) (UserProfile, error) {
	var up UserProfile
	if dbErr := db.QueryRowx(selectUserProfile, u.ID).StructScan(&up); dbErr != nil {
		return UserProfile{}, dbErr
	}
	return up, nil
}


const selectUserFromUsername = `SELECT id, username, password FROM users WHERE username=$1`

const selectUserCredentials = `SELECT * FROM user_credentials WHERE user_id=$1`

const selectUserProfile = `SELECT * FROM user_profile WHERE user_id=$1`

const updateUserCredentials = `UPDATE user_credentials SET (updated_at, access_token, refresh_token, expiration) =
                                ($1, $2, $3, $4) WHERE id=$5
								RETURNING id, user_id, created_at, updated_at, access_token, refresh_token, expiration, type`

const update_user = `UPDATE users SET (updated_at, username, password, email) =  ($1, $2, $3, $4) WHERE id=$5
					 RETURNING id, created_at, updated_at, username, password, email`

const createUser = `INSERT INTO users
				  (username, password, state)
				   VALUES ($1, $2, $3)
				   RETURNING id, created_at, username, password, state`

const createUserProfile = `INSERT INTO user_profile
				  (updated_at, user_id, first_name, last_name, guid, nickname)
				   VALUES ($1, $2, $3, $4, $5, $6)
				   RETURNING id, user_id, created_at, updated_at, first_name, last_name, guid, nickname`

const createUserCredentials = `INSERT INTO user_credentials
				  (updated_at, user_id, access_token, refresh_token, expiration, type)
				   VALUES ($1, $2, $3, $4, $5, $6)
				   RETURNING id, user_id, created_at, updated_at, access_token, refresh_token, expiration, type`
