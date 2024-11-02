package app

import "sync"

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}
