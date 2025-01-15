package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	auth "notas/authentication"
	"notas/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
    ErrDuplicateEmail    = errors.New("email already exists")
    ErrDuplicateUsername = errors.New("username already exists")
)

func AddUserToDB(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	var userID int
	hashedPassword, err := auth.GenerateHashedPassword(user.Password)
	if err != nil {
		return -1, fmt.Errorf("failed to hash password: %v", err)
	}

	exists, err := UserExists("email", user.Email)
	if err != nil {
		return -1, fmt.Errorf("error checking email existence: %v", err)
	}
	if exists {
		return -1, ErrDuplicateEmail
	}

	exists, err = UserExists("username", user.Username)
	if err != nil {
		return -1, fmt.Errorf("error checking username existence: %v", err)
	}
	if exists {
		return -1, ErrDuplicateUsername
	}

	addUserQuery := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	err = dbPool.QueryRow(ctx, addUserQuery, user.Username, user.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return -1, fmt.Errorf("failed to insert user: %v", err)
	}

	return userID, nil
}

func UserExists(field, value string) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
    defer cancel()

    query := ""
    switch field {
    case "email":
        query = "SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)"
    case "username":
        query = "SELECT EXISTS (SELECT 1 FROM users WHERE username=$1)"
    default:
        return false, fmt.Errorf("invalid field: %s", field)
    }

    var exists bool
    err := dbPool.QueryRow(ctx, query, value).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func getUserPasswordHash(username string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	queryString := "SELECT id, password_hash FROM users WHERE userName=$1"
	var hash string
	var id int
	err := dbPool.QueryRow(ctx, queryString, username).Scan(&id, &hash)
	if err != nil {
		return -1, "", fmt.Errorf("failed to retrieve password hash: %v", err)
	}
	return id, hash, nil
}

func ValidateUserLogin(username, password string) (int, error) {
	id, hash, err := getUserPasswordHash(username)
	if err != nil {
		return -1, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetUserByID(userID int) (*models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    query := "SELECT username, email FROM users WHERE id = $1"
    var user models.User
    err := dbPool.QueryRow(ctx, query, userID).Scan(&user.Username, &user.Email)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to query user: %v", err)
    }

    return &models.User{
        Username: user.Username,
        Email:    user.Email,
    }, nil
}
