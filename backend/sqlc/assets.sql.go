// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: assets.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getFilesForAssets = `-- name: getFilesForAssets :many
SELECT 0::int as episodes_id, f.id, f.user_created, f.date_created, f.user_updated, f.date_updated, f.type, f.asset_id, f.extra_metadata, f.path, f.audio_language_id, f.subtitle_language_id, f.mime_type, f.storage
FROM assets a
         JOIN assetfiles f ON a.id = f.asset_id
WHERE a.id = ANY ($1::int[])
`

type getFilesForAssetsRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesID"`
	ID                 int32                 `db:"id" json:"id"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	Type               string                `db:"type" json:"type"`
	AssetID            int32                 `db:"asset_id" json:"assetID"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	Path               string                `db:"path" json:"path"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageID"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageID"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Storage            string                `db:"storage" json:"storage"`
}

func (q *Queries) getFilesForAssets(ctx context.Context, dollar_1 []int32) ([]getFilesForAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFilesForAssets, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getFilesForAssetsRow
	for rows.Next() {
		var i getFilesForAssetsRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.ID,
			&i.UserCreated,
			&i.DateCreated,
			&i.UserUpdated,
			&i.DateUpdated,
			&i.Type,
			&i.AssetID,
			&i.ExtraMetadata,
			&i.Path,
			&i.AudioLanguageID,
			&i.SubtitleLanguageID,
			&i.MimeType,
			&i.Storage,
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

const getFilesForEpisodes = `-- name: getFilesForEpisodes :many
SELECT e.id AS episodes_id, f.id, f.user_created, f.date_created, f.user_updated, f.date_updated, f.type, f.asset_id, f.extra_metadata, f.path, f.audio_language_id, f.subtitle_language_id, f.mime_type, f.storage
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetfiles f ON a.id = f.asset_id
WHERE e.id = ANY ($1::int[])
`

type getFilesForEpisodesRow struct {
	EpisodesID         int32                 `db:"episodes_id" json:"episodesID"`
	ID                 int32                 `db:"id" json:"id"`
	UserCreated        uuid.NullUUID         `db:"user_created" json:"userCreated"`
	DateCreated        time.Time             `db:"date_created" json:"dateCreated"`
	UserUpdated        uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	DateUpdated        time.Time             `db:"date_updated" json:"dateUpdated"`
	Type               string                `db:"type" json:"type"`
	AssetID            int32                 `db:"asset_id" json:"assetID"`
	ExtraMetadata      pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	Path               string                `db:"path" json:"path"`
	AudioLanguageID    null_v4.String        `db:"audio_language_id" json:"audioLanguageID"`
	SubtitleLanguageID null_v4.String        `db:"subtitle_language_id" json:"subtitleLanguageID"`
	MimeType           string                `db:"mime_type" json:"mimeType"`
	Storage            string                `db:"storage" json:"storage"`
}

func (q *Queries) getFilesForEpisodes(ctx context.Context, dollar_1 []int32) ([]getFilesForEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getFilesForEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getFilesForEpisodesRow
	for rows.Next() {
		var i getFilesForEpisodesRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.ID,
			&i.UserCreated,
			&i.DateCreated,
			&i.UserUpdated,
			&i.DateUpdated,
			&i.Type,
			&i.AssetID,
			&i.ExtraMetadata,
			&i.Path,
			&i.AudioLanguageID,
			&i.SubtitleLanguageID,
			&i.MimeType,
			&i.Storage,
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

const getStreamsForAssets = `-- name: getStreamsForAssets :many
WITH audiolang AS (SELECT s.id, array_agg(al.languages_code) langs
                   FROM assets a
                            LEFT JOIN assetstreams s ON a.id = s.asset_id
                            LEFT JOIN assetstreams_audio_languages al ON al.assetstreams_id = s.id
                   WHERE al.languages_code IS NOT NULL
                   GROUP BY s.id),
     sublang AS (SELECT s.id, array_agg(al.languages_code) langs
                 FROM assets a
                          LEFT JOIN assetstreams s ON a.id = s.asset_id
                          LEFT JOIN assetstreams_subtitle_languages al ON al.assetstreams_id = s.id
                 WHERE al.languages_code IS NOT NULL
                 GROUP BY s.id)
SELECT 0::int as episodes_id, s.asset_id, s.date_created, s.date_updated, s.encryption_key_id, s.extra_metadata, s.id, s.legacy_videourl_id, s.path, s.service, s.status, s.type, s.url, s.user_created, s.user_updated, COALESCE(al.langs, '{}')::text[] audio_languages, COALESCE(sl.langs, '{}')::text[] subtitle_languages
FROM assets a
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE a.id = ANY ($1::int[])
`

type getStreamsForAssetsRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesID"`
	ID                int32                 `db:"id" json:"id"`
	Status            string                `db:"status" json:"status"`
	UserCreated       uuid.NullUUID         `db:"user_created" json:"userCreated"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	UserUpdated       uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	Url               string                `db:"url" json:"url"`
	Type              string                `db:"type" json:"type"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	AssetID           int32                 `db:"asset_id" json:"assetID"`
	Path              string                `db:"path" json:"path"`
	Service           string                `db:"service" json:"service"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyID"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlID"`
	AudioLanguages    []string              `db:"audio_languages" json:"audioLanguages"`
	SubtitleLanguages []string              `db:"subtitle_languages" json:"subtitleLanguages"`
}

