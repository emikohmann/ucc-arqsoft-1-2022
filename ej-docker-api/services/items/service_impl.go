package items

import (
	"context"
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/items"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/search"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/utils/errs"
)

type ServiceImpl struct{}

func NewServiceImpl() ServiceImpl {
	return ServiceImpl{}
}

func (ServiceImpl) GetItem(ctx context.Context, itemID int64) (items.Item, errs.APIError) {
	return items.Item{
		ID:          itemID,
		Name:        fmt.Sprintf("Item %d", itemID),
		Description: "Hi! I'm an fake item, implement the service please",
	}, nil
}

func (ServiceImpl) SearchItems(ctx context.Context, query search.Query) (search.Result, errs.APIError) {
	results := []interface{}{
		items.Item{
			ID:          100,
			Name:        "Item 100",
			Description: "Hi! I'm an fake item 100, implement the service please",
		},
		items.Item{
			ID:          200,
			Name:        "Item 100",
			Description: "Hi! I'm an fake item 200, implement the service please",
		},
		items.Item{
			ID:          300,
			Name:        "Item 300",
			Description: "Hi! I'm an fake item 300, implement the service please",
		},
	}
	return search.Result{
		Results: results,
		Paging: search.Paging{
			Total:  len(results),
			Limit:  query.Limit,
			Offset: query.Offset,
		},
	}, nil
}
