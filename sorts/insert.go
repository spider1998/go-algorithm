package main

import (
	`fmt`
)

func main()  {
	var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	for i := range sli{
		tmp := sli[i]
		for j:=i-1;j>=0;j--{
			fmt.Println(j)
			if tmp < sli[j]{
				sli[j+1],sli[j] = sli[j],tmp
			}else{
				break
			}
		}
	}
	fmt.Println(sli)
}
