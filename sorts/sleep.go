package main

import (
	`fmt`
	`time`
)

func main()  {
	var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	ch := make(chan int)
	for _,v := range sli{
		go func(v int) {
			time.Sleep(time.Duration(v)*10000000)
			ch <- v
		}(v)
	}
	for range sli{
		fmt.Println(<- ch)
	}
}
