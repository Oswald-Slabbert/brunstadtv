-- name: ListEpisodeTranslations :many
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
WHERE et.languages_code = ANY ($1::varchar[]);

-- name: ListSeasonTranslations :many
WITH seasons AS (SELECT s.id
                 FROM seasons s
                          LEFT JOIN shows sh ON sh.id = s.show_id
                 WHERE s.status = 'published'
                   AND sh.status = 'published')
SELECT et.id, seasons_id as parent_id, languages_code, title, description
FROM seasons_translations et
         JOIN seasons e ON e.id = et.seasons_id
WHERE et.languages_code = ANY ($1::varchar[]);

-- name: ListShowTranslations :many
WITH shows AS (SELECT s.id
               FROM shows s
               WHERE s.status = 'published')
SELECT et.id, shows_id as parent_id, languages_code, title, description
FROM shows_translations et
         JOIN shows e ON e.id = et.shows_id
WHERE et.languages_code = ANY ($1::varchar[]);

-- name: ListSectionTranslations :many
WITH sections AS (SELECT s.id
                  FROM sections s
                  WHERE s.status = 'published'
                    AND s.show_title = true)
SELECT st.id, sections_id as parent_id, languages_code, title, description
FROM sections_translations st
         JOIN sections e ON e.id = st.sections_id
WHERE st.languages_code = ANY ($1::varchar[]);

-- name: ListPageTranslations :many
WITH pages AS (SELECT s.id
               FROM pages s
               WHERE s.status = 'published')
SELECT st.id, pages_id as parent_id, languages_code, title, description
FROM pages_translations st
         JOIN pages e ON e.id = st.pages_id
WHERE st.languages_code = ANY ($1::varchar[]);

-- name: ListStudyTopicTranslations :many
WITH items AS (SELECT i.id
               FROM studytopics i
               WHERE i.status = 'published')
SELECT ts.id, studytopics_id as parent_id, languages_code, title
FROM studytopics_translations ts
         JOIN items i ON i.id = ts.studytopics_id
WHERE ts.languages_code = ANY ($1::varchar[]);

-- name: ListLessonTranslations :many
WITH lessons AS (SELECT s.id
                 FROM lessons s
                 WHERE s.status = 'published')
SELECT st.id, lessons_id as parent_id, languages_code, title
FROM lessons_translations st
         JOIN lessons e ON e.id = st.lessons_id
WHERE st.languages_code = ANY ($1::varchar[]);

-- name: ListTaskTranslations :many
WITH items AS (SELECT i.id
               FROM tasks i
               WHERE i.status = 'published')
SELECT ts.id, tasks_id as parent_id, languages_code, title
FROM tasks_translations ts
         JOIN items i ON i.id = ts.tasks_id
WHERE ts.languages_code = ANY ($1::varchar[]);

-- name: ListAlternativeTranslations :many
WITH items AS (SELECT i.id
               FROM questionalternatives i)
SELECT ts.id, questionalternatives_id as parent_id, languages_code, title
FROM questionalternatives_translations ts
         JOIN items i ON i.id = ts.questionalternatives_id
WHERE ts.languages_code = ANY ($1::varchar[]);

-- name: ListStudyTopicOriginalTranslations :many
SELECT items.id, items.title
FROM studytopics items
WHERE status = 'published';

-- name: ListLessonOriginalTranslations :many
SELECT items.id, items.title
FROM lessons items
WHERE status = 'published';

-- name: ListTaskOriginalTranslations :many
SELECT items.id, items.title
FROM tasks items
WHERE status = 'published';

-- name: ListQuestionAlternativesOriginalTranslations :many
SELECT items.id, items.title
FROM questionalternatives items
         JOIN tasks t ON t.id = items.task_id
WHERE t.status = 'published';
