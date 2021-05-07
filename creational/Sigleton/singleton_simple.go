package Sigleton
// 单例模式
import "sync"

type singleton struct {

}

var instance *singleton
var once1 sync.Once

func GetInstance1() *singleton {
	once1.Do(func() {
		instance = &singleton{}
	})
	return instance
}