// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: translations.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const listEpisodeTranslations = `-- name: ListEpisodeTranslations :many
WITH episodes AS (SELECT e.id
                  FROM episodes e
                           LEFT JOIN seasons s ON s.id = e.season_id
                           LEFT JOIN shows sh ON sh.id = s.show_id
                  WHERE e.status = 'published'
                    AND s.status = 'published'
                    AND sh.status = 'published')
SELECT et.id, episodes_id as parent_id, languages_code, title, description, extra_description
FROM episodes_translations et
         JOIN episodes e ON e.id = et.episodes_id
WHERE et.languages_code = ANY($1::varchar[])
`

type ListEpisodeTranslationsRow struct {
	ID               int32          `db:"id" json:"id"`
	ParentID         int32          `db:"parent_id" json:"parentID"`
	LanguagesCode    string         `db:"languages_code" json:"languagesCode"`
	Title            null_v4.String `db:"title" json:"title"`
	Description      null_v4.String `db:"description" json:"description"`
	ExtraDescription null_v4.String `db:"extra_description" json:"extraDescription"`
}

func (q *Queries) ListEpisodeTranslations(ctx context.Context, dollar_1 []string) ([]ListEpisodeTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listEpisodeTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEpisodeTranslationsRow
	for rows.Next() {
		var i ListEpisodeTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
			&i.Title,
			&i.Description,
			&i.ExtraDescription,
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

const listSeasonTranslations = `-- name: ListSeasonTranslations :many
WITH seasons AS (SELECT s.id
                 FROM seasons s
                          LEFT JOIN shows sh ON sh.id = s.show_id
                 WHERE s.status = 'published'
                   AND sh.status = 'published')
SELECT et.id, seasons_id as parent_id, languages_code, title, description
FROM seasons_translations et
         JOIN seasons e ON e.id = et.seasons_id
WHERE et.languages_code = ANY($1::varchar[])
`

type ListSeasonTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      int32          `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListSeasonTranslations(ctx context.Context, dollar_1 []string) ([]ListSeasonTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSeasonTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSeasonTranslationsRow
	for rows.Next() {
		var i ListSeasonTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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

const listShowTranslations = `-- name: ListShowTranslations :many
WITH shows AS (SELECT s.id
               FROM shows s
               WHERE s.status = 'published')
SELECT et.id, shows_id as parent_id, languages_code, title, description
FROM shows_translations et
         JOIN shows e ON e.id = et.shows_id
WHERE et.languages_code = ANY($1::varchar[])
`

type ListShowTranslationsRow struct {
	ID            int32          `db:"id" json:"id"`
	ParentID      int32          `db:"parent_id" json:"parentID"`
	LanguagesCode string         `db:"languages_code" json:"languagesCode"`
	Title         null_v4.String `db:"title" json:"title"`
	Description   null_v4.String `db:"description" json:"description"`
}

func (q *Queries) ListShowTranslations(ctx context.Context, dollar_1 []string) ([]ListShowTranslationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listShowTranslations, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListShowTranslationsRow
	for rows.Next() {
		var i ListShowTranslationsRow
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.LanguagesCode,
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