// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: sections.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getLinksForSection = `-- name: getLinksForSection :many
SELECT
    sl.id,
    sl.section_id,
    sl.page_id,
    sl.title,
    sl.url,
    df.filename_disk
FROM sections_links sl
    LEFT JOIN directus_files df on sl.icon = df.id
WHERE sl.section_id = ANY ($1::int[])
`

type getLinksForSectionRow struct {
	ID           int32          `db:"id" json:"id"`
	SectionID    int32          `db:"section_id" json:"sectionID"`
	PageID       null_v4.Int    `db:"page_id" json:"pageID"`
	Title        string         `db:"title" json:"title"`
	Url          null_v4.String `db:"url" json:"url"`
	FilenameDisk null_v4.String `db:"filename_disk" json:"filenameDisk"`
}

func (q *Queries) getLinksForSection(ctx context.Context, dollar_1 []int32) ([]getLinksForSectionRow, error) {
	rows, err := q.db.QueryContext(ctx, getLinksForSection, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getLinksForSectionRow
	for rows.Next() {
		var i getLinksForSectionRow
		if err := rows.Scan(
			&i.ID,
			&i.SectionID,
			&i.PageID,
			&i.Title,
			&i.Url,
			&i.FilenameDisk,
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

const getPermissionsForSections = `-- name: getPermissionsForSections :many
WITH u AS (SELECT ug.sections_id,
                  array_agg(ug.usergroups_code) AS roles
           FROM sections_usergroups ug
           GROUP BY ug.sections_id)
SELECT s.id,
       u.roles::varchar[] AS roles
FROM sections s
         JOIN pages p ON s.page_id = p.id
         LEFT JOIN u ON u.sections_id = s.id
WHERE s.id = ANY ($1::int[])
  AND s.status = 'published'
  AND p.status = 'published'
`

type getPermissionsForSectionsRow struct {
	ID    int32    `db:"id" json:"id"`
	Roles []string `db:"roles" json:"roles"`
}

func (q *Queries) getPermissionsForSections(ctx context.Context, dollar_1 []int32) ([]getPermissionsForSectionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPermissionsForSections, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPermissionsForSectionsRow
	for rows.Next() {
		var i getPermissionsForSectionsRow
		if err := rows.Scan(&i.ID, pq.Array(&i.Roles)); err != nil {
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

const getSectionIDsForPages = `-- name: getSectionIDsForPages :many
SELECT s.id::int AS id,
       p.id::int AS page_id
FROM sections s
         JOIN pages p ON s.page_id = p.id
WHERE p.id = ANY ($1::int[])
  AND s.status = 'published'
  AND p.status = 'published'
ORDER BY s.sort
`

type getSectionIDsForPagesRow struct {
	ID     int32 `db:"id" json:"id"`
	PageID int32 `db:"page_id" json:"pageID"`
}

func (q *Queries) getSectionIDsForPages(ctx context.Context, dollar_1 []int32) ([]getSectionIDsForPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getSectionIDsForPages, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSectionIDsForPagesRow
	for rows.Next() {
		var i getSectionIDsForPagesRow
		if err := rows.Scan(&i.ID, &i.PageID); err != nil {
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

const getSections = `-- name: getSections :many
WITH t AS (SELECT ts.sections_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM sections_translations ts
           GROUP BY ts.sections_id)
SELECT s.id,
       p.id::int                          AS page_id,
       s.type,
       s.style,
       s.link_style,
       s.size,
       s.grid_size,
       s.show_title,
       s.sort,
       s.status::text = 'published'::text AS published,
       s.collection_id,
       t.title,
       t.description
FROM sections s
         JOIN pages p ON s.page_id = p.id
         LEFT JOIN t ON s.id = t.sections_id
WHERE s.id = ANY ($1::int[])
  AND s.status = 'published'
  AND p.status = 'published'
`

type getSectionsRow struct {
	ID           int32                 `db:"id" json:"id"`
	PageID       int32                 `db:"page_id" json:"pageID"`
	Type         null_v4.String        `db:"type" json:"type"`
	Style        null_v4.String        `db:"style" json:"style"`
	LinkStyle    null_v4.String        `db:"link_style" json:"linkStyle"`
	Size         null_v4.String        `db:"size" json:"size"`
	GridSize     null_v4.String        `db:"grid_size" json:"gridSize"`
	ShowTitle    sql.NullBool          `db:"show_title" json:"showTitle"`
	Sort         null_v4.Int           `db:"sort" json:"sort"`
	Published    bool                  `db:"published" json:"published"`
	CollectionID null_v4.Int           `db:"collection_id" json:"collectionID"`
	Title        pqtype.NullRawMessage `db:"title" json:"title"`
	Description  pqtype.NullRawMessage `db:"description" json:"description"`
}

func (q *Queries) getSections(ctx context.Context, dollar_1 []int32) ([]getSectionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSections, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getSectionsRow
	for rows.Next() {
		var i getSectionsRow
		if err := rows.Scan(
			&i.ID,
			&i.PageID,
			&i.Type,
			&i.Style,
			&i.LinkStyle,
			&i.Size,
			&i.GridSize,
			&i.ShowTitle,
			&i.Sort,
			&i.Published,
			&i.CollectionID,
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
