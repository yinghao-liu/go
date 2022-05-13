package repomem

import (
	cata "ddd/domain/catalog"
	"ddd/domain/errs"
)

type CatalogsMemory struct {
	repo map[string]cata.Catalog
}

func (cts *CatalogsMemory) Init() {
	if nil == cts.repo {
		cts.repo = make(map[string]cata.Catalog)
	}
}
func (cts CatalogsMemory) Get(id string) (ct cata.Catalog, err errs.ErrorCoder) {
	ct, exist := cts.repo[id]
	if !exist {
		err = errs.ErrorCodeDefault(errs.ErrorResourceNotFound)
	}
	return
}
func (cts *CatalogsMemory) Add(ct cata.Catalog) errs.ErrorCoder {
	it, exist := cts.repo[ct.ID]
	if exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceIsExist)
	}
	cts.repo[it.ID] = it
	return nil

}
func (cts *CatalogsMemory) Update(id string, ct cata.Catalog) errs.ErrorCoder {
	it, exist := cts.repo[ct.ID]
	if !exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceNotFound)
	}
	cts.repo[it.ID] = it
	return nil

}
func (cts *CatalogsMemory) Save(id string, ct cata.Catalog) errs.ErrorCoder {
	it, exist := cts.repo[ct.ID]
	if !exist {
		return errs.ErrorCodeDefault(errs.ErrorResourceNotFound)
	}
	cts.repo[it.ID] = it
	return nil

}
func (cts *CatalogsMemory) Delete(id string) errs.ErrorCoder {
	delete(cts.repo, id)
	return nil
}
