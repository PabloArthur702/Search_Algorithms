package main

import(
	"fmt"
)

//Implace
func SelectionSortIP(v []int) {
	for varredura := 0; varredura < len(v) - 1; varredura ++ {
	  iMenor := varredura						//n-1 vezes
	  for i:=varredura + 1; i < len(v); i++{
			//qndo varredura => {0,1,2,...,n-2}
			//qndo varredura = 0 então i => {1,2,...,n-1} => n-1
			//qndo varredura = 1 então i => {2,...,n-1} => n-2
			//qndo varredura = 2 então i => {3,...,n-1} => n-3
			//...
			//qndo varredura = n-2 então i => {n-2} => 1

		  if v[i] < v[iMenor]{
			iMenor = i 
		  }
	  }     
	  v[iMenor], v[varredura] = v[varredura], v[iMenor]
	}
  }

func main(){
	v := [] int {5, 9, 1, 2, 8, 10, 0}
	SelectionSortIP(v)
	fmt.Println(v)
}

