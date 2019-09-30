package main

import (
	`fmt`
)

func main()  {
	var slis = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	var quik func([]int)[]int
	quik = func(sli []int)[]int {
		if len(sli) <= 1{
			return sli
		}
		base := sli[0]
		var left  []int
		var right  []int
		for i:=1;i<len(sli);i++{
			if sli[i] < base{
				left = append(left, sli[i])
			}else{
				right = append(right, sli[i])
			}
		}
		left = quik(left)
		right = quik(right)
		return append(append(left,base),right...)
	}
	fmt.Println(quik(slis))
}

