// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: studies.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getAnsweredTasks = `-- name: GetAnsweredTasks :many
SELECT ta.task_id
FROM "users"."taskanswers" ta
WHERE ta.profile_id = $1
  AND ta.task_id = ANY ($2::uuid[])
`

type GetAnsweredTasksParams struct {
	ProfileID uuid.UUID   `db:"profile_id" json:"profileID"`
	Column2   []uuid.UUID `db:"column_2" json:"column2"`
}

func (q *Queries) GetAnsweredTasks(ctx context.Context, arg GetAnsweredTasksParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getAnsweredTasks, arg.ProfileID, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var task_id uuid.UUID
		if err := rows.Scan(&task_id); err != nil {
			return nil, err
		}
		items = append(items, task_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setMessage = `-- name: SetMessage :exec
INSERT INTO "users"."messages" (id, item_id, message, updated_at, created_at, metadata)
VALUES ($1, $2, $3, NOW(), NOW(), $4)
ON CONFLICT (id) DO UPDATE SET message    = EXCLUDED.message,
                               metadata   = EXCLUDED.metadata,
                               updated_at = EXCLUDED.updated_at
`

type SetMessageParams struct {
	ID       string                `db:"id" json:"id"`
	ItemID   uuid.UUID             `db:"item_id" json:"itemID"`
	Message  string                `db:"message" json:"message"`
	Metadata pqtype.NullRawMessage `db:"metadata" json:"metadata"`
}

func (q *Queries) SetMessage(ctx context.Context, arg SetMessageParams) error {
	_, err := q.db.ExecContext(ctx, setMessage,
		arg.ID,
		arg.ItemID,
		arg.Message,
		arg.Metadata,
	)
	return err
}

const setTaskCompleted = `-- name: SetTaskCompleted :exec
INSERT INTO "users"."taskanswers" (profile_id, task_id, updated_at)
VALUES ($1, $2, NOW())
ON CONFLICT (profile_id, task_id) DO UPDATE SET updated_at = EXCLUDED.updated_at
`

type SetTaskCompletedParams struct {
	ProfileID uuid.UUID `db:"profile_id" json:"profileID"`
	TaskID    uuid.UUID `db:"task_id" json:"taskID"`
}

func (q *Queries) SetTaskCompleted(ctx context.Context, arg SetTaskCompletedParams) error {
	_, err := q.db.ExecContext(ctx, setTaskCompleted, arg.ProfileID, arg.TaskID)
	return err
}

const getCompletedLessons = `-- name: getCompletedLessons :many
WITH total AS (SELECT t.lesson_id,
                      COUNT(t.id) task_count
               FROM tasks t
               GROUP BY t.lesson_id),
     completed AS (SELECT t.lesson_id, ta.profile_id, COUNT(t.id) completed_count
                   FROM tasks t
                            JOIN "users"."taskanswers" ta ON ta.task_id = t.id
                   GROUP BY t.lesson_id, ta.profile_id)
SELECT total.lesson_id as id, p.id as parent_id
FROM users.profiles p
         JOIN completed ON completed.profile_id = p.id
         JOIN total ON total.lesson_id = completed.lesson_id
WHERE p.id = ANY ($1::uuid[])
  AND completed.completed_count = total.task_count
`

type getCompletedLessonsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getCompletedLessons(ctx context.Context, dollar_1 []uuid.UUID) ([]getCompletedLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompletedLessons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCompletedLessonsRow
	for rows.Next() {
		var i getCompletedLessonsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getCompletedTasks = `-- name: getCompletedTasks :many
SELECT ta.task_id as id, ta.profile_id as parent_id
FROM "users"."taskanswers" ta
WHERE ta.profile_id = ANY ($1::uuid[])
`

type getCompletedTasksRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getCompletedTasks(ctx context.Context, dollar_1 []uuid.UUID) ([]getCompletedTasksRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompletedTasks, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getCompletedTasksRow
	for rows.Next() {
		var i getCompletedTasksRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getEpisodesForLessons = `-- name: getEpisodesForLessons :many
SELECT rl.item       AS id,
       rl.lessons_id AS parent_id
FROM lessons_relations rl
         JOIN episode_availability access ON access.id = rl.item::int
         JOIN episode_roles roles ON roles.id = rl.item::int
WHERE rl.collection = 'episodes'
  AND access.published
  AND access.available_to > now()
  AND (
        (roles.roles && $2::varchar[] AND access.available_from < now()) OR
        (roles.roles_earlyaccess && $2::varchar[])
    )
  AND rl.lessons_id = ANY ($1::uuid[])
