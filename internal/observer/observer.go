package observer

import "giteasy/internal/constants"

type Observer interface {
	Notify()
}

type UnstageObserver struct {
	OnNotify func()
}

type StagedObserver struct {
	OnNotify func()
}

var observers map[constants.StatusType][]Observer = make(map[constants.StatusType][]Observer)

func (observer UnstageObserver) Notify() {
	observer.OnNotify()
}

func Register(observerType constants.StatusType, observer Observer) {
	existingObservers := observers[observerType]
	if existingObservers == nil {
		observers[observerType] = []Observer{observer}
	}
	existingObservers = append(existingObservers, observer)
}

func Get(observerType constants.StatusType) []Observer {
	return observers[observerType]
}
