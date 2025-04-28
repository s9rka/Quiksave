package database

import (
	"context"
	"fmt"
	"notas/models"
	"time"
)

func CreateVault(vault models.Vault) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var vaultID int
	createVaultQuery := `INSERT INTO vaults (name, description, user_id) VALUES ($1, $2, $3) RETURNING id`

	err := dbPool.QueryRow(ctx, createVaultQuery, vault.Name, vault.Description, vault.UserID).Scan(&vaultID)
	if err != nil {
		return -1, fmt.Errorf("failed to create vault: %w", err)
	}

	return vaultID, nil
}

func GetVaults(userID int) ([]models.Vault, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT id, name, description, created_at
		FROM vaults
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := dbPool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vaults: %w", err)
	}
	defer rows.Close()

	var vaults []models.Vault
	for rows.Next() {
		var vault models.Vault
		err := rows.Scan(&vault.ID, &vault.Name, &vault.Description, &vault.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vault: %w", err)
		}
		vault.UserID = userID
		vaults = append(vaults, vault)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return vaults, nil
}

func GetVaultByID(vaultID, userID int) (*models.Vault, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT id, name, description, created_at
		FROM vaults
		WHERE id = $1 AND user_id = $2
	`

	var vault models.Vault
	err := dbPool.QueryRow(ctx, query, vaultID, userID).Scan(&vault.ID, &vault.Name, &vault.Description, &vault.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vault: %w", err)
	}

	vault.UserID = userID
	return &vault, nil
}
