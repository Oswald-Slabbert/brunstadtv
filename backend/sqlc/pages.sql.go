// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: pages.sql

package sqlc

import (
	"context"
	"encoding/json"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getPageIDsForCodes = `-- name: getPageIDsForCodes :many
SELECT p.id, p.code
FROM pages p
WHERE p.code = ANY ($1::varchar[])
`

type getPageIDsForCodesRow struct {
	ID   int32          `db:"id" json:"id"`
	Code null_v4.String `db:"code" json:"code"`
}

func (q *Queries) getPageIDsForCodes(ctx context.Context, dollar_1 []string) ([]getPageIDsForCodesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPageIDsForCodes, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPageIDsForCodesRow
	for rows.Next() {
		var i getPageIDsForCodesRow
		if err := rows.Scan(&i.ID, &i.Code); err != nil {
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

const getPages = `-- name: getPages :many
WITH t AS (SELECT ts.pages_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM pages_translations ts
           GROUP BY ts.pages_id),
     images AS (WITH images AS (SELECT page_id, style, language, filename_disk
                                FROM images img
                                         JOIN directus_files df on img.file = df.id)
                SELECT page_id, json_agg(images) as json
                FROM images
                GROUP BY page_id)
SELECT p.id::int                AS id,
       p.code::varchar          AS code,
       p.status = 'published'   AS published,
       COALESCE(img.json, '[]') as images,
       t.title,
       t.description
FROM pages p
         LEFT JOIN t ON t.pages_id = p.id
         LEFT JOIN images img ON img.page_id = p.id
WHERE p.id = ANY ($1::int[])
  AND p.status = 'published'
`

type getPagesRow struct {
	ID          int32                 `db:"id" json:"id"`
	Code        string                `db:"code" json:"code"`
	Published   bool                  `db:"published" json:"published"`
	Images      json.RawMessage       `db:"images" json:"images"`
	Title       pqtype.NullRawMessage `db:"title" json:"title"`
	Description pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) getPages(ctx context.Context, dollar_1 []int32) ([]getPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPages, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPagesRow
	for rows.Next() {
		var i getPagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Published,
			&i.Images,
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

const getPermissionsForPages = `-- name: getPermissionsForPages :many
WITH r AS (SELECT id                            AS page_id,
                  (SELECT array_agg(DISTINCT eu.usergroups_code) AS array_agg
                   FROM sections_usergroups eu) AS roles
           FROM pages)
SELECT p.id::int              AS id,
       p.status = 'published' AS published,
       roles.roles::varchar[] AS roles
FROM pages p
         LEFT JOIN r roles ON roles.page_id = p.id
WHERE p.id = ANY ($1::int[])
  AND p.status = 'published'
`

type getPermissionsForPagesRow struct {
	ID        int32    `db:"id" json:"id"`
	Published bool     `db:"published" json:"published"`
	Roles     []string `db:"roles" json:"roles"`
}

func (q *Queries) getPermissionsForPages(ctx context.Context, dollar_1 []int32) ([]getPermissionsForPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPermissionsForPages, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPermissionsForPagesRow
	for rows.Next() {
		var i getPermissionsForPagesRow
		if err := rows.Scan(&i.ID, &i.Published, pq.Array(&i.Roles)); err != nil {
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

const listPages = `-- name: listPages :many
WITH t AS (SELECT ts.pages_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM pages_translations ts
           GROUP BY ts.pages_id),
     images AS (WITH images AS (SELECT page_id, style, language, filename_disk
                                FROM images img
                                         JOIN directus_files df on img.file = df.id)
                SELECT page_id, json_agg(images) as json
                FROM images
                GROUP BY page_id)
SELECT p.id::int                AS id,
       p.code::varchar          AS code,
       p.status = 'published'   AS published,
       COALESCE(img.json, '[]') as images,
       t.title,
       t.description
FROM pages p
         LEFT JOIN t ON t.pages_id = p.id
         LEFT JOIN images img ON img.page_id = p.id
WHERE p.status = 'published'
`

type listPagesRow struct {
	ID          int32                 `db:"id" json:"id"`
	Code        string                `db:"code" json:"code"`
	Published   bool                  `db:"published" json:"published"`
	Images      json.RawMessage       `db:"images" json:"images"`
	Title       pqtype.NullRawMessage `db:"title" json:"title"`
	Description pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) listPages(ctx context.Context) ([]listPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, listPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listPagesRow
	for rows.Next() {
		var i listPagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Published,
			&i.Images,
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
