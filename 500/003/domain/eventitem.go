package domain

// Item 相关事件
const (
	EventItemAdded = iota + 1
	EventItemDeleted
)

var EventItem Event[EventItemParam]

type EventItemParam struct {
	id   string
	name string
}
