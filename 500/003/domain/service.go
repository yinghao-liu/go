package domain

import (
	"ddd/domain/errs"
	"ddd/domain/item"
	repo "ddd/domain/repository"
)

var ItemsRepo repo.ItemsRepoer
var CatalogsRepo repo.CatalogsRepoer

func ItemsGet(id string) (item.Item, errs.ErrorCoder) {
	return ItemsRepo.Get(id)
}
func ItemsAdd(it item.Item) errs.ErrorCoder {
	return ItemsRepo.Add(it)
}