func (q *Queries) getStreamsForAssets(ctx context.Context, dollar_1 []int32) ([]getStreamsForAssetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getStreamsForAssets, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getStreamsForAssetsRow
	for rows.Next() {
		var i getStreamsForAssetsRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.ID,
			&i.Status,
			&i.UserCreated,
			&i.DateCreated,
			&i.UserUpdated,
			&i.DateUpdated,
			&i.Url,
			&i.Type,
			&i.ExtraMetadata,
			&i.AssetID,
			&i.Path,
			&i.Service,
			&i.EncryptionKeyID,
			&i.LegacyVideourlID,
			pq.Array(&i.AudioLanguages),
			pq.Array(&i.SubtitleLanguages),
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

const getStreamsForEpisodes = `-- name: getStreamsForEpisodes :many
WITH audiolang AS (SELECT s.id, array_agg(al.languages_code) langs
                   FROM episodes e
                            JOIN assets a ON e.asset_id = a.id
                            LEFT JOIN assetstreams s ON a.id = s.asset_id
                            LEFT JOIN assetstreams_audio_languages al ON al.assetstreams_id = s.id
                   WHERE al.languages_code IS NOT NULL
                   GROUP BY s.id),
     sublang AS (SELECT s.id, array_agg(al.languages_code) langs
                 FROM episodes e
                          JOIN assets a ON e.asset_id = a.id
                          LEFT JOIN assetstreams s ON a.id = s.asset_id
                          LEFT JOIN assetstreams_subtitle_languages al ON al.assetstreams_id = s.id
                 WHERE al.languages_code IS NOT NULL
                 GROUP BY s.id)
SELECT e.id AS episodes_id, s.id, s.status, s.user_created, s.date_created, s.user_updated, s.date_updated, s.url, s.type, s.extra_metadata, s.asset_id, s.path, s.service, s.encryption_key_id, s.legacy_videourl_id, COALESCE(al.langs, array[])::text[] audio_languages, COALESCE(sl.langs, array[])::text[] subtitle_languages
FROM episodes e
         JOIN assets a ON e.asset_id = a.id
         JOIN assetstreams s ON a.id = s.asset_id
         LEFT JOIN audiolang al ON al.id = s.id
         LEFT JOIN sublang sl ON sl.id = s.id
WHERE e.id = ANY($1::int[])
`

type getStreamsForEpisodesRow struct {
	EpisodesID        int32                 `db:"episodes_id" json:"episodesID"`
	ID                int32                 `db:"id" json:"id"`
	Status            string                `db:"status" json:"status"`
	UserCreated       uuid.NullUUID         `db:"user_created" json:"userCreated"`
	DateCreated       time.Time             `db:"date_created" json:"dateCreated"`
	UserUpdated       uuid.NullUUID         `db:"user_updated" json:"userUpdated"`
	DateUpdated       time.Time             `db:"date_updated" json:"dateUpdated"`
	Url               string                `db:"url" json:"url"`
	Type              string                `db:"type" json:"type"`
	ExtraMetadata     pqtype.NullRawMessage `db:"extra_metadata" json:"extraMetadata"`
	AssetID           int32                 `db:"asset_id" json:"assetID"`
	Path              string                `db:"path" json:"path"`
	Service           string                `db:"service" json:"service"`
	EncryptionKeyID   null_v4.String        `db:"encryption_key_id" json:"encryptionKeyID"`
	LegacyVideourlID  null_v4.Int           `db:"legacy_videourl_id" json:"legacyVideourlID"`
	AudioLanguages    []string              `db:"audio_languages" json:"audioLanguages"`
	SubtitleLanguages []string              `db:"subtitle_languages" json:"subtitleLanguages"`
}

func (q *Queries) getStreamsForEpisodes(ctx context.Context, dollar_1 []int32) ([]getStreamsForEpisodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getStreamsForEpisodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getStreamsForEpisodesRow
	for rows.Next() {
		var i getStreamsForEpisodesRow
		if err := rows.Scan(
			&i.EpisodesID,
			&i.ID,
			&i.Status,
			&i.UserCreated,
			&i.DateCreated,
			&i.UserUpdated,
			&i.DateUpdated,
			&i.Url,
			&i.Type,
			&i.ExtraMetadata,
			&i.AssetID,
			&i.Path,
			&i.Service,
			&i.EncryptionKeyID,
			&i.LegacyVideourlID,
			pq.Array(&i.AudioLanguages),
			pq.Array(&i.SubtitleLanguages),
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
