package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/admin/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/admin/model"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/samber/lo"
)

// Collection is the resolver for the collection field.
func (r *previewResolver) Collection(ctx context.Context, obj *model.Preview, filter string) (*model.PreviewCollection, error) {
	ctx = context.WithValue(ctx, "preview", true)

	var f common.Filter

	_ = json.Unmarshal([]byte(filter), &f)

	items, err := r.getItemsForFilter(ctx, f)
	if err != nil {
		return nil, err
	}

	return &model.PreviewCollection{
		Items: items,
	}, nil
}

// Asset is the resolver for the asset field.
func (r *previewResolver) Asset(ctx context.Context, obj *model.Preview, id string) (*model.PreviewAsset, error) {
	ctx = context.WithValue(ctx, "preview", true)

	intID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}

	streams, err := r.Queries.GetStreamsForAssets(ctx, []int{int(intID)})
	if err != nil || len(streams) == 0 {
		return nil, err
	}

	stream, found := lo.Find(streams, func(s common.Stream) bool {
		return s.Type == "hls_cmaf"
	})
	if !found {
		stream = streams[0]
	}
	return &model.PreviewAsset{
		URL:  stream.Url,
		Type: stream.Type,
	}, nil
}

// Preview is the resolver for the preview field.
func (r *queryRootResolver) Preview(ctx context.Context) (*model.Preview, error) {
	return &model.Preview{}, nil
}

// Statistics is the resolver for the statistics field.
func (r *queryRootResolver) Statistics(ctx context.Context) (*model.Statistics, error) {
	return &model.Statistics{}, nil
}

// LessonProgressGroupedByOrg is the resolver for the lessonProgressGroupedByOrg field.
func (r *statisticsResolver) LessonProgressGroupedByOrg(ctx context.Context, obj *model.Statistics, lessonID string, ageGroups []string, orgMaxSize *int, orgMinSize *int) ([]*model.ProgressByOrg, error) {
	minSize := 0
	maxSize := 9999999
	if orgMinSize != nil {
		minSize = *orgMinSize
	}

	if orgMaxSize != nil {
		maxSize = *orgMaxSize
	}

	stats, err := r.Queries.GetLessonProgressGroupedByOrg(ctx, sqlc.GetLessonProgressGroupedByOrgParams{
		AgeGroups: ageGroups,
		LessonID:  utils.AsUuid(lessonID),
		MinSize:   int32(minSize),
		MaxSize:   int32(maxSize),
	})
	if err != nil {
		return nil, err
	}

	out := []*model.ProgressByOrg{}
	for _, s := range stats {
		out = append(out, &model.ProgressByOrg{
			Name:     s.Name.String,
			Progress: s.Perc,
		})
	}

	return out, err
}

// Preview returns generated.PreviewResolver implementation.
func (r *Resolver) Preview() generated.PreviewResolver { return &previewResolver{r} }

// QueryRoot returns generated.QueryRootResolver implementation.
func (r *Resolver) QueryRoot() generated.QueryRootResolver { return &queryRootResolver{r} }

// Statistics returns generated.StatisticsResolver implementation.
func (r *Resolver) Statistics() generated.StatisticsResolver { return &statisticsResolver{r} }

type previewResolver struct{ *Resolver }
type queryRootResolver struct{ *Resolver }
type statisticsResolver struct{ *Resolver }
