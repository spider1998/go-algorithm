package main

import (
	`fmt`
)

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func main()  {
	for i := range sli{
		k := i
		for j:=i+1;j<len(sli);j++{
			if sli[k] > sli[j] {
				k = j
			}
		}
		if k != i{
			sli[k],sli[i] = sli[i],sli[k]
		}
	}
	fmt.Println(sli)
}
