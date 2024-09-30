package demo03

import (
	"fmt"
	"sync"
	"testing"
)

// 测试用例函数，测试add函数
func Test_3_2(t *testing.T) {
	/*
		使用channel 在三个goroutine，传递数据。
		场景：product程序产生maxs个消息需要处理，指定work_num个worker（负责处理消息的goroutine），处理完成够任务自动结束。
		备注：当前会发出work_num个多余的message，用于结束product程序
	*/
	var wg sync.WaitGroup
	inc := make(chan int, 2)
	ouc := make(chan int, 1)
	maxs := 4 //需要处理maxs 消息
	work_num := 3
	// 创建两个goroutines
	for i := 1; i < work_num+1; i++ {
		wg.Add(1)
		go func(id int, cc, oo chan int) {
			worker(id, cc, oo, maxs, &wg)
		}(i, inc, ouc)
	}
	wg.Add(1)
	go product_msg(inc, ouc, maxs, &wg)
	ouc <- 0 // 数据产生信号
	wg.Wait()
	fmt.Println("Main: Both workers are finished")

}
