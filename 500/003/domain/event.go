package domain

// 观察者
type Observer[P EventParam] interface {
	Notify(param P)
}

// 事件
type Event[P EventParam] struct {
	obs map[int]map[Observer[P]]bool // set
}

// 订阅
func (evt *Event[P]) Subscribe(evtype int, ob Observer[P]) {
	if nil == evt.obs {
		evt.obs = make(map[int]map[Observer[P]]bool)
	}
	if nil == evt.obs[evtype] {
		evt.obs[evtype] = make(map[Observer[P]]bool)
	}
	evt.obs[evtype][ob] = true
}

// 取消订阅
func (evt *Event[P]) UnSubscribe(evtype int, ob Observer[P]) {
	if nil == evt.obs {
		return
	}
	delete(evt.obs[evtype], ob)
}

// 发布事件
func (evt Event[P]) Publish(evtype int, param P) {
	if nil == evt.obs {
		return
	}
	for i, _ := range evt.obs[evtype] {
		go i.Notify(param)
	}
}

type EventParam interface {
	EventItemParam | EventCatalogParam
}
