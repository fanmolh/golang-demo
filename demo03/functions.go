package demo03

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, inc, ouc chan int, maxs int, wg *sync.WaitGroup) {
	fmt.Printf("Worker%d starting\n", id)
	var enddata int
	for {
		enddata = <-inc
		fmt.Printf("Worker%d ,Recived: [%d]\n", id, enddata)
		// time.Sleep(time.Second / 2)
		if enddata < maxs+1 || enddata == 10 {
			fmt.Printf("Worker%d , processing complete mes: [%d]\n", id, enddata)
			ouc <- 0 // 通知生产者生继续生产
		} else {
			break
		}

	}
	defer func() {
		ouc <- 1
		fmt.Printf("Worker %d done,End Data %d (Not processed and returned)\n", id, enddata)
		wg.Done()
	}()
}

func product_msg(inc, ouc chan int, maxs int, wg *sync.WaitGroup) {
	fmt.Println("product_msg", runtime.NumGoroutine())
	count_work := runtime.NumGoroutine() - 3
	status := 0
	fmt.Println(maxs + count_work + 2)
	for i := 1; i < maxs+count_work+2; i++ {
		for {
			fmt.Println("generate number", i)
			// time.Sleep(time.Second * 1)
			status = (<-ouc + status)
			if status < count_work {
				fmt.Printf("product_msg for %d\n", i)
				inc <- i
			}
			break
		}
		if status == count_work {
			break
		}

	}
	defer func() {
		fmt.Println("product func done")
		wg.Done()
	}()
}
