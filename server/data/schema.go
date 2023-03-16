package data

import (
	"context"
	"database/sql"
)

// Database SQL schemas

var assetSchema = []string{
	`CREATE TABLE IF NOT EXISTS assets (
		asset_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		exec_contract INTEGER NOT NULL,
    asset_contract INTEGER NOT NULL,
    img_preview TEXT NOT NULL,
    executor INTEGER NOT NULL,
    desc TEXT NOT NULL,
    price REAL NOT NULL,
    FOREIGN KEY(exec_contract) REFERENCES contracts(contract_id),
    FOREIGN KEY(asset_contract) REFERENCES contracts(contract_id),
    FOREIGN KEY(executor) REFERENCES users(user_id))`,
}

var contractSchema = []string{
	`CREATE TABLE IF NOT EXISTS contracts (
    contract_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    addr TEXT NOT NULL,
    is_exec INTEGER NOT NULL,
    is_finished INTEGER)`,
}

var userSchema = []string{
	`CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    addr TEXT NOT NULL,
    nonce TEXT NOT NULL,
    profile_pic_url TEXT NOT NULL,
    bio TEXT NOT NULL)`,
}

// Contribution tx status:
// 0 - Pending
// 1 - Success
// 2 - Failure
var contributionSchema = []string{
	`CREATE TABLE IF NOT EXISTS contributions (
    contrib_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    exec_contract INTEGER NOT NULL,
    user INTEGER NOT NULL,
    amt REAL NOT NULL,
    status INTEGER NOT NULL,
    FOREIGN KEY(exec_contract) REFERENCES contracts(contract_id),
    FOREIGN KEY(user) REFERENCES users(user_id)
  )`,
}

// Drop db expressions

var assetResetExprs = []string{
	`DROP TABLE IF EXISTS assets`,
}

var contractResetExprs = []string{
	`DROP TABLE IF EXISTS contracts`,
}

var userResetExprs = []string{
	`DROP TABLE IF EXISTS users`,
}

var contributionResetExprs = []string{
	`DROP TABLE IF EXISTS contributions`,
}

// Atomic transaction that instantiates tables if they don't already exist
func initTables(ctx context.Context, tx *sql.Tx) error {
	// First, create tables with no foreign keys
	if _, err := tx.Exec(userSchema[0]); err != nil {
		return err
	}
	if _, err := tx.Exec(contractSchema[0]); err != nil {
		return err
	}

	// The rest of the tables
	if _, err := tx.Exec(contributionSchema[0]); err != nil {
		return err
	}
	if _, err := tx.Exec(assetSchema[0]); err != nil {
		return err
	}

	return nil
}

// Atomic transaction that drops all tables if they exist
func dropTables(ctx context.Context, tx *sql.Tx) error {
	// First, drop tables without dependencies
	if _, err := tx.Exec(contributionResetExprs[0]); err != nil {
		return err
	}
	if _, err := tx.Exec(assetResetExprs[0]); err != nil {
		return err
	}

	// Then, the rest of the tables
	if _, err := tx.Exec(userResetExprs[0]); err != nil {
		return err
	}
	if _, err := tx.Exec(contractResetExprs[0]); err != nil {
		return err
	}

	return nil
}
