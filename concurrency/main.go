package main

import (
	"fmt"
)

/* GO routine
เหมือนต่อคิว
เช่น ละเทียนเรียน
ลงทะเบียนไป
web รับมา ส่ง something ไปว่าเราได้รับข้อมูลแล้ว
หลังบ้านทำงานต่อ
*/

// // ถ้า main จบก่อนจะไม่ทำ doTask()
// func main() {
// 	fmt.Println("start program")
// 	wg := &sync.WaitGroup()
// 	wg.add(1)

// 	go doTask(wg)

// 	wg.wait()
// 	fmt.Println("end program")
// }

//	func doTask(wg *sync.WaitGroup) {
//		time.Sleep(200 * time.Millisecond)
//		fmt.Println("Do something..")
//		wg.Done()
//	}
//
// create channel (ท่อไว้ส่งค่า)
func main() {
	ch := make(chan int, 10)

	//ch <- "V" //send value to channel

	fmt.Println(<-ch) // receive date out of channel

	fibonacci(cap(ch), ch)
	for c := range ch {
		fmt.Println(c)
	}
}

func fibonacci(num int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		fmt.Println(x)
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
