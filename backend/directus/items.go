package directus

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	"gopkg.in/guregu/null.v4"
	"strconv"
)

// Item model
type Item struct {
	ID     int           `json:"id,omitempty"`
	Status common.Status `json:"status"`
}

// Episode model
type Episode struct {
	Item
	SeasonID null.Int `json:"season_id"`
}

// Season model
type Season struct {
	Item
	ShowID int `json:"show_id"`
}

// Show model
type Show struct {
	Item
}

// UID retrieves the not so unique ID (but unique in collection)
func (i Item) UID() string {
	if i.ID == 0 {
		return ""
	}
	return strconv.Itoa(i.ID)
}

// ForUpdate retrieves a model with only mutable properties
func (i Item) ForUpdate() interface{} {
	return map[string]string{}
}

// GetStatus for this item
func (i Item) GetStatus() common.Status {
	return i.Status
}

// TypeName episodes
func (i Episode) TypeName() string {
	return "episodes"
}

// TypeName seasons
func (i Season) TypeName() string {
	return "seasons"
}

// TypeName shows
func (i Show) TypeName() string {
	return "shows"
}

// GetEpisode get an episode by id
func (h *Handler) GetEpisode(ctx context.Context, id int) (Episode, error) {
	return GetItem[Episode](ctx, h.c, "episodes", id)
}

// GetSeason get a season by id
func (h *Handler) GetSeason(ctx context.Context, id int) (Season, error) {
	return GetItem[Season](ctx, h.c, "seasons", id)
}

// GetShow get a show by id
func (h *Handler) GetShow(ctx context.Context, id int) (Show, error) {
	return GetItem[Show](ctx, h.c, "shows", id)
}

// ListEpisodes lists all episodes
func (h *Handler) ListEpisodes(ctx context.Context) ([]Episode, error) {
	return ListItems[Episode](ctx, h.c, "episodes", nil)
}

// ListSeasons lists all seasons
func (h *Handler) ListSeasons(ctx context.Context) ([]Season, error) {
	return ListItems[Season](ctx, h.c, "seasons", nil)
}

// ListShows lists all shows
func (h *Handler) ListShows(ctx context.Context) ([]Show, error) {
	return ListItems[Show](ctx, h.c, "shows", nil)
}
