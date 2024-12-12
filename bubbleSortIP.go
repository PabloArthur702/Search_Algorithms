package main

import(
	"fmt"
)


func bublleSort(v []int) {
	for varredura := 1; varredura < len(v); varredura++ {
		//varredura => (0,1,2,...,n-1)
		trocou := false
		for i := 1; i <= len(v) - varredura; i++{
			//quando varredura i = 0 => (0,1,2,...,n-1) => n-1 elementos
			//quando varredura i = 1 => (0,1,2,...,n-2) => n-2 elementos
			//quando varredura i = 2 => (0,1,2,...,n-3) => n-3 elementos
			//...
			//quando varredura i = n-1 => (0) => 1 elemento

			if v[i] < v[i-1]{
				v[i], v[i-1] = v[i-1], v[i]
				trocou = true
			}
	  	}
		if ! trocou{
			return
		}     
	}
  }

func main(){
	v := [] int {5, 9, 1, 2, 8, 10, 0}
	bublleSort(v)
	fmt.Println(v)
}

