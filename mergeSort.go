package main

import(
	"fmt"
)


func Merge(v []int, e []int, d []int) {  //O(n)
	index_Esq := 0 
	index_Dir := 0
	index_v := 0 

	for (index_Dir < len(d)) && (index_Esq < len(e)){
		if e[index_Esq] <= d[index_Dir] {
			v[index_v] = e[index_Esq]
			index_Esq++	
		}else{
			v[index_v] = d[index_Dir]
			index_Dir++
		}
		index_v++
	}

	for (index_Dir < len(d)) {
			v[index_v] = d[index_Dir]
			index_Dir++
			index_v++	
	}

	for (index_Esq < len(e)) {
		v[index_v] = e[index_Esq]
		index_Esq++
		index_v++	
	}
}

func MergeSort(v []int) {      //T(n)
	if len(v)>1{
			tam_Esq := len(v)/2
			tam_Dir := len(v)-tam_Esq
		
			e := make([]int,tam_Esq)
			for index_Esq := 0; index_Esq < tam_Esq; index_Esq++{
					e[index_Esq] = v[index_Esq]
			}

			d := make([]int, tam_Dir)
			for index_Dir := tam_Esq; index_Dir < len(v); index_Dir++{
					d[index_Dir-tam_Esq] = v[index_Dir]
			}

			MergeSort(e)           
			MergeSort(d)
			Merge(v,e,d)           
	}
}

func main(){
	v := [] int{8, 7, 5, 3, 2, 1, 4, 6, 9, 10} 
	MergeSort(v)
	fmt.Println(v)
}

