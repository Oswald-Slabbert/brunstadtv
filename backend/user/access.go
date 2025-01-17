package user

import (
	"context"
	"time"

	"github.com/bcc-code/brunstadtv/backend/batchloaders"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/samber/lo"
)

// Sentinel errors
var (
	ErrItemNotPublished    = common.ErrItemNotPublished
	ErrItemNoAccess        = common.ErrItemNoAccess
	ErrPublishDateInFuture = merry.Sentinel("Publish date in the future")
)

type CheckConditions struct {
	FromDate    bool
	PublishDate bool
}

// ValidateAccess returns error if user in context does not have access to the specified item
func ValidateAccess[k comparable](
	ctx context.Context,
	permissionLoader *dataloader.Loader[k, *common.Permissions[k]],
	id k,
	conditions CheckConditions,
) error {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return err
	}
	rs := GetRolesFromCtx(ginCtx)

	perms, err := batchloaders.GetByID(ctx, permissionLoader, id)
	if err != nil {
		return err
	}

	roles := perms.Roles
	availability := perms.Availability

	if len(lo.Intersect(rs, roles.EarlyAccess)) > 0 && (availability.Published || availability.Unlisted) {
		return nil
	}

	if !(availability.Published || availability.Unlisted) ||
		availability.To.Before(time.Now()) ||
		(conditions.FromDate && availability.From.After(time.Now())) {
		return merry.Wrap(ErrItemNotPublished)
	}

	if len(lo.Intersect(rs, roles.Access)) == 0 {
		return merry.Wrap(ErrItemNoAccess)
	}

	if conditions.PublishDate && availability.PublishedOn.After(time.Now()) {
		return merry.Wrap(ErrPublishDateInFuture)
	}

	return nil
}
