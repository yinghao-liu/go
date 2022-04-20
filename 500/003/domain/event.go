package domain

const (
	EventItemAdded = iota + 1
	EventItemDeleted
)

// 观察者
type Observer[E Events] interface {
	Notify(ev E)
}

// 事件
type Event[E Events] struct {
	obs map[int]map[Observer[E]]int
}

// 订阅
func (evt *Event[E]) Subscribe(evtype int, ob Observer[E]) {
	if nil == evt.obs {
		evt.obs = make(map[int]map[Observer[E]]int)
	}
	if nil == evt.obs[evtype] {
		evt.obs[evtype] = make(map[Observer[E]]int)
	}
}

// 取消订阅
func (evt *Event[E]) UnSubscribe(evtype int, ob Observer[E]) {
	if nil == evt.obs {
		return
	}
	delete(evt.obs[evtype], ob)
}

// 发布事件
func (evt Event[E]) Publish(evtype int, ev E) {
	if nil == evt.obs {
		return
	}
	for i, _ := range evt.obs[evtype] {
		go i.Notify(ev)
	}
}

type EventItem struct {
	id   string
	name string
}

type EventCatalog struct {
	id   string
	name string
}

type Events interface {
	EventItem | EventCatalog
}
