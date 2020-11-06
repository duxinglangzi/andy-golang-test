package basic

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 协程
func TestWait(t *testing.T) {
	// 类似于 java CyclicBarrier
	wait := sync.WaitGroup{}
	wait.Add(2)
	for i := 0; i < 2; i++ {
		go func(index int) {
			defer wait.Done()
			time.Sleep(time.Duration(rand.Int()%10) * time.Second)
			fmt.Printf("%d stop \n", index)
		}(i)
	}

	wait.Wait()

	fmt.Println("Finish")

}
