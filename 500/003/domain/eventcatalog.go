package domain

// catalog 相关事件
const (
	EventCatalogAdded = iota + 1
	EventCatalogDeleted
)

var EventCatalog Event[EventCatalogParam]

type EventCatalogParam struct {
	id   string
	name string
}
