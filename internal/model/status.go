package model

import (
	"giteasy/internal/constants"
	"giteasy/internal/observer"
)

var unstaged = make(map[string]string)
var staged = make(map[string]string)
var commited = make(map[string]string)

func Set(newSet map[string]string, status constants.StatusType) {
	switch status {
	case constants.UNSTAGED:
		unstaged = newSet
		for _, obs := range observer.Get(status) {
			obs.Notify()
		}
	case constants.STAGED:
		staged = newSet
		for _, obs := range observer.Get(status) {
			obs.Notify()
		}
	case constants.COMMITED:
		staged = newSet
		for _, obs := range observer.Get(status) {
			obs.Notify()
		}
	}

}

func Get(status constants.StatusType) map[string]string {
	switch status {
	case constants.UNSTAGED:
		return unstaged
	case constants.STAGED:
		return staged
	case constants.COMMITED:
		return staged
	}
	return nil
}
