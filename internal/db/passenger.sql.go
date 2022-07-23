// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: passenger.sql

package db

import (
	"context"
	"time"
)

const createPassenger = `-- name: CreatePassenger :one
INSERT INTO passenger (
  phone,
  hashed_password,
  name,
  date_of_birth
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at
`

type CreatePassengerParams struct {
	Phone          string    `json:"phone"`
	HashedPassword string    `json:"hashed_password"`
	Name           string    `json:"name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
}

func (q *Queries) CreatePassenger(ctx context.Context, arg CreatePassengerParams) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, createPassenger,
		arg.Phone,
		arg.HashedPassword,
		arg.Name,
		arg.DateOfBirth,
	)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const deletePassenger = `-- name: DeletePassenger :exec
DELETE FROM passenger
WHERE phone = $1
`

func (q *Queries) DeletePassenger(ctx context.Context, phone string) error {
	_, err := q.db.ExecContext(ctx, deletePassenger, phone)
	return err
}

const getPassenger = `-- name: GetPassenger :one
SELECT id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at FROM passenger
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPassenger(ctx context.Context, id int64) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, getPassenger, id)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getPassengerByPhone = `-- name: GetPassengerByPhone :one
SELECT id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at FROM passenger
WHERE phone = $1 LIMIT 1
`

func (q *Queries) GetPassengerByPhone(ctx context.Context, phone string) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, getPassengerByPhone, phone)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getPassengerForUpdate = `-- name: GetPassengerForUpdate :one
SELECT id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at FROM passenger
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetPassengerForUpdate(ctx context.Context, id int64) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, getPassengerForUpdate, id)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const listPassengers = `-- name: ListPassengers :many
SELECT id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at FROM passenger
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPassengersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPassengers(ctx context.Context, arg ListPassengersParams) ([]Passenger, error) {
	rows, err := q.db.QueryContext(ctx, listPassengers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Passenger
	for rows.Next() {
		var i Passenger
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.HashedPassword,
			&i.Name,
			&i.DateOfBirth,
			&i.AvatarUrl,
			&i.Verified,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePassenger = `-- name: UpdatePassenger :one

UPDATE passenger
SET phone = $2,
  name = $3,
  date_of_birth = $4
WHERE id = $1
RETURNING id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at
`

type UpdatePassengerParams struct {
	ID          int64     `json:"id"`
	Phone       string    `json:"phone"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

// pagination: offset: skip many rows
func (q *Queries) UpdatePassenger(ctx context.Context, arg UpdatePassengerParams) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, updatePassenger,
		arg.ID,
		arg.Phone,
		arg.Name,
		arg.DateOfBirth,
	)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE passenger
SET hashed_password = $2
WHERE id = $1
RETURNING id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at
`

type UpdatePasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.HashedPassword)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const verify = `-- name: Verify :one
UPDATE passenger
SET verified = true
WHERE phone = $1
RETURNING id, phone, hashed_password, name, date_of_birth, avatar_url, verified, created_at
`

func (q *Queries) Verify(ctx context.Context, phone string) (Passenger, error) {
	row := q.db.QueryRowContext(ctx, verify, phone)
	var i Passenger
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.HashedPassword,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}
