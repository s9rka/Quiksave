package database

import (
	"context"
	"errors"
	"fmt"
	auth "notas/authentication"
	"notas/models"
	"time"
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
		return 0, fmt.Errorf("failed to hash password: %v", err)
	}

	exists, err := UserExists("email", user.Email)
	if err != nil {
		return 0, fmt.Errorf("error checking email existence: %v", err)
	}
	if exists {
		return 0, ErrDuplicateEmail
	}

	exists, err = UserExists("username", user.Username)
	if err != nil {
		return 0, fmt.Errorf("error checking username existence: %v", err)
	}
	if exists {
		return 0, ErrDuplicateUsername
	}

	addUserQuery := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	err = dbPool.QueryRow(ctx, addUserQuery, user.Username, user.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %v", err)
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