package main

import (
	`fmt`
)

func main()  {
	sli  := []string{"2","3","9","1","4","0","3","7"}
	for i:=1;i<len(sli);i++{
		for j:=0;j<len(sli)-1;j++{
			if sli[i] < sli[j]{
				sli[i],sli[j] = sli[j],sli[i]
			}
		}
	}
	fmt.Println(sli)
}

