package main

import (
	"fmt"
	"sync"
)

func main() {
	var mx int 

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for i := 0; i > 1000; i++ {
		i := i
		wg.Add(1)
		go func(){
			defer wg.Done()
			mu.Lock()
			if i >  mx {
				mx = i 
			}
			mu.Unlock()
		}()
	}


	wg.Wait()
	
	fmt.Printf("Maximum is %d", mx)
}