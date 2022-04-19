package mem

import (
	"ddd/domain/errs"
	"ddd/domain/item"
	"fmt"
)

type ItemsMemory struct {
	repo map[string]item.Item
}

func (its *ItemsMemory) Init() {
	if nil == its.repo {
		its.repo = make(map[string]item.Item)
	}
}
func (its ItemsMemory) Get(id string) (it item.Item, err errs.ErrorCoder) {
	it, exist := its.repo[id]
	if !exist {
		err = errs.ErrorCodeDefault(errs.ErrorResourceNotExist)
	}
	return
}
func (its *ItemsMemory) Add(it item.Item) errs.ErrorCoder {
	_, exist := its.repo[it.ID]
	if exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceIsExist)
	}
	its.repo[it.ID] = it
	fmt.Printf("its.repo %+v\n", its.repo)
	return nil

}
func (its *ItemsMemory) Update(id string, it item.Item) errs.ErrorCoder {
	_, exist := its.repo[it.ID]
	if !exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceNotExist)
	}
	its.repo[it.ID] = it
	return nil

}
func (its *ItemsMemory) Save(id string, it item.Item) errs.ErrorCoder {
	_, exist := its.repo[it.ID]
	if !exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceNotExist)
	}
	its.repo[it.ID] = it
	return nil

}
func (its *ItemsMemory) Delete(id string) errs.ErrorCoder {
	delete(its.repo, id)
	return nil
}
