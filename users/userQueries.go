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
	err := db.QueryRowx(get_user_from_username, username).StructScan(&u)
	if err != nil {
		return User{}, api.ApiError{
			Status: api.UsernameNotFound,
			Message: "Username not found",
		}
	}
	validated := ValidatePassword([]byte(u.Password), password)
	if validated {  // Fix this to only set error
		return u, nil
	}
	return u, api.ApiError{
		Status: api.IncorrectPassword,
		Message: "Incorrect password",
	}
}

func insertUserProfile(db *sqlx.DB, u User, up UserProfile) (UserProfile, error) {
	var up_new UserProfile
	db_err := db.QueryRowx(create_user_profile, time.Now(), u.ID, up.Firstname, up.Lastname, up.Guid, up.Nickname).StructScan(&up_new)

	if db_err != nil {
		return UserProfile{}, db_err
	}
	return up_new, nil
}

func insertUserCredentials(db *sqlx.DB, u User, uc UserCredential) (UserCredential, error) {
	var uc_new UserCredential

	db_err := db.QueryRowx(create_user_credentials, time.Now(), u.ID, uc.AccessToken, uc.RefreshToken, uc.Expiration, uc.Type).StructScan(&uc_new)
	if db_err != nil {
		return UserCredential{}, db_err
	}
	return uc_new, nil
}

func insertUser(db *sqlx.DB, username, password, state string) (User, error) {
	var new_user User

	db_err := db.QueryRowx(create_user, username, password, state).StructScan(&new_user)
	if db_err != nil {
		pqErr := db_err.(*pq.Error)
		if pqErr.Code == "23505"  {
			return User{}, api.ApiError{
				Status: "23505",
				ErrorString: pqErr.Message,
				Message: fmt.Sprintf("Username [%s] already exists", username),
			}
		}
		return User{}, db_err
	}
	return new_user, nil
}

func updateUser(db *sqlx.DB, user User, new_username, new_password, new_email string) (User, error) {
	var new_user User

	if new_username == "" {
		new_username = user.Username
	}

	//TODO Encrypt password
	if new_password == "" {
		new_password = user.Password
	}

	if new_email == "" {
		new_email = user.Email
	}

	db_err := db.QueryRowx(update_user, new_username, new_password, new_email).StructScan(&new_user)
	if db_err != nil {
		return User{}, nil
	}
	return new_user, nil
}


// UserCredential queries
func getUserCredentialsFromUser(db *sqlx.DB, u User) (UserCredential, error) {
	var uc UserCredential
	db_err := db.QueryRowx(get_user_credential, u.ID).StructScan(&uc)
	if db_err != nil {
		return UserCredential{}, db_err
	}
	return uc, nil
}

func getUserProfileFromUser(db *sqlx.DB, u User) (UserProfile, error) {
	var up UserProfile
	db_err := db.QueryRowx(get_user_profile, u.ID).StructScan(&up)
	if db_err != nil {
		return UserProfile{}, db_err
	}
	return up, nil
}


const get_user_from_username = `SELECT id, username, password FROM users WHERE username=$1`

const get_user_credential = `SELECT * FROM user_credentials WHERE user_id=$1`

const get_user_profile = `SELECT * FROM user_profile WHERE user_id=$1`

const update_user_credentials = `UPDATE user_credentials SET (updated_at, access_token, refresh_token, expiration) =
                                ($1, $2, $3, $4) WHERE id=$5
								RETURNING id, user_id, created_at, updated_at, access_token, refresh_token, expiration, type`

const update_user = `UPDATE users SET (updated_at, username, password, email) =  ($1, $2, $3, $4) WHERE id=$5
					 RETURNING id, created_at, updated_at, username, password, email`

const create_user = `INSERT INTO users
				  (username, password, state)
				   VALUES ($1, $2, $3)
				   RETURNING id, created_at, username, password, state`

const create_user_profile = `INSERT INTO user_profile
				  (updated_at, user_id, first_name, last_name, guid, nickname)
				   VALUES ($1, $2, $3, $4, $5, $6)
				   RETURNING id, user_id, created_at, updated_at, first_name, last_name, guid, nickname`

const create_user_credentials = `INSERT INTO user_credentials
				  (updated_at, user_id, access_token, refresh_token, expiration, type)
				   VALUES ($1, $2, $3, $4, $5, $6)
				   RETURNING id, user_id, created_at, updated_at, access_token, refresh_token, expiration, type`
