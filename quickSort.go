package main

import(
	"fmt"
)

//Conquista
func partitioning (v []int, ini int, fim int) int{
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++{
		if v[i] <= pivot{
			aux := v[pIndex]
			v[pIndex] = v[i]
			v[i] = aux	
			pIndex += 1
		}
	}
	aux := v[pIndex]
	v[pIndex] = pivot
	v[fim] = aux
	return pIndex	
}

//Divisão
func quickSort(v []int, ini int, fim int){
	if fim > ini{
		pivotIndex := partitioning(v, ini, fim)
		quickSort(v, ini, pivotIndex-1)
		quickSort(v, pivotIndex+1, fim)
	}
}

func main(){
	v := [] int{8, 7, 5, 3, 2, 3, 4, 6, 9, 10} 
	quickSort(v, 0, len(v)-1)
	fmt.Println(v)
}