// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: seasons.sql

package sqlc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const refreshSeasonAccessView = `-- name: RefreshSeasonAccessView :one
SELECT update_access('seasons_access')
`

func (q *Queries) RefreshSeasonAccessView(ctx context.Context) (bool, error) {
	row := q.db.QueryRowContext(ctx, refreshSeasonAccessView)
	var update_access bool
	err := row.Scan(&update_access)
	return update_access, err
}

const getPermissionsForSeasons = `-- name: getPermissionsForSeasons :many
WITH sa AS (SELECT se.id,
                   se.status::text = 'published'::text AND
                   s.status::text = 'published'::text                           AS published,
                   COALESCE(GREATEST(se.available_from, s.available_from),
                            '1800-01-01 00:00:00'::timestamp without time zone) AS available_from,
                   COALESCE(LEAST(se.available_to, s.available_to),
                            '3000-01-01 00:00:00'::timestamp without time zone) AS available_to
            FROM seasons se
                     LEFT JOIN shows s ON se.show_id = s.id),
     sr AS (SELECT se.id,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups eu
                             WHERE (eu.episodes_id IN (SELECT e.id
                                                       FROM episodes e
                                                       WHERE e.season_id = se.id))),
                            ARRAY []::character varying[]) AS roles,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups_download eu
                             WHERE (eu.episodes_id IN (SELECT e.id
                                                       FROM episodes e
                                                       WHERE e.season_id = se.id))),
                            ARRAY []::character varying[]) AS roles_download,
                   COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                             FROM episodes_usergroups_earlyaccess eu
                             WHERE (eu.episodes_id IN (SELECT e.id
                                                       FROM episodes e
                                                       WHERE e.season_id = se.id))),
                            ARRAY []::character varying[]) AS roles_earlyaccess
            FROM seasons se)
SELECT se.id,
       access.published::bool             AS published,
       access.available_from::timestamp   AS available_from,
       access.available_to::timestamp     AS available_to,
       roles.roles::varchar[]             AS usergroups,
       roles.roles_download::varchar[]    AS usergroups_downloads,
       roles.roles_earlyaccess::varchar[] AS usergroups_earlyaccess
FROM seasons se
         LEFT JOIN sa access ON access.id = se.id
         LEFT JOIN sr roles ON roles.id = se.id
WHERE se.id = ANY ($1::int[])
`

type getPermissionsForSeasonsRow struct {
	ID                    int32     `db:"id" json:"id"`
	Published             bool      `db:"published" json:"published"`
	AvailableFrom         time.Time `db:"available_from" json:"availableFrom"`
	AvailableTo           time.Time `db:"available_to" json:"availableTo"`
	Usergroups            []string  `db:"usergroups" json:"usergroups"`
	UsergroupsDownloads   []string  `db:"usergroups_downloads" json:"usergroupsDownloads"`
	UsergroupsEarlyaccess []string  `db:"usergroups_earlyaccess" json:"usergroupsEarlyaccess"`
}

