package items

import (
	"context"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/items"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/domain/search"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/utils/errs"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) GetItem(ctx context.Context, itemID int64) (items.Item, errs.APIError) {
	//TODO implement me
	panic("implement me")
}

func (ServiceMock) SearchItems(ctx context.Context, query search.Query) (search.Result, errs.APIError) {
	//TODO implement me
	panic("implement me")
}
