package mysync_test

import (
	"testing"

	"github.com/zaemon1251-hesty/mysync"
)

func TestOnce_Do(t *testing.T) {
	var once mysync.Once

	defer func() {
		recover()
		once.Do(func() {
			panic("fuga")
		})
	}()

	once.Do(func() {
		panic("hoge")
	})
}