func (q *Queries) getPermissionsForSeasons(ctx context.Context, dollar_1 []int32) ([]getPermissionsForSeasonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPermissionsForSeasons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPermissionsForSeasonsRow
	for rows.Next() {
		var i getPermissionsForSeasonsRow
		if err := rows.Scan(
			&i.ID,
			&i.Published,
			&i.AvailableFrom,
			&i.AvailableTo,
			pq.Array(&i.Usergroups),
			pq.Array(&i.UsergroupsDownloads),
			pq.Array(&i.UsergroupsEarlyaccess),
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

const getSeasonIDsForShows = `-- name: getSeasonIDsForShows :many
SELECT s.id, s.show_id
FROM seasons s
WHERE s.show_id = ANY ($1::int[])
ORDER BY s.season_number
`

type getSeasonIDsForShowsRow struct {
	ID     int32 `db:"id" json:"id"`
	ShowID int32 `db:"show_id" json:"showID"`
}

func (q *Queries) getSeasonIDsForShows(ctx context.Context, dollar_1 []int32) ([]getSeasonIDsForShowsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSeasonIDsForShows, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSeasonIDsForShowsRow
	for rows.Next() {
		var i getSeasonIDsForShowsRow
		if err := rows.Scan(&i.ID, &i.ShowID); err != nil {
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

const getSeasons = `-- name: getSeasons :many
WITH ts AS (SELECT seasons_id,
                   json_object_agg(languages_code, title)       AS title,
                   json_object_agg(languages_code, description) AS description
            FROM seasons_translations
            GROUP BY seasons_id),
     tags AS (SELECT seasons_id,
                     array_agg(tags_id) AS tags
              FROM seasons_tags
              GROUP BY seasons_id)
SELECT s.id,
       s.legacy_id,
       s.season_number,
       fs.filename_disk as image_file_name,
       s.show_id,
       COALESCE(s.agerating_code, 'A')                   as agerating,
       tags.tags::int[]                                  AS tag_ids,
       ts.title,
       ts.description
FROM seasons s
         JOIN ts ON s.id = ts.seasons_id
         LEFT JOIN tags ON tags.seasons_id = s.id
         JOIN shows sh ON s.show_id = sh.id
         LEFT JOIN directus_files fs ON fs.id = COALESCE(s.image_file_id, sh.image_file_id)
WHERE s.id = ANY ($1::int[])
`

type getSeasonsRow struct {
	ID            int32           `db:"id" json:"id"`
	LegacyID      null_v4.Int     `db:"legacy_id" json:"legacyID"`
	SeasonNumber  int32           `db:"season_number" json:"seasonNumber"`
	ImageFileName null_v4.String  `db:"image_file_name" json:"imageFileName"`
	ShowID        int32           `db:"show_id" json:"showID"`
	Agerating     string          `db:"agerating" json:"agerating"`
	TagIds        []int32         `db:"tag_ids" json:"tagIds"`
	Title         json.RawMessage `db:"title" json:"title"`
	Description   json.RawMessage `db:"description" json:"description"`
}

func (q *Queries) getSeasons(ctx context.Context, dollar_1 []int32) ([]getSeasonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSeasons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSeasonsRow
	for rows.Next() {
		var i getSeasonsRow
		if err := rows.Scan(
			&i.ID,
			&i.LegacyID,
			&i.SeasonNumber,
			&i.ImageFileName,
			&i.ShowID,
			&i.Agerating,
			pq.Array(&i.TagIds),
			&i.Title,
			&i.Description,
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

const listSeasons = `-- name: listSeasons :many
WITH ts AS (SELECT seasons_id,
                   json_object_agg(languages_code, title)       AS title,
                   json_object_agg(languages_code, description) AS description
            FROM seasons_translations
            GROUP BY seasons_id),
     tags AS (SELECT seasons_id,
                     array_agg(tags_id) AS tags
              FROM seasons_tags
              GROUP BY seasons_id)
SELECT s.id,
       s.legacy_id,
       s.season_number,
       fs.filename_disk as image_file_name,
       s.show_id,
       COALESCE(s.agerating_code, 'A')                   as agerating,
       tags.tags::int[]                                  AS tag_ids,
       ts.title,
       ts.description
FROM seasons s
         JOIN ts ON s.id = ts.seasons_id
         LEFT JOIN tags ON tags.seasons_id = s.id
         JOIN shows sh ON s.show_id = sh.id
         LEFT JOIN directus_files fs ON fs.id = COALESCE(s.image_file_id, sh.image_file_id)
`

type listSeasonsRow struct {
	ID            int32           `db:"id" json:"id"`
	LegacyID      null_v4.Int     `db:"legacy_id" json:"legacyID"`
	SeasonNumber  int32           `db:"season_number" json:"seasonNumber"`
	ImageFileName null_v4.String  `db:"image_file_name" json:"imageFileName"`
	ShowID        int32           `db:"show_id" json:"showID"`
	Agerating     string          `db:"agerating" json:"agerating"`
	TagIds        []int32         `db:"tag_ids" json:"tagIds"`
	Title         json.RawMessage `db:"title" json:"title"`
	Description   json.RawMessage `db:"description" json:"description"`
}

func (q *Queries) listSeasons(ctx context.Context) ([]listSeasonsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSeasons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listSeasonsRow
	for rows.Next() {
		var i listSeasonsRow
		if err := rows.Scan(
			&i.ID,
			&i.LegacyID,
			&i.SeasonNumber,
			&i.ImageFileName,
			&i.ShowID,
			&i.Agerating,
			pq.Array(&i.TagIds),
			&i.Title,
			&i.Description,
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
