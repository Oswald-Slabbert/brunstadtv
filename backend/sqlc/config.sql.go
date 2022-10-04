// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: config.sql

package sqlc

import (
	"context"
	"database/sql"
)

const getAppConfig = `-- name: getAppConfig :one
SELECT id, app_version FROM appconfig LIMIT 1
`

type getAppConfigRow struct {
	ID         int32  `db:"id" json:"id"`
	AppVersion string `db:"app_version" json:"appVersion"`
}

func (q *Queries) getAppConfig(ctx context.Context) (getAppConfigRow, error) {
	row := q.db.QueryRowContext(ctx, getAppConfig)
	var i getAppConfigRow
	err := row.Scan(&i.ID, &i.AppVersion)
	return i, err
}

const getGlobalConfig = `-- name: getGlobalConfig :one
SELECT id, live_online, npaw_enabled FROM globalconfig LIMIT 1
`

type getGlobalConfigRow struct {
	ID          int32        `db:"id" json:"id"`
	LiveOnline  sql.NullBool `db:"live_online" json:"liveOnline"`
	NpawEnabled sql.NullBool `db:"npaw_enabled" json:"npawEnabled"`
}

func (q *Queries) getGlobalConfig(ctx context.Context) (getGlobalConfigRow, error) {
	row := q.db.QueryRowContext(ctx, getGlobalConfig)
	var i getGlobalConfigRow
	err := row.Scan(&i.ID, &i.LiveOnline, &i.NpawEnabled)
	return i, err
}