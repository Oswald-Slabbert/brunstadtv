package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Period is the resolver for the period field.
func (r *calendarResolver) Period(ctx context.Context, obj *model.Calendar, from string, to string) (*model.CalendarPeriod, error) {
	fromTime, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, err
	}
	toTime, err := time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, err
	}
	return r.periodResolver(ctx, fromTime, toTime)
}

// Day is the resolver for the day field.
func (r *calendarResolver) Day(ctx context.Context, obj *model.Calendar, day string) (*model.CalendarDay, error) {
	source, err := time.Parse(time.RFC3339, day)
	if err != nil {
		return nil, err
	}

	yy, mm, dd := source.Date()
	zone, offset := source.Zone()
	location := time.FixedZone(zone, offset)
	midnight := time.Date(yy, mm, dd, 0, 0, 0, 0, location)
	nextMidnight := midnight.Add(time.Hour * 24)

	events, err := getForPeriod(ctx, r.Loaders.EventLoader, r.Queries.GetEventsForPeriod, midnight, nextMidnight)
	if err != nil {
		return nil, err
	}
	entries, err := getForPeriod(ctx, r.FilteredLoaders(ctx).CalendarEntryLoader, r.Queries.GetCalendarEntriesForPeriod, midnight, nextMidnight)
	if err != nil {
		return nil, err
	}
	return &model.CalendarDay{
		Events:  utils.MapWithCtx(ctx, events, model.EventFrom),
		Entries: utils.MapWithCtx(ctx, entries, model.CalendarEntryFrom),
	}, nil
}

// Event is the resolver for the event field.
func (r *episodeCalendarEntryResolver) Event(ctx context.Context, obj *model.EpisodeCalendarEntry) (*model.Event, error) {
	if obj.Event == nil {
		return nil, nil
	}
	return r.QueryRoot().Event(ctx, obj.Event.ID)
}

// Title is the resolver for the title field.
func (r *episodeCalendarEntryResolver) Title(ctx context.Context, obj *model.EpisodeCalendarEntry) (string, error) {
	if obj.Title != "" {
		return obj.Title, nil
	}
	e, err := r.QueryRoot().Episode(ctx, obj.Episode.ID, nil)
	if err != nil {
		return "", nil
	}
	return e.Title, nil
}

// Description is the resolver for the description field.
func (r *episodeCalendarEntryResolver) Description(ctx context.Context, obj *model.EpisodeCalendarEntry) (string, error) {
	if obj.Description != "" {
		return obj.Description, nil
	}
	e, err := r.QueryRoot().Episode(ctx, obj.Episode.ID, nil)
	if err != nil {
		return "", nil
	}
	return e.Description, nil
}

// Episode is the resolver for the episode field.
func (r *episodeCalendarEntryResolver) Episode(ctx context.Context, obj *model.EpisodeCalendarEntry) (*model.Episode, error) {
	e, _ := r.QueryRoot().Episode(ctx, obj.Episode.ID, nil)
	return e, nil
}

// Event is the resolver for the event field.
func (r *seasonCalendarEntryResolver) Event(ctx context.Context, obj *model.SeasonCalendarEntry) (*model.Event, error) {
	if obj.Event == nil {
		return nil, nil
	}
	return r.QueryRoot().Event(ctx, obj.Event.ID)
}

// Title is the resolver for the title field.
func (r *seasonCalendarEntryResolver) Title(ctx context.Context, obj *model.SeasonCalendarEntry) (string, error) {
	if obj.Title != "" {
		return obj.Title, nil
	}
	s, err := r.QueryRoot().Season(ctx, obj.Season.ID)
	if err != nil {
		return "", nil
	}
	return s.Title, nil
}

// Description is the resolver for the description field.
func (r *seasonCalendarEntryResolver) Description(ctx context.Context, obj *model.SeasonCalendarEntry) (string, error) {
	if obj.Description != "" {
		return obj.Description, nil
	}
	s, err := r.QueryRoot().Season(ctx, obj.Season.ID)
	if err != nil {
		return "", nil
	}
	return s.Description, nil
}

// Season is the resolver for the season field.
func (r *seasonCalendarEntryResolver) Season(ctx context.Context, obj *model.SeasonCalendarEntry) (*model.Season, error) {
	return r.QueryRoot().Season(ctx, obj.Season.ID)
}

// Event is the resolver for the event field.
func (r *showCalendarEntryResolver) Event(ctx context.Context, obj *model.ShowCalendarEntry) (*model.Event, error) {
	if obj.Event == nil {
		return nil, nil
	}
	return r.QueryRoot().Event(ctx, obj.Event.ID)
}

// Title is the resolver for the title field.
func (r *showCalendarEntryResolver) Title(ctx context.Context, obj *model.ShowCalendarEntry) (string, error) {
	if obj.Title != "" {
		return obj.Title, nil
	}
	s, err := r.QueryRoot().Show(ctx, obj.Show.ID)
	if err != nil {
		return "", nil
	}
	return s.Title, nil
}

// Description is the resolver for the description field.
func (r *showCalendarEntryResolver) Description(ctx context.Context, obj *model.ShowCalendarEntry) (string, error) {
	if obj.Description != "" {
		return obj.Description, nil
	}
	s, err := r.QueryRoot().Show(ctx, obj.Show.ID)
	if err != nil {
		return "", nil
	}
	return s.Description, nil
}

// Show is the resolver for the show field.
func (r *showCalendarEntryResolver) Show(ctx context.Context, obj *model.ShowCalendarEntry) (*model.Show, error) {
	s, _ := r.QueryRoot().Show(ctx, obj.Show.ID)
	return s, nil
}

// Event is the resolver for the event field.
func (r *simpleCalendarEntryResolver) Event(ctx context.Context, obj *model.SimpleCalendarEntry) (*model.Event, error) {
	if obj.Event == nil {
		return nil, nil
	}
	return r.QueryRoot().Event(ctx, obj.Event.ID)
}

// Calendar returns generated.CalendarResolver implementation.
func (r *Resolver) Calendar() generated.CalendarResolver { return &calendarResolver{r} }

// EpisodeCalendarEntry returns generated.EpisodeCalendarEntryResolver implementation.
func (r *Resolver) EpisodeCalendarEntry() generated.EpisodeCalendarEntryResolver {
	return &episodeCalendarEntryResolver{r}
}

// SeasonCalendarEntry returns generated.SeasonCalendarEntryResolver implementation.
func (r *Resolver) SeasonCalendarEntry() generated.SeasonCalendarEntryResolver {
	return &seasonCalendarEntryResolver{r}
}

// ShowCalendarEntry returns generated.ShowCalendarEntryResolver implementation.
func (r *Resolver) ShowCalendarEntry() generated.ShowCalendarEntryResolver {
	return &showCalendarEntryResolver{r}
}

// SimpleCalendarEntry returns generated.SimpleCalendarEntryResolver implementation.
func (r *Resolver) SimpleCalendarEntry() generated.SimpleCalendarEntryResolver {
	return &simpleCalendarEntryResolver{r}
}

type calendarResolver struct{ *Resolver }
type episodeCalendarEntryResolver struct{ *Resolver }
type seasonCalendarEntryResolver struct{ *Resolver }
type showCalendarEntryResolver struct{ *Resolver }
type simpleCalendarEntryResolver struct{ *Resolver }
