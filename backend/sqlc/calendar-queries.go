package sqlc

import (
	"context"
	"encoding/json"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/samber/lo"
	"gopkg.in/guregu/null.v4"
	"time"
)

func mapToEvents(items []getEventsRow) []common.Event {
	return lo.Map(items, func(i getEventsRow, _ int) common.Event {
		var title common.LocaleString

		_ = json.Unmarshal(i.Title.RawMessage, &title)

		return common.Event{
			ID:    int(i.ID),
			Title: title,
			End:   i.End,
			Start: i.Start,
		}
	})
}

// GetEvents returns specific calendar events
func (q *Queries) GetEvents(ctx context.Context, ids []int) ([]common.Event, error) {
	items, err := q.getEvents(ctx, intToInt32(ids))
	if err != nil {
		return nil, err
	}
	return mapToEvents(items), nil
}

// ListEvents returns all CalendarEvents
func (q *Queries) ListEvents(ctx context.Context) ([]common.Event, error) {
	items, err := q.listEvents(ctx)
	if err != nil {
		return nil, err
	}
	return mapToEvents(lo.Map(items, func(i listEventsRow, _ int) getEventsRow {
		return getEventsRow(i)
	})), nil
}

// GetEventsForPeriod returns events for the specific period
func (q *Queries) GetEventsForPeriod(ctx context.Context, from time.Time, to time.Time) ([]int, error) {
	ids, err := q.getEventIDsForPeriod(ctx, getEventIDsForPeriodParams{
		from,
		to,
	})
	if err != nil {
		return nil, err
	}
	return int32ToInt(ids), nil
}

func mapToCalendarEntries(items []getCalendarEntriesRow) []common.CalendarEntry {
	return lo.Map(items, func(i getCalendarEntriesRow, _ int) common.CalendarEntry {
		var title common.LocaleString
		var description common.LocaleString

		_ = json.Unmarshal(i.Title.RawMessage, &title)
		_ = json.Unmarshal(i.Description.RawMessage, &description)

		var itemID null.Int
		switch i.LinkType.ValueOrZero() {
		case "episode":
			itemID = i.EpisodeID
		case "show":
			itemID = i.ShowID
		case "season":
			itemID = i.SeasonID
		}

		return common.CalendarEntry{
			ID:          int(i.ID),
			EventID:     i.EventID,
			Title:       title,
			Description: description,
			Start:       i.Start,
			End:         i.End,
			Type:        i.LinkType,
			ItemID:      itemID,
		}
	})
}

// GetCalendarEntries returns the specified entries
func (rq *RoleQueries) GetCalendarEntries(ctx context.Context, ids []int) ([]common.CalendarEntry, error) {
	items, err := rq.queries.getCalendarEntries(ctx, getCalendarEntriesParams{
		Column1: intToInt32(ids),
		Column2: rq.roles,
	})
	if err != nil {
		return nil, err
	}
	return mapToCalendarEntries(items), nil
}

// ListCalendarEntries returns all entries
func (rq *RoleQueries) ListCalendarEntries(ctx context.Context) ([]common.CalendarEntry, error) {
	items, err := rq.queries.listCalendarEntries(ctx, rq.roles)
	if err != nil {
		return nil, err
	}
	return mapToCalendarEntries(lo.Map(items, func(i listCalendarEntriesRow, _ int) getCalendarEntriesRow {
		return getCalendarEntriesRow(i)
	})), nil
}

// GetCalendarEntriesForPeriod returns events for the specific period
func (q *Queries) GetCalendarEntriesForPeriod(ctx context.Context, from time.Time, to time.Time) ([]int, error) {
	ids, err := q.getCalendarEntryIDsForPeriod(ctx, getCalendarEntryIDsForPeriodParams{
		from,
		to,
	})
	if err != nil {
		return nil, err
	}
	return int32ToInt(ids), err
}
