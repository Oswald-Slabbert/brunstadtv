-- name: listCalendarEntries :many
WITH t AS (SELECT ts.calendarentries_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM calendarentries_translations ts
           GROUP BY ts.calendarentries_id)
SELECT e.id,
       e.event_id,
       e.link_type,
       e.start,
       e.end,
       ea.id AS episode_id,
       se.id AS season_id,
       sh.id AS show_id,
       t.title,
       t.description
FROM calendarentries e
         LEFT JOIN episode_roles er ON er.id = e.episode_id AND er.roles && $1::varchar[]
         LEFT JOIN episode_availability ea ON ea.id = er.id AND ea.published
         LEFT JOIN seasons se ON se.id = e.season_id AND se.status = 'published'
         LEFT JOIN shows sh ON sh.id = e.show_id AND sh.status = 'published'
         LEFT JOIN t ON e.id = t.calendarentries_id
WHERE e.status = 'published';

-- name: getCalendarEntries :many
WITH t AS (SELECT ts.calendarentries_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM calendarentries_translations ts
           GROUP BY ts.calendarentries_id)
SELECT e.id,
       e.event_id,
       e.link_type,
       e.start,
       e.end,
       ea.id AS episode_id,
       se.id AS season_id,
       sh.id AS show_id,
       t.title,
       t.description
FROM calendarentries e
         LEFT JOIN episode_roles er ON er.id = e.episode_id AND er.roles && $2::varchar[]
         LEFT JOIN episode_availability ea ON ea.id = er.id AND ea.published
         LEFT JOIN seasons se ON se.id = e.season_id AND se.status = 'published'
         LEFT JOIN shows sh ON sh.id = e.show_id AND sh.status = 'published'
         LEFT JOIN t ON e.id = t.calendarentries_id
WHERE e.status = 'published'
  AND e.id = ANY ($1::int[]);

-- name: getCalendarEntriesForEvents :many
WITH t AS (SELECT ts.calendarentries_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM calendarentries_translations ts
           GROUP BY ts.calendarentries_id)
SELECT e.id,
       e.event_id,
       e.link_type,
       e.start,
       e.end,
       ea.id AS episode_id,
       se.id AS season_id,
       sh.id AS show_id,
       t.title,
       t.description
FROM calendarentries e
         LEFT JOIN episode_roles er ON er.id = e.episode_id AND er.roles && $2::varchar[]
         LEFT JOIN episode_availability ea ON ea.id = er.id AND ea.published
         LEFT JOIN seasons se ON se.id = e.season_id AND se.status = 'published'
         LEFT JOIN shows sh ON sh.id = e.show_id AND sh.status = 'published'
         LEFT JOIN t ON e.id = t.calendarentries_id
WHERE e.status = 'published'
  AND e.event_id = ANY ($1::int[]);

-- name: getCalendarEntryIDsForPeriod :many
SELECT e.id
FROM calendarentries e
WHERE e.status = 'published'
  AND ((e.start >= $1::timestamptz AND e.start <= $2::timestamptz)
    OR (e.end >= $1::timestamptz AND e.end <= $2::timestamptz)
    OR (e.start <= $1::timestamptz AND e.end >= $2::timestamptz))
ORDER BY e.start;

-- name: listEvents :many
WITH t AS (SELECT ts.events_id,
                  json_object_agg(ts.languages_code, ts.title) AS title
           FROM events_translations ts
           GROUP BY ts.events_id)
SELECT e.id,
       e.start,
       e.end,
       t.title
FROM events e
         LEFT JOIN t ON e.id = t.events_id
WHERE e.status = 'published';

-- name: getEvents :many
WITH t AS (SELECT ts.events_id,
                  json_object_agg(ts.languages_code, ts.title) AS title
           FROM events_translations ts
           GROUP BY ts.events_id)
SELECT e.id,
       e.start,
       e.end,
       t.title
FROM events e
         LEFT JOIN t ON e.id = t.events_id
WHERE e.status = 'published'
  AND e.id = ANY ($1::int[]);

-- name: getEventIDsForPeriod :many
SELECT e.id
FROM events e
WHERE e.status = 'published'
  AND ((e.start >= $1::timestamptz AND e.start <= $2::timestamptz) OR
       (e.end >= $1::timestamptz AND e.end <= $2::timestamptz))
ORDER BY e.start;
