package syncTest

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
	golang channel 是一种类型化的消息队列
	1、chan 信道、也叫通道、管道，，可以通过它们发送类型化的数据在协程之间通信，可以避开所有内存共享导致的坑；
		通道的通信方式保证了同步性。数据通过通道：同一时间只有一个协程可以访问数据：所以不会出现数据竞争，设计如此。数据的归属（可以读写数据的能力）被传递。
	2、通道 普遍用于并发编程时 和 goroutine 配合使用
	3、chan 可以是任意类型，如: int、string、bool、等go内置数据类型，当然也可以自定义函数及struct 结构体
	4、chan 通道也是引用类型，所以我们使用 make() 函数来给它分配内存。
	5、chan 通道只有在 写满数据、或者读空数据是 才会阻塞
	

 */

func TestMessageQ1(t *testing.T) {
	// 模拟一个 mq ，生产者 消费者
	fmt.Println("Begin doing something!")
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			s := "shasjdk" + strconv.FormatInt(int64(i), 10)
			c <- s
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 5; i++ {
		s := <-c
		fmt.Println("Done!", s)
	}
	
	// 睡眠3秒
	time.Sleep(time.Second * 3)
	fmt.Println("最终完成")
}

func TestMessageQ2(t *testing.T) {
	// 创建一个 管道 c ， 循环放入5个消息之后，关闭管道。
	fmt.Println("Begin doing something!")
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			s := "shasjdk" + strconv.FormatInt(int64(i), 10)
			c <- s
			// time.Sleep(time.Second)
		}
		// 不关闭会造成死锁现象
		close(c)
	}()
	
	// 循环读取管道内的消息，
	// TODO 循环读管道时，必须要关闭管道，否则报错 all goroutines are asleep - deadlock!(所有的 goroutines 都已休眠， 形成死锁)
	for msg := range c {
		fmt.Println(msg)
	}
	
	// 睡眠3秒
	time.Sleep(time.Second * 3)
	fmt.Println("最终完成")
}

func TestMessageQ3(t *testing.T) {
	// 创建一个 有缓冲的管道 C ， 设置缓冲值为 3
	fmt.Println("Begin doing something!")
	c := make(chan string, 3)
	go func() {
		for i := 0; i < 5; i++ {
			s := "shasjdk" + strconv.FormatInt(int64(i), 10)
			c <- s
			// time.Sleep(time.Second)
		}
		// 不关闭会造成死锁现象
		close(c)
	}()
	
	// 循环读取管道内的消息，
	// TODO 循环读管道时，必须要关闭管道，否则报错 all goroutines are asleep - deadlock!(所有的 goroutines 都已休眠， 形成死锁)
	for msg := range c {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
	
	// 睡眠3秒
	time.Sleep(time.Second * 10)
	fmt.Println("最终完成")
}

func TestMessageQ4(t *testing.T) {
	sum := func(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // 将和送入 c
	}
	
	
	s := []int{7, 2, 8, -9, 4, 0}
	
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 睡眠一秒， 等待异步协程执行完毕
	time.Sleep(time.Second * 1)
	
	// 打印管道长度和容量
	fmt.Printf("容量: %v , 长度:%v  \n" , cap(c) , len(c))
	time.Sleep(time.Second * 1)
	// 这里需要注意， 管道的取值，每次只能 取一个， 那么如果想取多个值，就得用循环 或者以下方式
	x, y := <-c, <-c
	// 两次 <-c  表示取值两次，分别付给 x,y 两个变量
	
	fmt.Println(x, y, x+y)
	
}
