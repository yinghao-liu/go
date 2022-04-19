package repo

import (
	cata "ddd/domain/catalog"
	"ddd/domain/errs"
	"ddd/domain/item"
)

// item仓储
type ItemsRepoer interface {
	Get(id string) (item.Item, errs.ErrorCoder)
	Add(it item.Item) errs.ErrorCoder
	Update(id string, it item.Item) errs.ErrorCoder
	Save(id string, it item.Item) errs.ErrorCoder
	Delete(id string) errs.ErrorCoder
}

// catalog仓储
type CatalogsRepoer interface {
	Get(id string) (cata.Catalog, errs.ErrorCoder)
	Add(ct cata.Catalog) errs.ErrorCoder
	Update(id string, ct cata.Catalog) errs.ErrorCoder
	Save(id string, ct cata.Catalog) errs.ErrorCoder
	Delete(id string) errs.ErrorCoder
}
