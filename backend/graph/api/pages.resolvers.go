package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/graph/api/generated"
	"github.com/bcc-code/brunstadtv/backend/graph/api/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
)

// Items is the resolver for the items field.
func (r *collectionResolver) Items(ctx context.Context, obj *model.Collection, first *int, offset *int) (*model.CollectionItemPagination, error) {
	pagination, err := collectionItemResolverFromCollection(ctx, r.Resolver, obj.ID, first, offset)
	if err != nil {
		return nil, err
	}

	return &model.CollectionItemPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Page is the resolver for the page field.
func (r *itemSectionResolver) Page(ctx context.Context, obj *model.ItemSection) (*model.Page, error) {
	return r.QueryRoot().Page(ctx, &obj.Page.ID, nil)
}

// Items is the resolver for the items field.
func (r *itemSectionResolver) Items(ctx context.Context, obj *model.ItemSection, first *int, offset *int) (*model.CollectionItemPagination, error) {
	pagination, err := collectionItemResolver(ctx, r.Resolver, obj.ID, first, offset)
	if err != nil {
		return nil, err
	}
	return &model.CollectionItemPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Sections is the resolver for the sections field.
func (r *pageResolver) Sections(ctx context.Context, obj *model.Page, first *int, offset *int) (*model.SectionPagination, error) {
	sections, err := itemsResolverForIntID(ctx, &itemLoaders[int, common.Section]{
		Item:        r.Loaders.SectionLoader,
		Permissions: r.Loaders.SectionPermissionLoader,
	}, r.Loaders.SectionsLoader, obj.ID, model.SectionFrom)
	if err != nil {
		return nil, err
	}
	pagination := utils.Paginate(sections, first, offset)
	return &model.SectionPagination{
		Total:  pagination.Total,
		First:  pagination.First,
		Offset: pagination.Offset,
		Items:  pagination.Items,
	}, nil
}

// Collection returns generated.CollectionResolver implementation.
func (r *Resolver) Collection() generated.CollectionResolver { return &collectionResolver{r} }

// ItemSection returns generated.ItemSectionResolver implementation.
func (r *Resolver) ItemSection() generated.ItemSectionResolver { return &itemSectionResolver{r} }

// Page returns generated.PageResolver implementation.
func (r *Resolver) Page() generated.PageResolver { return &pageResolver{r} }

type collectionResolver struct{ *Resolver }
type itemSectionResolver struct{ *Resolver }
type pageResolver struct{ *Resolver }