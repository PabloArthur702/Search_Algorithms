package main

import(
	"fmt"
	"math"
)


func SelectionSortOP(v []int) []int {
	ordenado:= make( []int, len(v))
	for varredura := 1; varredura <= len(v); varredura ++ {
	  iMenor := varredura - 1
	  for i:=0; i < len(v); i++{
		  if v[i] < v[iMenor]{
			iMenor = i  
		  }
	  }    
	  ordenado[varredura -1] = v[iMenor] 
	  v[iMenor] = math.MaxInt
	}
	return ordenado  
  }

func main(){
	v := [] int {5, 9, 1, 2, 8, 10, 0}
	ord := SelectionSortOP(v)
	for i := 0; i < len(ord); i++{
		fmt.Println(ord[i])
	}
}