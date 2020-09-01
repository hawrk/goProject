package design_pattern

import "sync"

// 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

// 懒汉式单例
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
