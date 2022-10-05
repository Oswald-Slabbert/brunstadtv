package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Streams is the resolver for the streams field.
func (r *episodeResolver) Streams(ctx context.Context, obj *model.Episode) ([]*model.Stream, error) {
	intID, _ := strconv.ParseInt(obj.ID, 10, 32)
	streams, err := common.GetFromLoaderForKey(ctx, r.Resolver.Loaders.StreamsLoader, int(intID))
	if err != nil {
		return nil, err
	}

	var out []*model.Stream
	for _, s := range streams {
		stream, err := model.StreamFrom(ctx, r.URLSigner, r.Resolver.APIConfig, s)
		if err != nil {
			return nil, err
		}

		out = append(out, stream)
	}

	return out, nil
}

// Files is the resolver for the files field.
func (r *episodeResolver) Files(ctx context.Context, obj *model.Episode) ([]*model.File, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 32)
	if err != nil {
		return nil, err
	}

	files, err := common.GetFromLoaderForKey(ctx, r.Resolver.Loaders.FilesLoader, int(intID))
	if err != nil {
		return nil, err
	}

	var out []*model.File
	for _, f := range files {
		out = append(out, model.FileFrom(ctx, r.URLSigner, r.Resolver.APIConfig.GetFilesCDNDomain(), f))
	}
	return out, nil
}

// Season is the resolver for the season field.
func (r *episodeResolver) Season(ctx context.Context, obj *model.Episode) (*model.Season, error) {
	if obj.Season != nil {
		return r.QueryRoot().Season(ctx, obj.Season.ID)
	}
	return nil, nil
}

// Show is the resolver for the show field.
func (r *seasonResolver) Show(ctx context.Context, obj *model.Season) (*model.Show, error) {
	return r.QueryRoot().Show(ctx, obj.Show.ID)
}

// Episodes is the resolver for the episodes field.
func (r *seasonResolver) Episodes(ctx context.Context, obj *model.Season, first *int, offset *int) (*model.EpisodePagination, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemIDs, err := common.GetFromLoaderForKey(ctx, r.FilteredLoaders(ctx).EpisodesLoader, int(intID))
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(itemIDs, first, offset)

	episodes, err := common.GetManyFromLoader(ctx, r.Loaders.EpisodeLoader, utils.PointerIntArrayToIntArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.EpisodePagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, episodes, model.EpisodeFrom),
	}, nil
}

// FirstEpisode is the resolver for the firstEpisode field.
func (r *seasonResolver) FirstEpisode(ctx context.Context, obj *model.Season) (*model.Episode, error) {
	return firstOf(ctx, obj.ID, r.Loaders.EpisodePermissionLoader, r.Loaders.EpisodesLoader, r.QueryRoot().Episode)
}

// LastEpisode is the resolver for the lastEpisode field.
func (r *seasonResolver) LastEpisode(ctx context.Context, obj *model.Season) (*model.Episode, error) {
	return lastOf(ctx, obj.ID, r.Loaders.EpisodePermissionLoader, r.Loaders.EpisodesLoader, r.QueryRoot().Episode)
}

// Seasons is the resolver for the seasons field.
func (r *showResolver) Seasons(ctx context.Context, obj *model.Show, first *int, offset *int) (*model.SeasonPagination, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemIDs, err := common.GetFromLoaderForKey(ctx, r.FilteredLoaders(ctx).SeasonsLoader, int(intID))
	if err != nil {
		return nil, err
	}

	page := utils.Paginate(itemIDs, first, offset)

	seasons, err := common.GetManyFromLoader(ctx, r.Loaders.SeasonLoader, utils.PointerIntArrayToIntArray(page.Items))
	if err != nil {
		return nil, err
	}

	return &model.SeasonPagination{
		Total:  page.Total,
		First:  page.First,
		Offset: page.Offset,
		Items:  utils.MapWithCtx(ctx, seasons, model.SeasonFrom),
	}, nil
}

// FirstSeason is the resolver for the firstSeason field.
func (r *showResolver) FirstSeason(ctx context.Context, obj *model.Show) (*model.Season, error) {
	return firstOf(ctx, obj.ID, r.Loaders.SeasonPermissionLoader, r.Loaders.SeasonsLoader, r.QueryRoot().Season)
}

// LastSeason is the resolver for the lastSeason field.
func (r *showResolver) LastSeason(ctx context.Context, obj *model.Show) (*model.Season, error) {
	return lastOf(ctx, obj.ID, r.Loaders.SeasonPermissionLoader, r.Loaders.SeasonsLoader, r.QueryRoot().Season)
}

// Episode returns generated.EpisodeResolver implementation.
func (r *Resolver) Episode() generated.EpisodeResolver { return &episodeResolver{r} }

// Season returns generated.SeasonResolver implementation.
func (r *Resolver) Season() generated.SeasonResolver { return &seasonResolver{r} }

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type episodeResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
type showResolver struct{ *Resolver }
