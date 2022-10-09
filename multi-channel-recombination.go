// @Title  multi-channel-recombination
// @Description  多路复合计算器
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-09 23:54
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @title    doCompute
// @description   这个函数可以用来处理比较耗时的事情，比如计算
// @auth      MGAronya（张健）             2022-10-09 23:54
// @param     x int				入参
// @return    int				结果
func doCompute(x int) int {
	// TODO 模拟计算
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	// TODO 此处假设1 + x是一个很费时的计算
	return 1 + x
}

// @title    branch
// @description   每个分支开出一个goroutine来做计算，并把计算结果发送到各自通道里
// @auth      MGAronya（张健）             2022-10-09 23:54
// @param     x int					入参
// @return    chan int				用于传出结果
func branch(x int) chan int {
	ch := make(chan int)
	// TODO 开出一个goroutine来做计算，并把计算结果发送通道里
	go func() {
		ch <- doCompute(x)
	}()
	return ch
}

// @title    Recombination
// @description   将传入的多路通道复合
// @auth      MGAronya（张健）             2022-10-09 23:54
// @param     branches ...chan int			多路通道
// @return    chan int					    用于传出结果的通达
func Recombination(branches ...chan int) chan int {
	ch := make(chan int)

	// TODO select 会尝试着依次取出各个通道中的值
	for i := 0; i < len(branches); i++ {
		select {
		case v1 := <-branches[i]:
			// TODO 复合
			ch <- v1
		}
	}

	return ch
}

// @title    main
// @description   进行多路复合计算
// @auth      MGAronya（张健）             2022-10-09 23:54
// @param     void
// @return    void
func main() {
	// TODO 返回多路复合的结果
	result := Recombination(branch(10), branch(20), branch(30))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
