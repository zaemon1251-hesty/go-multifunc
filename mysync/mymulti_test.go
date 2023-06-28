package mysync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/zaemon1251-hesty/mysync"
)

func TestMulti_Do(t *testing.T) {
	m := mysync.NewMulti(5)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Do(func() {
				fmt.Println("hoge ", i)
				time.Sleep(10 * time.Second)
			})
		}(i)
	}

	wg.Wait()
}
