package items

import (
	"context"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/items"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/search"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/utils/errs"
)

type Service interface {
	GetItem(ctx context.Context, itemID int64) (items.Item, errs.APIError)
	SearchItems(ctx context.Context, query search.Query) (search.Result, errs.APIError)
}
