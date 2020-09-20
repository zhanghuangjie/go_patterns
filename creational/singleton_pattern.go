package creational

import "sync"

type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

func NewSingleton() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	return instance
}