`

type getEpisodesForLessonsParams struct {
	Column1 []uuid.UUID `db:"column_1" json:"column1"`
	Column2 []string    `db:"column_2" json:"column2"`
}

type getEpisodesForLessonsRow struct {
	ID       null_v4.String `db:"id" json:"id"`
	ParentID uuid.NullUUID  `db:"parent_id" json:"parentID"`
}

func (q *Queries) getEpisodesForLessons(ctx context.Context, arg getEpisodesForLessonsParams) ([]getEpisodesForLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodesForLessons, pq.Array(arg.Column1), pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getEpisodesForLessonsRow
	for rows.Next() {
		var i getEpisodesForLessonsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getLessons = `-- name: getLessons :many
WITH ts AS (SELECT lessons_id,
                   json_object_agg(languages_code, title) as title
            FROM lessons_translations
            GROUP BY lessons_id)
SELECT l.id,
       l.topic_id,
       l.title as original_title,
       ts.title
FROM lessons l
         LEFT JOIN ts ON ts.lessons_id = l.id
WHERE l.status = 'published'
  AND l.id = ANY ($1::uuid[])
`

type getLessonsRow struct {
	ID            uuid.UUID             `db:"id" json:"id"`
	TopicID       uuid.UUID             `db:"topic_id" json:"topicID"`
	OriginalTitle string                `db:"original_title" json:"originalTitle"`
	Title         pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getLessons(ctx context.Context, dollar_1 []uuid.UUID) ([]getLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getLessons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getLessonsRow
	for rows.Next() {
		var i getLessonsRow
		if err := rows.Scan(
			&i.ID,
			&i.TopicID,
			&i.OriginalTitle,
			&i.Title,
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

const getLessonsForItemsInCollection = `-- name: getLessonsForItemsInCollection :many
SELECT rl.lessons_id AS id,
       rl.item       AS parent_id
FROM lessons_relations rl
WHERE rl.collection = $1
  AND rl.item = ANY ($2::varchar[])
`

type getLessonsForItemsInCollectionParams struct {
	Collection null_v4.String `db:"collection" json:"collection"`
	Column2    []string       `db:"column_2" json:"column2"`
}

type getLessonsForItemsInCollectionRow struct {
	ID       uuid.NullUUID  `db:"id" json:"id"`
	ParentID null_v4.String `db:"parent_id" json:"parentID"`
}

func (q *Queries) getLessonsForItemsInCollection(ctx context.Context, arg getLessonsForItemsInCollectionParams) ([]getLessonsForItemsInCollectionRow, error) {
	rows, err := q.db.QueryContext(ctx, getLessonsForItemsInCollection, arg.Collection, pq.Array(arg.Column2))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getLessonsForItemsInCollectionRow
	for rows.Next() {
		var i getLessonsForItemsInCollectionRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getLessonsForTopics = `-- name: getLessonsForTopics :many
SELECT l.id, l.topic_id AS parent_id
FROM lessons l
WHERE l.status = 'published'
  AND l.topic_id = ANY ($1::uuid[])
`

type getLessonsForTopicsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getLessonsForTopics(ctx context.Context, dollar_1 []uuid.UUID) ([]getLessonsForTopicsRow, error) {
	rows, err := q.db.QueryContext(ctx, getLessonsForTopics, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getLessonsForTopicsRow
	for rows.Next() {
		var i getLessonsForTopicsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getLinksForLessons = `-- name: getLinksForLessons :many
SELECT rl.item       AS id,
       rl.lessons_id AS parent_id
FROM lessons_relations rl
WHERE rl.collection = 'links'
  AND rl.lessons_id = ANY ($1::uuid[])
`

type getLinksForLessonsRow struct {
	ID       null_v4.String `db:"id" json:"id"`
	ParentID uuid.NullUUID  `db:"parent_id" json:"parentID"`
}

func (q *Queries) getLinksForLessons(ctx context.Context, dollar_1 []uuid.UUID) ([]getLinksForLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getLinksForLessons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getLinksForLessonsRow
	for rows.Next() {
		var i getLinksForLessonsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getQuestionAlternatives = `-- name: getQuestionAlternatives :many
WITH ts AS (SELECT questionalternatives_id, json_object_agg(languages_code, title) AS title
            FROM questionalternatives_translations
            GROUP BY questionalternatives_id)
SELECT qa.id, qa.title as original_title, qa.task_id, qa.is_correct, ts.title
FROM questionalternatives qa
         LEFT JOIN ts ON ts.questionalternatives_id = qa.id
WHERE qa.task_id = ANY ($1::uuid[])
`

type getQuestionAlternativesRow struct {
	ID            uuid.UUID             `db:"id" json:"id"`
	OriginalTitle null_v4.String        `db:"original_title" json:"originalTitle"`
	TaskID        uuid.NullUUID         `db:"task_id" json:"taskID"`
	IsCorrect     bool                  `db:"is_correct" json:"isCorrect"`
	Title         pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getQuestionAlternatives(ctx context.Context, dollar_1 []uuid.UUID) ([]getQuestionAlternativesRow, error) {
	rows, err := q.db.QueryContext(ctx, getQuestionAlternatives, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getQuestionAlternativesRow
	for rows.Next() {
		var i getQuestionAlternativesRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.TaskID,
			&i.IsCorrect,
			&i.Title,
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

const getTasks = `-- name: getTasks :many
WITH ts AS (SELECT tasks_id,
                   json_object_agg(languages_code, title)           as title,
                   json_object_agg(languages_code, description)     as description,
                   json_object_agg(languages_code, secondary_title) as secondary_title
            FROM tasks_translations
            GROUP BY tasks_id),
     images AS (SELECT img.task_id, json_object_agg(img.language, df.filename_disk) as images
                FROM tasks_images img
                         JOIN directus_files df ON df.id = img.image
                GROUP BY img.task_id)
