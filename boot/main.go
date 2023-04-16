package boot

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)

}