SELECT t.id,
       t.title           as original_title,
       t.secondary_title as original_secondary_title,
       t.description     as original_description,
       t.type,
       t.question_type,
       t.lesson_id,
       t.alternatives_multiselect,
       t.image_type,
       t.link_id,
       t.episode_id,
       ts.title,
       ts.secondary_title,
       ts.description,
       images.images
FROM tasks t
         LEFT JOIN ts ON ts.tasks_id = t.id
         LEFT JOIN images ON images.task_id = t.id
WHERE t.status = 'published'
  AND t.id = ANY ($1::uuid[])
`

type getTasksRow struct {
	ID                      uuid.UUID             `db:"id" json:"id"`
	OriginalTitle           null_v4.String        `db:"original_title" json:"originalTitle"`
	OriginalSecondaryTitle  null_v4.String        `db:"original_secondary_title" json:"originalSecondaryTitle"`
	OriginalDescription     null_v4.String        `db:"original_description" json:"originalDescription"`
	Type                    string                `db:"type" json:"type"`
	QuestionType            null_v4.String        `db:"question_type" json:"questionType"`
	LessonID                uuid.UUID             `db:"lesson_id" json:"lessonID"`
	AlternativesMultiselect sql.NullBool          `db:"alternatives_multiselect" json:"alternativesMultiselect"`
	ImageType               null_v4.String        `db:"image_type" json:"imageType"`
	LinkID                  null_v4.Int           `db:"link_id" json:"linkID"`
	EpisodeID               null_v4.Int           `db:"episode_id" json:"episodeID"`
	Title                   pqtype.NullRawMessage `db:"title" json:"title"`
	SecondaryTitle          pqtype.NullRawMessage `db:"secondary_title" json:"secondaryTitle"`
	Description             pqtype.NullRawMessage `db:"description" json:"description"`
	Images                  pqtype.NullRawMessage `db:"images" json:"images"`
}

func (q *Queries) getTasks(ctx context.Context, dollar_1 []uuid.UUID) ([]getTasksRow, error) {
	rows, err := q.db.QueryContext(ctx, getTasks, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getTasksRow
	for rows.Next() {
		var i getTasksRow
		if err := rows.Scan(
			&i.ID,
			&i.OriginalTitle,
			&i.OriginalSecondaryTitle,
			&i.OriginalDescription,
			&i.Type,
			&i.QuestionType,
			&i.LessonID,
			&i.AlternativesMultiselect,
			&i.ImageType,
			&i.LinkID,
			&i.EpisodeID,
			&i.Title,
			&i.SecondaryTitle,
			&i.Description,
			&i.Images,
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

const getTasksForLessons = `-- name: getTasksForLessons :many
SELECT t.id, t.lesson_id AS parent_id
FROM tasks t
WHERE t.status = 'published'
  AND t.lesson_id = ANY ($1::uuid[])
`

type getTasksForLessonsRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	ParentID uuid.UUID `db:"parent_id" json:"parentID"`
}

func (q *Queries) getTasksForLessons(ctx context.Context, dollar_1 []uuid.UUID) ([]getTasksForLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTasksForLessons, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getTasksForLessonsRow
	for rows.Next() {
		var i getTasksForLessonsRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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

const getTopics = `-- name: getTopics :many
WITH ts AS (SELECT studytopics_id,
                   json_object_agg(languages_code, title) as title
            FROM studytopics_translations
            GROUP BY studytopics_id)
SELECT s.id,
       s.title as original_title,
       ts.title
FROM studytopics s
         LEFT JOIN ts ON ts.studytopics_id = s.id
WHERE s.status = 'published'
  AND s.id = ANY ($1::uuid[])
`

type getTopicsRow struct {
	ID            uuid.UUID             `db:"id" json:"id"`
	OriginalTitle string                `db:"original_title" json:"originalTitle"`
	Title         pqtype.NullRawMessage `db:"title" json:"title"`
}

func (q *Queries) getTopics(ctx context.Context, dollar_1 []uuid.UUID) ([]getTopicsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTopics, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getTopicsRow
	for rows.Next() {
		var i getTopicsRow
		if err := rows.Scan(&i.ID, &i.OriginalTitle, &i.Title); err != nil {
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
